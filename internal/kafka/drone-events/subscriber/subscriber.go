package subscriber

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/VeneLooool/mission-control/internal/config"
	"github.com/VeneLooool/mission-control/internal/model"
	"github.com/segmentio/kafka-go"
)

const (
	droneEventsTopic                 = "drone-events"
	missionControlDroneEventsGroupID = "mission-control-drone-events"
)

type Subscriber struct {
	reader  Reader
	handler Handler
}

func New(ctx context.Context, handler Handler, cfg *config.KafkaConfig) *Subscriber {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("%s:%s", cfg.KafkaHost, cfg.KafkaPort)},
		GroupID: missionControlDroneEventsGroupID,
		Topic:   droneEventsTopic,
	})

	return &Subscriber{
		reader:  r,
		handler: handler,
	}
}

func (s *Subscriber) Subscribe(ctx context.Context) {
	go func() {
		defer s.Close()

		for {
			message, err := s.reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("Error reading message: %s", err.Error())
				break
			}

			log.Printf("Received message key: %s, value: %s", message.Key, message.Value)

			var event model.Event
			if err = json.Unmarshal(message.Value, &event); err != nil {
				log.Printf("Error unmarshalling event: %s", err.Error())
				continue
			}

			if err = s.handler.Handle(ctx, event); err != nil {
				log.Printf("Error handling event: %s", err.Error())
			}
		}
	}()
}

func (s *Subscriber) Close() error {
	return s.reader.Close()
}
