package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"` // lưu hash
	CreatedAt time.Time `json:"created_at"`
	Role     string `gorm:"default:user"`
}
