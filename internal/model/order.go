package model

import "time"

type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `json:"user_id"`
	Items     []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"` // pending, paid, shipped, canceled
	CreatedAt time.Time   `json:"created_at"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"` // price at the time of order
}
