package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"product-service/internal/handlers"
	"product-service/internal/repository"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize repository with sample data
	productRepo := repository.NewInMemoryProductRepository()

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productRepo)

	// Setup routes
	router := setupRoutes(productHandler)

	// Configure server
	server := &http.Server{
		Addr:         ":8082",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Println("🚀 Product Service starting on port 8082...")
		log.Println("📚 API Documentation:")
		log.Println("  GET  /products               - List all products")
		log.Println("  GET  /products/{id}          - Get product by ID")
		log.Println("  POST /products               - Create product")
		log.Println("  PUT  /products/{id}          - Update product")
		log.Println("  PATCH /products/{id}/stock   - Update stock")
		log.Println("  GET  /products/category/{cat} - Get by category")
		log.Println("  GET  /health                 - Health check")
		log.Println("---")
		log.Println("📦 Sample products loaded!")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down Product Service...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	} else {
		log.Println("✅ Product Service shutdown complete")
	}
}

// setupRoutes configures all the HTTP routes
func setupRoutes(productHandler *handlers.ProductHandler) *mux.Router {
	router := mux.NewRouter()

	// Add CORS middleware
	router.Use(corsMiddleware)
	
	// Add logging middleware
	router.Use(loggingMiddleware)

	// API routes
	api := router.PathPrefix("/").Subrouter()

	// Product routes
	api.HandleFunc("/products", productHandler.ListProducts).Methods("GET")
	api.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	api.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
	api.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	api.HandleFunc("/products/{id}/stock", productHandler.UpdateStock).Methods("PATCH")
	api.HandleFunc("/products/category/{category}", productHandler.GetProductsByCategory).Methods("GET")

	// Health check
	api.HandleFunc("/health", productHandler.HealthCheck).Methods("GET")

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
