package config

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	MQConn    *amqp.Connection
	MQChannel *amqp.Channel
)

func ConnectRabbitMQ() {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://guest:guest@rabbitmq:5672/" // default Docker URL
	}

	var err error

	// ✅ Mở kết nối
	MQConn, err = amqp.Dial(url)
	if err != nil {
		log.Fatalf("❌ Failed to connect to RabbitMQ: %v", err)
	}

	// ✅ Mở channel
	MQChannel, err = MQConn.Channel()
	if err != nil {
		log.Fatalf("❌ Failed to open a channel: %v", err)
	}

	// ✅ Khai báo exchange
	err = MQChannel.ExchangeDeclare(
		"order_exchange", // name
		"fanout",         // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatalf("❌ Failed to declare exchange: %v", err)
	}

	fmt.Println("✅ Connected to RabbitMQ")
}

func CloseRabbitMQ() {
	if MQChannel != nil {
		_ = MQChannel.Close()
	}
	if MQConn != nil {
		_ = MQConn.Close()
	}
}
