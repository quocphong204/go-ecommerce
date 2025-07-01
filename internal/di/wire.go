//go:build wireinject
// +build wireinject

package di

import (
	"go-ecommerce/internal/api/handler"
	"go-ecommerce/internal/config"
	"go-ecommerce/internal/repository"
	"go-ecommerce/internal/service"

	"github.com/google/wire"
)

func InitializeProductHandler() *handler.ProductHandler {
	wire.Build(
		config.GetDB,
		repository.NewProductRepository,
		service.NewProductService,
		handler.NewProductHandler,
	)
	return &handler.ProductHandler{}
}

func InitializeAuthHandler() *handler.AuthHandler {
	wire.Build(
		config.GetDB,
		repository.NewUserRepository,
		service.NewAuthService,
		handler.NewAuthHandler,
	)
	return &handler.AuthHandler{}
}

func InitializeOrderHandler() *handler.OrderHandler {
	wire.Build(
		config.GetDB,
		repository.NewOrderRepository,
		service.NewOrderService,
		handler.NewOrderHandler,
	)
	return &handler.OrderHandler{}
}
