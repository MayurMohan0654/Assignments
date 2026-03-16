package producers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"server/internal/models"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitializeOrderProducer() {
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "order-topic",
	})
	log.Println("Kafka producer initialized.")
}

func CloseOrder() {
	writer.Close()
}


func SendOrderToKAfka(order *models.Orders) error {
	orderBytes, err := json.Marshal(order)

	// log.Printf("this is orderBytes: %v", orderBytes)

	if err != nil {
		return fmt.Errorf("failed to marshal order: %w", err)
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(order.ID),
		Value: orderBytes,
	})

	if err != nil {
		return fmt.Errorf("Failed to send order to kafka: %w", err)
	}

	log.Printf("Oredr sent to kafka successfully: order id: %d", order.ID)
	return nil

}
