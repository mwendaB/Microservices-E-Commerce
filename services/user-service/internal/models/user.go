package models

import (
	"time"
	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"` // omitempty prevents password from being returned in JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest represents the request payload for creating a user
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// LoginRequest represents the request payload for user login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse represents the response for successful login
type LoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

// NewUser creates a new user with generated ID and timestamps
func NewUser(name, email, password string) *User {
	now := time.Now()
	return &User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Password:  password, // In production, this should be hashed
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
