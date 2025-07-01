// internal/consumer/order_consumer.go
package consumer

import (
	"encoding/json"
	"fmt"
	"go-ecommerce/internal/config"
	"go-ecommerce/internal/model"
)

func StartOrderConsumer() {
	ch := config.MQChannel

	q, _ := ch.QueueDeclare("", false, true, true, false, nil)
	_ = ch.QueueBind(q.Name, "", "order_exchange", false, nil)

	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	go func() {
		for msg := range msgs {
			var order model.Order
			if err := json.Unmarshal(msg.Body, &order); err == nil {
				fmt.Printf("📧 Đã gửi email xác nhận đơn hàng #%d cho user %d\n", order.ID, order.UserID)
			}
		}
	}()
}
