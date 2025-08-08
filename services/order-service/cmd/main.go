package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"order-service/internal/client"
	"order-service/internal/handlers"
	"order-service/internal/repository"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize repository
	orderRepo := repository.NewInMemoryOrderRepository()

	// Initialize service client for inter-service communication
	// In production, these URLs would come from service discovery or environment variables
	userServiceURL := "http://localhost:8081"
	productServiceURL := "http://localhost:8082"
	serviceClient := client.NewServiceClient(userServiceURL, productServiceURL)

	// Initialize handlers
	orderHandler := handlers.NewOrderHandler(orderRepo, serviceClient)

	// Setup routes
	router := setupRoutes(orderHandler)

	// Configure server
	server := &http.Server{
		Addr:         ":8083",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Println("🚀 Order Service starting on port 8083...")
		log.Println("📚 API Documentation:")
		log.Println("  POST  /orders              - Create order")
		log.Println("  GET   /orders/{id}         - Get order by ID")
		log.Println("  GET   /orders/user/{id}    - Get orders by user")
		log.Println("  PATCH /orders/{id}/status  - Update order status")
		log.Println("  GET   /orders              - List all orders")
		log.Println("  GET   /health              - Health check")
		log.Println("---")
		log.Printf("🔗 Connected to User Service: %s", userServiceURL)
		log.Printf("🔗 Connected to Product Service: %s", productServiceURL)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down Order Service...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	} else {
		log.Println("✅ Order Service shutdown complete")
	}
}

// setupRoutes configures all the HTTP routes
func setupRoutes(orderHandler *handlers.OrderHandler) *mux.Router {
	router := mux.NewRouter()

	// Add CORS middleware
	router.Use(corsMiddleware)
	
	// Add logging middleware
	router.Use(loggingMiddleware)

	// API routes
	api := router.PathPrefix("/").Subrouter()

	// Order routes
	api.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	api.HandleFunc("/orders", orderHandler.ListOrders).Methods("GET")
	api.HandleFunc("/orders/{id}", orderHandler.GetOrder).Methods("GET")
	api.HandleFunc("/orders/user/{user_id}", orderHandler.GetUserOrders).Methods("GET")
	api.HandleFunc("/orders/{id}/status", orderHandler.UpdateOrderStatus).Methods("PATCH")

	// Health check
	api.HandleFunc("/health", orderHandler.HealthCheck).Methods("GET")

	return router
}

// corsMiddleware adds CORS headers to responses
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
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
