package dtos

import (
	"time"
)

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,min=0"`
	Stock       int     `json:"stock" validate:"required,min=0"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"omitempty"`
	Description string  `json:"description" validate:"omitempty"`
	Price       float64 `json:"price" validate:"omitempty,min=0"`
	Stock       int     `json:"stock" validate:"omitempty,min=0"`
}

type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
