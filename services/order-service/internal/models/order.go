package models

import (
	"time"
	"github.com/google/uuid"
)

// OrderStatus represents the status of an order
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

// Order represents an order in the system
type Order struct {
	ID         string      `json:"id"`
	UserID     string      `json:"user_id"`
	Items      []OrderItem `json:"items"`
	TotalPrice float64     `json:"total_price"`
	Status     OrderStatus `json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

// OrderItem represents a single item in an order
type OrderItem struct {
	ProductID string  `json:"product_id"`
	ProductName string `json:"product_name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}

// CreateOrderRequest represents the request payload for creating an order
type CreateOrderRequest struct {
	UserID string            `json:"user_id" validate:"required"`
	Items  []CreateOrderItem `json:"items" validate:"required,min=1"`
}

// CreateOrderItem represents an item in the order creation request
type CreateOrderItem struct {
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

// UpdateOrderStatusRequest represents the request payload for updating order status
type UpdateOrderStatusRequest struct {
	Status OrderStatus `json:"status" validate:"required"`
}

// User represents user data from user service
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Product represents product data from product service
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// NewOrder creates a new order with generated ID and timestamps
func NewOrder(userID string, items []OrderItem) *Order {
	now := time.Now()
	
	// Calculate total price
	var totalPrice float64
	for _, item := range items {
		totalPrice += item.Subtotal
	}

	return &Order{
		ID:         uuid.New().String(),
		UserID:     userID,
		Items:      items,
		TotalPrice: totalPrice,
		Status:     OrderStatusPending,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

// NewOrderItem creates a new order item with calculated subtotal
func NewOrderItem(productID, productName string, price float64, quantity int) OrderItem {
	return OrderItem{
		ProductID:   productID,
		ProductName: productName,
		Price:       price,
		Quantity:    quantity,
		Subtotal:    price * float64(quantity),
	}
}

// UpdateStatus updates the order status and timestamp
func (o *Order) UpdateStatus(status OrderStatus) {
	o.Status = status
	o.UpdatedAt = time.Now()
}

// CanBeCancelled checks if the order can be cancelled
func (o *Order) CanBeCancelled() bool {
	return o.Status == OrderStatusPending || o.Status == OrderStatusConfirmed
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
