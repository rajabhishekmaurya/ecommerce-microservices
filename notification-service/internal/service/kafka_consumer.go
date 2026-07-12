package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"

	"github.com/rajabhishekmaurya/ecommerce-microservices/notification-service/internal/model"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer() *KafkaConsumer {

	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{"localhost:9092"},
			Topic:   "payment-events",
			GroupID: "notification-service",
		}),
	}
}

func (c *KafkaConsumer) Start() error {

	fmt.Println("Notification Service is waiting for events...")

	for {

		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}

		var event model.PaymentEvent

		if err := json.Unmarshal(msg.Value, &event); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("====================================")
		fmt.Println("Notification Received")
		fmt.Println("Order ID      :", event.OrderID)
		fmt.Println("TransactionID :", event.TransactionID)
		fmt.Println("Amount        :", event.Amount)
		fmt.Println("Status        :", event.Status)
		fmt.Println("Email Sent Successfully")
		fmt.Println("====================================")
	}
}
