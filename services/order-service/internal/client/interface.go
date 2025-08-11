package client

import "order-service/internal/models"

// OrderValidationClient abstracts the validation operations needed by the order handler.
// Implemented by ServiceClient; enables mocking in tests.
type OrderValidationClient interface {
	CheckUserExists(userID string) error
	ValidateOrderItems(items []models.CreateOrderItem) ([]models.OrderItem, error)
}
