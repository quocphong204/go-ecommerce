package repository

import (
	"go-ecommerce/internal/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) GetByUserID(userID uint) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) GetByID(id string) (*model.Order, error) {
	var order model.Order
	err := r.db.First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetAll() ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) UpdateStatus(id string, status string) error {
	return r.db.Model(&model.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (r *OrderRepository) Delete(id string) error {
	return r.db.Delete(&model.Order{}, id).Error
}
