package publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/VeneLooool/mission-control/internal/config"
	"github.com/VeneLooool/mission-control/internal/model"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

const (
	analyticTaskTopic = "analytic-task"
	retryCount        = 3
	retryDelay        = 500 * time.Millisecond
)

type Publisher struct {
	writer writer
}

func New(ctx context.Context, cfg *config.KafkaConfig) *Publisher {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(fmt.Sprintf("%s:%s", cfg.KafkaHost, cfg.KafkaPort)),
		Topic:                  analyticTaskTopic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	return &Publisher{
		writer: w,
	}
}

func (p *Publisher) Publish(ctx context.Context, task model.AnalyticTask) (err error) {
	marshalledEvent, err := json.Marshal(task)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte(task.GetEventKey()),
		Value: marshalledEvent,
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	for range retryCount {
		err := p.writer.WriteMessages(ctx, message)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(retryDelay)
			continue
		}
		if err != nil {
			log.Printf("failed to write messages: %s", err.Error())
		}
		break
	}
	return nil
}

func (p *Publisher) Close() error {
	return p.writer.Close()
}
