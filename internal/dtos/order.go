package dtos

import (
	"time"
)

type PlaceOrderRequest struct {
	Items []OrderItemRequest `json:"items" validate:"required,min=1"`
}

type OrderItemRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,min=1"`
}

type OrderResponse struct {
	ID         uint              `json:"id"`
	UserID     uint              `json:"user_id"`
	Status     string            `json:"status"`
	TotalPrice float64           `json:"total_price"`
	Items      []OrderItemDetail `json:"items"`
	CreatedAt  time.Time         `json:"created_at"`
}

type OrderItemDetail struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
