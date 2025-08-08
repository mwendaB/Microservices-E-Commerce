package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"order-service/internal/models"
)

// ServiceClient handles communication with other microservices
type ServiceClient struct {
	httpClient *http.Client
	userServiceURL    string
	productServiceURL string
}

// NewServiceClient creates a new service client for inter-service communication
func NewServiceClient(userServiceURL, productServiceURL string) *ServiceClient {
	return &ServiceClient{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		userServiceURL:    userServiceURL,
		productServiceURL: productServiceURL,
	}
}

// UserServiceResponse represents the response from user service
type UserServiceResponse struct {
	Success bool        `json:"success"`
	Data    models.User `json:"data"`
	Error   string      `json:"error"`
}

// ProductServiceResponse represents the response from product service
type ProductServiceResponse struct {
	Success bool           `json:"success"`
	Data    models.Product `json:"data"`
	Error   string         `json:"error"`
}

// GetUser retrieves user information from the user service
func (c *ServiceClient) GetUser(userID string) (*models.User, error) {
	url := fmt.Sprintf("%s/users/%s", c.userServiceURL, userID)
	
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user service returned status %d", resp.StatusCode)
	}

	var userResp UserServiceResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		return nil, fmt.Errorf("failed to decode user service response: %w", err)
	}

	if !userResp.Success {
		return nil, fmt.Errorf("user service error: %s", userResp.Error)
	}

	return &userResp.Data, nil
}

// GetProduct retrieves product information from the product service
func (c *ServiceClient) GetProduct(productID string) (*models.Product, error) {
	url := fmt.Sprintf("%s/products/%s", c.productServiceURL, productID)
	
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call product service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product service returned status %d", resp.StatusCode)
	}

	var productResp ProductServiceResponse
	if err := json.NewDecoder(resp.Body).Decode(&productResp); err != nil {
		return nil, fmt.Errorf("failed to decode product service response: %w", err)
	}

	if !productResp.Success {
		return nil, fmt.Errorf("product service error: %s", productResp.Error)
	}

	return &productResp.Data, nil
}

// ValidateOrderItems validates all items in an order by checking with services
func (c *ServiceClient) ValidateOrderItems(items []models.CreateOrderItem) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem

	for _, item := range items {
		// Get product information
		product, err := c.GetProduct(item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("invalid product %s: %w", item.ProductID, err)
		}

		// Check stock availability
		if product.Stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s: available %d, requested %d", 
				product.Name, product.Stock, item.Quantity)
		}

		// Create order item
		orderItem := models.NewOrderItem(product.ID, product.Name, product.Price, item.Quantity)
		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}

// CheckUserExists verifies that a user exists
func (c *ServiceClient) CheckUserExists(userID string) error {
	_, err := c.GetUser(userID)
	return err
}
