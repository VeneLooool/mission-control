package drones_api

import (
	"context"
	"fmt"

	"github.com/VeneLooool/mission-control/internal/config"
	"github.com/VeneLooool/mission-control/internal/model"
	"github.com/VeneLooool/mission-control/internal/pb/drones-api/api/v1/drones"
	proto_model "github.com/VeneLooool/mission-control/internal/pb/drones-api/api/v1/model"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	dronesApi drones.DronesClient
}

func New(ctx context.Context, cfg *config.DroneApiClientConfig) (*Client, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.Host, cfg.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, errors.Wrap(err, "grpc.NewClient()")
	}

	return &Client{
		dronesApi: drones.NewDronesClient(conn),
	}, nil
}

func (c *Client) StartDroneMission(ctx context.Context, droneID uint64, mission model.MissionPlan) error {
	_, err := c.dronesApi.StartDroneMission(ctx, &drones.StartDroneMission_Request{
		Id:      droneID,
		Mission: transformMissionToProto(mission),
	})
	if err != nil {
		return errors.Wrap(err, "c.droneClient.StartDroneMission()")
	}
	return nil
}

func transformMissionToProto(mission model.MissionPlan) *proto_model.Mission {
	return &proto_model.Mission{
		Coordinates: transformCoordinatesToProto(mission.Coordinates),
	}
}

func transformCoordinatesToProto(coordinates model.Coordinates) []*proto_model.Coordinate {
	protoCoordinates := make([]*proto_model.Coordinate, 0, len(coordinates))
	for _, coordinate := range coordinates {
		protoCoordinates = append(protoCoordinates, &proto_model.Coordinate{
			Latitude:  coordinate.Latitude,
			Longitude: coordinate.Longitude,
		})
	}
	return protoCoordinates
}
