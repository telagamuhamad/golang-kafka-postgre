package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"golang-kafka-postgre/src/constant"
	"golang-kafka-postgre/src/interfaces"
	"golang-kafka-postgre/src/models"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader interfaces.KafkaReader
}

func StartConsumer() {
	reader := NewKafkaReader()
	defer reader.Close()

	consumer := NewConsumer(reader)
	ctx := context.Background()

	consumer.Consume(ctx)
}

func NewConsumer(reader interfaces.KafkaReader) *Consumer {
	return &Consumer{reader: reader}
}

func (c *Consumer) Consume(ctx context.Context) {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		if err := c.handleMessage(msg); err != nil {
			log.Printf("Error handling message: %v", err)
		}
	}
}

func (c *Consumer) handleMessage(msg kafka.Message) error {
	var connote models.Connote
	if err := json.Unmarshal(msg.Value, &connote); err != nil {
		return fmt.Errorf("failed to unmarshal message: %v", err)
	}

	fmt.Printf("Processing message: %+v\n", connote)

	// Process the message and interact with the database
	// Add your processing logic here

	return nil
}

func NewKafkaReader() interfaces.KafkaReader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{constant.BrokerAddress1, constant.BrokerAddress2, constant.BrokerAddress3},
		Topic:   constant.TopicCreateConnote,
		GroupID: "group-id",
	})
}
