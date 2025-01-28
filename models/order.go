package models

import (
	"time"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID         uint        `gorm:"primaryKey"`
	UserID     uint        `gorm:"not null"`
	Status     OrderStatus `gorm:"type:varchar(20);default:'pending'"`
	TotalPrice float64     `gorm:"not null"`
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID        uint      `gorm:"primaryKey"`
	OrderID   uint      `gorm:"not null"`
	ProductID uint      `gorm:"not null"`
	Quantity  int       `gorm:"not null"`
	Price     float64   `gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
