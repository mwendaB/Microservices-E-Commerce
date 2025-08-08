package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-service/internal/handlers"
	"user-service/internal/repository"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize repository
	userRepo := repository.NewInMemoryUserRepository()

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userRepo)

	// Setup routes
	router := setupRoutes(userHandler)

	// Configure server
	server := &http.Server{
		Addr:         ":8081",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Println("🚀 User Service starting on port 8081...")
		log.Println("📚 API Documentation:")
		log.Println("  POST /users           - Create user")
		log.Println("  GET  /users/{id}      - Get user by ID")
		log.Println("  GET  /users           - List all users")
		log.Println("  POST /auth/login      - User login")
		log.Println("  GET  /health          - Health check")
		log.Println("---")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down User Service...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	} else {
		log.Println("✅ User Service shutdown complete")
	}
}

// setupRoutes configures all the HTTP routes
func setupRoutes(userHandler *handlers.UserHandler) *mux.Router {
	router := mux.NewRouter()

	// Add CORS middleware
	router.Use(corsMiddleware)
	
	// Add logging middleware
	router.Use(loggingMiddleware)

	// API routes
	api := router.PathPrefix("/").Subrouter()

	// User routes
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users", userHandler.ListUsers).Methods("GET")

	// Auth routes
	api.HandleFunc("/auth/login", userHandler.Login).Methods("POST")

	// Health check
	api.HandleFunc("/health", userHandler.HealthCheck).Methods("GET")

	return router
}

// corsMiddleware adds CORS headers to responses
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// loggingMiddleware logs HTTP requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the request
		log.Printf(
			"[%s] %s %s %v",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}
