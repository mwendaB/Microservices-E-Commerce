package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"user-service/internal/models"
	"user-service/internal/repository"

	"github.com/gorilla/mux"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	repo repository.UserRepository
}

// NewUserHandler creates a new user handler
func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

// CreateUser handles POST /users - creates a new user
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// Basic validation
	if req.Name == "" || req.Email == "" || req.Password == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Name, email, and password are required")
		return
	}

	// Create user
	user := models.NewUser(req.Name, req.Email, req.Password)
	if err := h.repo.Create(user); err != nil {
		log.Printf("Error creating user: %v", err)
		h.sendErrorResponse(w, http.StatusConflict, err.Error())
		return
	}

	// Remove password from response
	user.Password = ""

	response := models.Response{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetUser handles GET /users/{id} - retrieves a user by ID
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userID := vars["id"]

	if userID == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "User ID is required")
		return
	}

	user, err := h.repo.GetByID(userID)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		h.sendErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	response := models.Response{
		Success: true,
		Data:    user,
	}

	json.NewEncoder(w).Encode(response)
}

// Login handles POST /auth/login - authenticates a user
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// Basic validation
	if req.Email == "" || req.Password == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Email and password are required")
		return
	}

	// Get user by email
	user, err := h.repo.GetByEmail(req.Email)
	if err != nil {
		log.Printf("Login attempt for non-existent user: %s", req.Email)
		h.sendErrorResponse(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Simple password check (in production, use proper password hashing)
	if user.Password != req.Password {
		log.Printf("Invalid password for user: %s", req.Email)
		h.sendErrorResponse(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Create login response (in production, generate JWT token)
	loginResp := models.LoginResponse{
		User:  *user,
		Token: "mock-jwt-token-" + user.ID, // Mock token for demonstration
	}
	loginResp.User.Password = "" // Don't return password

	response := models.Response{
		Success: true,
		Message: "Login successful",
		Data:    loginResp,
	}

	json.NewEncoder(w).Encode(response)
}

// ListUsers handles GET /users - retrieves all users
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.repo.List()
	if err != nil {
		log.Printf("Error listing users: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}

	response := models.Response{
		Success: true,
		Data:    users,
	}

	json.NewEncoder(w).Encode(response)
}

// HealthCheck handles GET /health - returns service health status
func (h *UserHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.Response{
		Success: true,
		Message: "User service is healthy",
		Data: map[string]string{
			"service": "user-service",
			"status":  "UP",
		},
	}

	json.NewEncoder(w).Encode(response)
}

// sendErrorResponse sends a standardized error response
func (h *UserHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)

	response := models.Response{
		Success: false,
		Error:   message,
	}

	json.NewEncoder(w).Encode(response)
}
