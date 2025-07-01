// internal/producer/order_publisher.go
package producer

import (
	"encoding/json"
	"fmt"
	"go-ecommerce/internal/config"
	"go-ecommerce/internal/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishOrderCreated(order *model.Order) {
	data, _ := json.Marshal(order)

	err := config.MQChannel.Publish(
		"order_exchange", // exchange
		"",               // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
	if err != nil {
		fmt.Println("❌ Failed to publish order:", err)
	} else {
		fmt.Println("📨 Order message sent to RabbitMQ")
	}
}
