package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"product-service/internal/models"
	"product-service/internal/repository"

	"github.com/gorilla/mux"
)

// ProductHandler handles HTTP requests related to products
type ProductHandler struct {
	repo repository.ProductRepository
}

// NewProductHandler creates a new product handler
func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		repo: repo,
	}
}

// CreateProduct handles POST /products - creates a new product
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// Basic validation
	if req.Name == "" || req.Category == "" || req.Price <= 0 {
		h.sendErrorResponse(w, http.StatusBadRequest, "Name, category, and positive price are required")
		return
	}

	// Create product
	product := models.NewProduct(req.Name, req.Description, req.Category, req.Price, req.Stock, req.ImageURL)
	if err := h.repo.Create(product); err != nil {
		log.Printf("Error creating product: %v", err)
		h.sendErrorResponse(w, http.StatusConflict, err.Error())
		return
	}

	response := models.Response{
		Success: true,
		Message: "Product created successfully",
		Data:    product,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetProduct handles GET /products/{id} - retrieves a product by ID
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	productID := vars["id"]

	if productID == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Product ID is required")
		return
	}

	product, err := h.repo.GetByID(productID)
	if err != nil {
		log.Printf("Error getting product: %v", err)
		h.sendErrorResponse(w, http.StatusNotFound, "Product not found")
		return
	}

	response := models.Response{
		Success: true,
		Data:    product,
	}

	json.NewEncoder(w).Encode(response)
}

// ListProducts handles GET /products - retrieves all products with optional filtering
func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters for filtering
	filter := &models.ProductFilter{}
	
	if category := r.URL.Query().Get("category"); category != "" {
		filter.Category = category
	}
	
	if minPriceStr := r.URL.Query().Get("min_price"); minPriceStr != "" {
		if minPrice, err := strconv.ParseFloat(minPriceStr, 64); err == nil {
			filter.MinPrice = minPrice
		}
	}
	
	if maxPriceStr := r.URL.Query().Get("max_price"); maxPriceStr != "" {
		if maxPrice, err := strconv.ParseFloat(maxPriceStr, 64); err == nil {
			filter.MaxPrice = maxPrice
		}
	}
	
	if inStockStr := r.URL.Query().Get("in_stock"); inStockStr == "true" {
		filter.InStock = true
	}

	products, err := h.repo.List(filter)
	if err != nil {
		log.Printf("Error listing products: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve products")
		return
	}

	response := models.Response{
		Success: true,
		Data:    products,
	}

	json.NewEncoder(w).Encode(response)
}

// GetProductsByCategory handles GET /products/category/{category} - retrieves products by category
func (h *ProductHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	category := vars["category"]

	if category == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Category is required")
		return
	}

	products, err := h.repo.GetByCategory(category)
	if err != nil {
		log.Printf("Error getting products by category: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve products")
		return
	}

	response := models.Response{
		Success: true,
		Data:    products,
	}

	json.NewEncoder(w).Encode(response)
}

// UpdateProduct handles PUT /products/{id} - updates an existing product
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	productID := vars["id"]

	if productID == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Product ID is required")
		return
	}

	// Get existing product
	existingProduct, err := h.repo.GetByID(productID)
	if err != nil {
		h.sendErrorResponse(w, http.StatusNotFound, "Product not found")
		return
	}

	var req models.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// Update fields if provided
	if req.Name != nil {
		existingProduct.Name = *req.Name
	}
	if req.Description != nil {
		existingProduct.Description = *req.Description
	}
	if req.Price != nil {
		existingProduct.Price = *req.Price
	}
	if req.Category != nil {
		existingProduct.Category = *req.Category
	}
	if req.Stock != nil {
		existingProduct.Stock = *req.Stock
	}
	if req.ImageURL != nil {
		existingProduct.ImageURL = *req.ImageURL
	}

	if err := h.repo.Update(existingProduct); err != nil {
		log.Printf("Error updating product: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	response := models.Response{
		Success: true,
		Message: "Product updated successfully",
		Data:    existingProduct,
	}

	json.NewEncoder(w).Encode(response)
}

// UpdateStock handles PATCH /products/{id}/stock - updates product stock
func (h *ProductHandler) UpdateStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	productID := vars["id"]

	if productID == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Product ID is required")
		return
	}

	var req struct {
		Stock int `json:"stock"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	if err := h.repo.UpdateStock(productID, req.Stock); err != nil {
		log.Printf("Error updating stock: %v", err)
		h.sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	response := models.Response{
		Success: true,
		Message: "Stock updated successfully",
	}

	json.NewEncoder(w).Encode(response)
}

// HealthCheck handles GET /health - returns service health status
func (h *ProductHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.Response{
		Success: true,
		Message: "Product service is healthy",
		Data: map[string]string{
			"service": "product-service",
			"status":  "UP",
		},
	}

	json.NewEncoder(w).Encode(response)
}

// sendErrorResponse sends a standardized error response
func (h *ProductHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)

	response := models.Response{
		Success: false,
		Error:   message,
	}

	json.NewEncoder(w).Encode(response)
}
