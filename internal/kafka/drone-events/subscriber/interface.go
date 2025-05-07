package subscriber

import (
	"context"

	"github.com/VeneLooool/mission-control/internal/model"
	"github.com/segmentio/kafka-go"
)

type Reader interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
	Close() error
}

type Handler interface {
	Handle(ctx context.Context, event model.Event) error
}
