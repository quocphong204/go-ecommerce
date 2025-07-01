package service

import (
	"go-ecommerce/internal/logger"
	"go-ecommerce/internal/model"
	"go-ecommerce/internal/producer"
	"go-ecommerce/internal/repository"
	"go.uber.org/zap"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetByUserID(userID uint) ([]model.Order, error) {
	return s.repo.GetByUserID(userID)
}

func (s *OrderService) GetByID(id string) (*model.Order, error) {
	return s.repo.GetByID(id)
}

func (s *OrderService) GetAll() ([]model.Order, error) {
	return s.repo.GetAll()
}

func (s *OrderService) UpdateStatus(id string, status string) error {
	return s.repo.UpdateStatus(id, status)
}

func (s *OrderService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *OrderService) Create(order *model.Order) error {
	if err := s.repo.Create(order); err != nil {
		logger.Log.Error("❌ Không thể tạo đơn hàng", zap.Error(err))
		return err
	}

	// Gửi message đến RabbitMQ (giả lập gửi email)
	producer.PublishOrderCreated(order)

	logger.Log.Info("✅ Đã tạo đơn hàng", zap.Int("order_id", int(order.ID)))
	return nil
}
