package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"order-service/internal/client"
	"order-service/internal/models"
	"order-service/internal/repository"

	"github.com/gorilla/mux"
)

// OrderHandler handles HTTP requests related to orders
type OrderHandler struct {
	repo   repository.OrderRepository
	client client.OrderValidationClient
}

// NewOrderHandler creates a new order handler
func NewOrderHandler(repo repository.OrderRepository, serviceClient client.OrderValidationClient) *OrderHandler {
	return &OrderHandler{
		repo:   repo,
		client: serviceClient,
	}
}

// CreateOrder handles POST /orders - creates a new order
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// Basic validation
	if req.UserID == "" || len(req.Items) == 0 {
		h.sendErrorResponse(w, http.StatusBadRequest, "User ID and at least one item are required")
		return
	}

	// Validate user exists
	if err := h.client.CheckUserExists(req.UserID); err != nil {
		log.Printf("User validation failed: %v", err)
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Validate and get order items
	orderItems, err := h.client.ValidateOrderItems(req.Items)
	if err != nil {
		log.Printf("Order items validation failed: %v", err)
		h.sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Create order
	order := models.NewOrder(req.UserID, orderItems)
	if err := h.repo.Create(order); err != nil {
		log.Printf("Error creating order: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to create order")
		return
	}

	response := models.Response{
		Success: true,
		Message: "Order created successfully",
		Data:    order,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetOrder handles GET /orders/{id} - retrieves an order by ID
func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	orderID := vars["id"]

	if orderID == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Order ID is required")
		return
	}

	order, err := h.repo.GetByID(orderID)
	if err != nil {
		log.Printf("Error getting order: %v", err)
		h.sendErrorResponse(w, http.StatusNotFound, "Order not found")
		return
	}

	response := models.Response{
		Success: true,
		Data:    order,
	}

	json.NewEncoder(w).Encode(response)
}

// GetUserOrders handles GET /orders/user/{user_id} - retrieves all orders for a user
func (h *OrderHandler) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userID := vars["user_id"]

	if userID == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// Validate user exists
	if err := h.client.CheckUserExists(userID); err != nil {
		log.Printf("User validation failed: %v", err)
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	orders, err := h.repo.GetByUserID(userID)
	if err != nil {
		log.Printf("Error getting user orders: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve orders")
		return
	}

	response := models.Response{
		Success: true,
		Data:    orders,
	}

	json.NewEncoder(w).Encode(response)
}

// UpdateOrderStatus handles PATCH /orders/{id}/status - updates order status
func (h *OrderHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	orderID := vars["id"]

	if orderID == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Order ID is required")
		return
	}

	var req models.UpdateOrderStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// Validate status
	validStatuses := []models.OrderStatus{
		models.OrderStatusPending,
		models.OrderStatusConfirmed,
		models.OrderStatusShipped,
		models.OrderStatusDelivered,
		models.OrderStatusCancelled,
	}

	isValidStatus := false
	for _, status := range validStatuses {
		if req.Status == status {
			isValidStatus = true
			break
		}
	}

	if !isValidStatus {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid order status")
		return
	}

	// Get existing order
	order, err := h.repo.GetByID(orderID)
	if err != nil {
		h.sendErrorResponse(w, http.StatusNotFound, "Order not found")
		return
	}

	// Check if order can be cancelled
	if req.Status == models.OrderStatusCancelled && !order.CanBeCancelled() {
		h.sendErrorResponse(w, http.StatusBadRequest, "Order cannot be cancelled in current status")
		return
	}

	// Update status
	order.UpdateStatus(req.Status)

	if err := h.repo.Update(order); err != nil {
		log.Printf("Error updating order status: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to update order status")
		return
	}

	response := models.Response{
		Success: true,
		Message: "Order status updated successfully",
		Data:    order,
	}

	json.NewEncoder(w).Encode(response)
}

// ListOrders handles GET /orders - retrieves all orders (admin function)
func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orders, err := h.repo.List()
	if err != nil {
		log.Printf("Error listing orders: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve orders")
		return
	}

	response := models.Response{
		Success: true,
		Data:    orders,
	}

	json.NewEncoder(w).Encode(response)
}

// HealthCheck handles GET /health - returns service health status
func (h *OrderHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.Response{
		Success: true,
		Message: "Order service is healthy",
		Data: map[string]string{
			"service": "order-service",
			"status":  "UP",
		},
	}

	json.NewEncoder(w).Encode(response)
}

// sendErrorResponse sends a standardized error response
func (h *OrderHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)

	response := models.Response{
		Success: false,
		Error:   message,
	}

	json.NewEncoder(w).Encode(response)
}
