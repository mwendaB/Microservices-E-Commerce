package repository

import (
	"errors"
	"sync"
	"order-service/internal/models"
)

// OrderRepository defines the interface for order data operations
type OrderRepository interface {
	Create(order *models.Order) error
	GetByID(id string) (*models.Order, error)
	GetByUserID(userID string) ([]*models.Order, error)
	Update(order *models.Order) error
	List() ([]*models.Order, error)
	Delete(id string) error
}

// InMemoryOrderRepository implements OrderRepository using in-memory storage
type InMemoryOrderRepository struct {
	orders map[string]*models.Order
	mutex  sync.RWMutex
}

// NewInMemoryOrderRepository creates a new in-memory order repository
func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		orders: make(map[string]*models.Order),
	}
}

// Create adds a new order to the repository
func (r *InMemoryOrderRepository) Create(order *models.Order) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.orders[order.ID] = order
	return nil
}

// GetByID retrieves an order by its ID
func (r *InMemoryOrderRepository) GetByID(id string) (*models.Order, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	order, exists := r.orders[id]
	if !exists {
		return nil, errors.New("order not found")
	}

	// Return a copy to prevent external modification
	orderCopy := *order
	return &orderCopy, nil
}

// GetByUserID retrieves all orders for a specific user
func (r *InMemoryOrderRepository) GetByUserID(userID string) ([]*models.Order, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var userOrders []*models.Order
	for _, order := range r.orders {
		if order.UserID == userID {
			// Create a copy to prevent external modification
			orderCopy := *order
			userOrders = append(userOrders, &orderCopy)
		}
	}

	return userOrders, nil
}

// Update modifies an existing order
func (r *InMemoryOrderRepository) Update(order *models.Order) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.orders[order.ID]; !exists {
		return errors.New("order not found")
	}

	r.orders[order.ID] = order
	return nil
}

// List returns all orders
func (r *InMemoryOrderRepository) List() ([]*models.Order, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	orders := make([]*models.Order, 0, len(r.orders))
	for _, order := range r.orders {
		// Create a copy to prevent external modification
		orderCopy := *order
		orders = append(orders, &orderCopy)
	}

	return orders, nil
}

// Delete removes an order from the repository
func (r *InMemoryOrderRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.orders[id]; !exists {
		return errors.New("order not found")
	}

	delete(r.orders, id)
	return nil
}
