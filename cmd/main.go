package main

import (
	"go-ecommerce/internal/api"
	"go-ecommerce/internal/config"
	"go-ecommerce/internal/consumer"
	"go-ecommerce/internal/logger"

	_ "go-ecommerce/docs" // 👈 rất quan trọng để import docs
)

// @title Go Ecommerce API
// @version 1.0
// @description Backend API for ecommerce site written in Go
// @contact.name Your Name
// @contact.email your.email@example.com
// @license.name MIT
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Khởi tạo logger
	logger.InitLogger()

	// Kết nối RabbitMQ
	config.ConnectRabbitMQ()
	defer config.CloseRabbitMQ()

	// Bắt đầu consumer xử lý đơn hàng (giả lập gửi email)
	go consumer.StartOrderConsumer()

	// Khởi động router
	r := api.NewRouter()
	r.Run(":8080")
}
