package Consumers

import (
	"context"
	"encoding/json"
	"log"
	"server/internal/models"

	"github.com/segmentio/kafka-go"
)

func InitializeOrderConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "order-topic",
		GroupID: "order-consumer-group",
	})

	defer reader.Close()
	log.Println("order consumer started")

	for {
		msg, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Printf("error in consumer %w", err)
			continue
		}
		var order models.Orders
		err = json.Unmarshal(msg.Value, &order)

		if err != nil {
			log.Printf("Error unmarshling the order: %w", err)
			continue
		}

		log.Printf("rder consumed — ID: %s | facilityCode %s | Status: %s",
			order.ID,
			order.FacilityCode,
			order.Status,
		)

	}
}
