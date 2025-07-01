package consumer

import (
	"fmt"
	"go-ecommerce/internal/config"

)

func StartOrderConsumer() {
	ch := config.MQChannel
	q, err := ch.QueueDeclare(
		"order_queue", // name
		true,          // durable
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	err = ch.QueueBind(
		q.Name,
		"", // routing key
		"order_exchange",
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true, // auto-ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	go func() {
		for d := range msgs {
			fmt.Printf("📥 Received Order Message: %s\n", d.Body)
			// tại đây bạn có thể gọi logic xử lý, gửi email v.v.
		}
	}()
}
