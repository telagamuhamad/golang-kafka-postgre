package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang-kafka-postgre/src/constant"
	"golang-kafka-postgre/src/interfaces"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer interfaces.KafkaWriter
}

func NewProducer(writer interfaces.KafkaWriter) *Producer {
	return &Producer{writer: writer}
}

func (p *Producer) Produce(ctx context.Context, msg string) {
	start := time.Now()
	i := 1

	err := p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(fmt.Sprintf("%d", i)),
		Value: []byte(msg),
	})

	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	elapsed := time.Since(start).Seconds()
	fmt.Printf("Message written in %v seconds\n", elapsed)
}

func NewKafkaWriter() interfaces.KafkaWriter {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{constant.BrokerAddress1, constant.BrokerAddress2, constant.BrokerAddress3},
		Topic:   constant.TopicCreateConnote,
	})
}
