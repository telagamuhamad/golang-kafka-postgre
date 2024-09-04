package interfaces

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaWriter interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type KafkaReader interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
	Close() error
}
