package models

import (
	"time"
	"github.com/google/uuid"
)

// Product represents a product in the catalog
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	Stock       int       `json:"stock"`
	ImageURL    string    `json:"image_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateProductRequest represents the request payload for creating a product
type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=2"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,min=0"`
	Category    string  `json:"category" validate:"required"`
	Stock       int     `json:"stock" validate:"required,min=0"`
	ImageURL    string  `json:"image_url,omitempty"`
}

// UpdateProductRequest represents the request payload for updating a product
type UpdateProductRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Category    *string  `json:"category,omitempty"`
	Stock       *int     `json:"stock,omitempty"`
	ImageURL    *string  `json:"image_url,omitempty"`
}

// ProductFilter represents filtering options for product queries
type ProductFilter struct {
	Category  string  `json:"category,omitempty"`
	MinPrice  float64 `json:"min_price,omitempty"`
	MaxPrice  float64 `json:"max_price,omitempty"`
	InStock   bool    `json:"in_stock,omitempty"`
}

// NewProduct creates a new product with generated ID and timestamps
func NewProduct(name, description, category string, price float64, stock int, imageURL string) *Product {
	now := time.Now()
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
		Stock:       stock,
		ImageURL:    imageURL,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// IsInStock checks if the product has available stock
func (p *Product) IsInStock() bool {
	return p.Stock > 0
}

// ReduceStock reduces the product stock by the specified quantity
func (p *Product) ReduceStock(quantity int) bool {
	if p.Stock >= quantity {
		p.Stock -= quantity
		p.UpdatedAt = time.Now()
		return true
	}
	return false
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
