package repository

import (
	"errors"
	"strings"
	"sync"
	"product-service/internal/models"
)

// ProductRepository defines the interface for product data operations
type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id string) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id string) error
	List(filter *models.ProductFilter) ([]*models.Product, error)
	GetByCategory(category string) ([]*models.Product, error)
	UpdateStock(id string, quantity int) error
}

// InMemoryProductRepository implements ProductRepository using in-memory storage
type InMemoryProductRepository struct {
	products map[string]*models.Product
	mutex    sync.RWMutex
}

// NewInMemoryProductRepository creates a new in-memory product repository with sample data
func NewInMemoryProductRepository() *InMemoryProductRepository {
	repo := &InMemoryProductRepository{
		products: make(map[string]*models.Product),
	}

	// Add sample products
	repo.seedData()
	return repo
}

// seedData adds sample products to the repository
func (r *InMemoryProductRepository) seedData() {
	sampleProducts := []*models.Product{
		models.NewProduct("MacBook Pro 16\"", "Apple MacBook Pro with M3 chip", "Electronics", 2499.99, 10, "https://example.com/macbook.jpg"),
		models.NewProduct("iPhone 15 Pro", "Latest iPhone with titanium design", "Electronics", 999.99, 25, "https://example.com/iphone.jpg"),
		models.NewProduct("Nike Air Max", "Comfortable running shoes", "Footwear", 129.99, 50, "https://example.com/nike.jpg"),
		models.NewProduct("Coffee Maker", "Automatic drip coffee maker", "Appliances", 89.99, 15, "https://example.com/coffee.jpg"),
		models.NewProduct("Wireless Headphones", "Noise-cancelling Bluetooth headphones", "Electronics", 199.99, 30, "https://example.com/headphones.jpg"),
	}

	for _, product := range sampleProducts {
		r.products[product.ID] = product
	}
}

// Create adds a new product to the repository
func (r *InMemoryProductRepository) Create(product *models.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Check if product with same name already exists
	for _, existingProduct := range r.products {
		if strings.EqualFold(existingProduct.Name, product.Name) {
			return errors.New("product with this name already exists")
		}
	}

	r.products[product.ID] = product
	return nil
}

// GetByID retrieves a product by its ID
func (r *InMemoryProductRepository) GetByID(id string) (*models.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	product, exists := r.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}

	// Return a copy to prevent external modification
	productCopy := *product
	return &productCopy, nil
}

// Update modifies an existing product
func (r *InMemoryProductRepository) Update(product *models.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.products[product.ID]; !exists {
		return errors.New("product not found")
	}

	r.products[product.ID] = product
	return nil
}

// Delete removes a product from the repository
func (r *InMemoryProductRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.products[id]; !exists {
		return errors.New("product not found")
	}

	delete(r.products, id)
	return nil
}

// List returns all products, optionally filtered
func (r *InMemoryProductRepository) List(filter *models.ProductFilter) ([]*models.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var products []*models.Product
	for _, product := range r.products {
		// Apply filters if provided
		if filter != nil {
			if filter.Category != "" && !strings.EqualFold(product.Category, filter.Category) {
				continue
			}
			if filter.MinPrice > 0 && product.Price < filter.MinPrice {
				continue
			}
			if filter.MaxPrice > 0 && product.Price > filter.MaxPrice {
				continue
			}
			if filter.InStock && product.Stock <= 0 {
				continue
			}
		}

		// Create a copy to prevent external modification
		productCopy := *product
		products = append(products, &productCopy)
	}

	return products, nil
}

// GetByCategory retrieves all products in a specific category
func (r *InMemoryProductRepository) GetByCategory(category string) ([]*models.Product, error) {
	filter := &models.ProductFilter{Category: category}
	return r.List(filter)
}

// UpdateStock updates the stock quantity for a product
func (r *InMemoryProductRepository) UpdateStock(id string, quantity int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	product, exists := r.products[id]
	if !exists {
		return errors.New("product not found")
	}

	if quantity < 0 {
		return errors.New("stock quantity cannot be negative")
	}

	product.Stock = quantity
	return nil
}
