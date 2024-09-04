package integration_tests

import (
	"context"
	"testing"
	"time"

	"golang-kafka-postgre/src/constant"
	kafkaPkg "golang-kafka-postgre/src/kafka"
	"golang-kafka-postgre/src/utils"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKafkaIntegration(t *testing.T) {
	// Initialize the database
	err := utils.InitDB("user=postgres password=yourRoamer14 host=localhost port=5432 dbname=gokafka-db sslmode=disable")
	require.NoError(t, err, "should connect to the database")

	// Create Kafka writer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{constant.BrokerAddress1},
		Topic:   constant.TopicCreateConnote,
	})

	// Produce a test message
	testMessage := "Test message for integration testing"
	testKey := "test-key"

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(testKey),
			Value: []byte(testMessage),
		},
	)
	require.NoError(t, err, "should write message to Kafka")

	// Start the consumer in a goroutine
	go kafkaPkg.StartConsumer() // Use the aliased package

	// Allow some time for the consumer to process the message
	time.Sleep(10 * time.Second)

	// Check the database for the message
	count, err := utils.CountEntries("test_table", "key1")
	require.NoError(t, err, "should query the database successfully")
	assert.Equal(t, 1, count, "should find one message in the database")

	writer.Close()
}
