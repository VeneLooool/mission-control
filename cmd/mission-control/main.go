package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/VeneLooool/mission-control/internal/app/api/v1/control"
	drones_api "github.com/VeneLooool/mission-control/internal/clients/drones-api"
	missions_api "github.com/VeneLooool/mission-control/internal/clients/missions-api"
	"github.com/VeneLooool/mission-control/internal/config"
	control_cron "github.com/VeneLooool/mission-control/internal/cron/control"
	"github.com/VeneLooool/mission-control/internal/handlers/drone_events"
	"github.com/VeneLooool/mission-control/internal/kafka/analytic-tasks/publisher"
	"github.com/VeneLooool/mission-control/internal/kafka/drone-events/subscriber"
	pb "github.com/VeneLooool/mission-control/internal/pb/api/v1/control"
	control_uc "github.com/VeneLooool/mission-control/internal/usecase/control"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.New(ctx)
	if err != nil {
		log.Fatalf("failed to create new config: %s", err.Error())
	}

	go func() {
		if err := runGRPC(ctx, cfg); err != nil {
			log.Fatal(err)
		}
	}()

	if err := runHTTPGateway(ctx, cfg); err != nil {
		log.Fatal(err)
	}
}

func runGRPC(ctx context.Context, cfg *config.Config) error {
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	controlServer, err := newServices(ctx, cfg)
	if err != nil {
		return err
	}
	pb.RegisterMissionControlServer(grpcServer, controlServer)

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	log.Printf("gRPC server listening on :%s\n", cfg.GrpcPort)
	if err = grpcServer.Serve(grpcListener); err != nil {
		return err
	}
	return nil
}

func runHTTPGateway(ctx context.Context, cfg *config.Config) error {
	mux := runtime.NewServeMux()
	err := pb.RegisterMissionControlHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%s", cfg.GrpcPort), []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		log.Fatalf("failed to register gateway: %s", err.Error())
	}

	// Serve Swagger JSON and Swagger UI
	fs := http.FileServer(http.Dir("./swagger-ui")) // директория со статикой UI
	http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", fs))

	// Serve Swagger JSON файл
	http.HandleFunc("/swagger/control.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./internal/pb/api/v1/control/control.swagger.json")
	})

	withCORS := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			// Для preflight-запросов
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			h.ServeHTTP(w, r)
		})
	}

	// gRPC → REST mux
	http.Handle("/", withCORS(mux))

	log.Printf("HTTP gateway listening on :%s\n", cfg.HttpPort)
	if err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.HttpPort), nil); err != nil {
		return err
	}

	return nil
}

func newServices(ctx context.Context, cfg *config.Config) (*control.Implementation, error) {
	analyticTaskPublisher := publisher.New(ctx, cfg.GetKafkaConfig())

	droneClient, err := drones_api.New(ctx, cfg.GetDroneApiClientConfig())
	if err != nil {
		return nil, err
	}

	missionClient, err := missions_api.New(ctx, cfg.GetMissionApiClientConfig())
	if err != nil {
		return nil, err
	}

	controlUC := control_uc.New(missionClient)

	newHandlers(ctx, missionClient, analyticTaskPublisher, cfg)
	newCron(ctx, missionClient, droneClient)

	return control.NewService(controlUC), nil
}

func newHandlers(ctx context.Context, missionClient *missions_api.Client, analyticPublisher *publisher.Publisher, cfg *config.Config) {
	handler := drone_events.New(missionClient, analyticPublisher)
	sub := subscriber.New(ctx, handler, cfg.GetKafkaConfig())
	sub.Subscribe(ctx)
}

func newCron(ctx context.Context, missionClient *missions_api.Client, droneClient *drones_api.Client) {
	cron := control_cron.New(missionClient, droneClient)
	go func() {
		if err := cron.Do(ctx); err != nil {
			log.Fatalf("Cron error: %s", err.Error())
		}
	}()
}
