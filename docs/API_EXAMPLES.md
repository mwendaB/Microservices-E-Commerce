# API Testing Examples

This document provides comprehensive examples for testing the Go microservices API endpoints.

## 🚀 Quick Start Testing

### 1. Health Checks
Test if all services are running:

```bash
# User Service
curl http://localhost:8081/health

# Product Service  
curl http://localhost:8082/health

# Order Service
curl http://localhost:8083/health
```

Expected response:
```json
{
  "success": true,
  "message": "User service is healthy",
  "data": {
    "service": "user-service",
    "status": "UP"
  }
}
```

## 👤 User Service API (Port 8081)

### Create User
```bash
curl -X POST http://localhost:8081/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com", 
    "password": "securepassword123"
  }'
```

Response:
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": "uuid-here",
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2025-01-XX...",
    "updated_at": "2025-01-XX..."
  }
}
```

### Get User by ID
```bash
# Replace USER_ID with actual ID from creation response
curl http://localhost:8081/users/USER_ID
```

### List All Users
```bash
curl http://localhost:8081/users
```

### User Login
```bash
curl -X POST http://localhost:8081/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword123"
  }'
```

Response:
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "user": {
      "id": "uuid-here",
      "name": "John Doe",
      "email": "john@example.com"
    },
    "token": "mock-jwt-token-uuid"
  }
}
```

## 📦 Product Service API (Port 8082)

### List All Products
```bash
curl http://localhost:8082/products
```

### Get Product by ID
```bash
# Replace PRODUCT_ID with actual ID from list response
curl http://localhost:8082/products/PRODUCT_ID
```

### Create Product
```bash
curl -X POST http://localhost:8082/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Gaming Laptop",
    "description": "High-performance gaming laptop with RTX graphics",
    "price": 1299.99,
    "category": "Electronics",
    "stock": 10,
    "image_url": "https://example.com/laptop.jpg"
  }'
```

### Update Product
```bash
curl -X PUT http://localhost:8082/products/PRODUCT_ID \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Gaming Laptop",
    "price": 1199.99,
    "stock": 15
  }'
```

### Update Product Stock
```bash
curl -X PATCH http://localhost:8082/products/PRODUCT_ID/stock \
  -H "Content-Type: application/json" \
  -d '{
    "stock": 25
  }'
```

### Get Products by Category
```bash
curl http://localhost:8082/products/category/Electronics
```

### Filter Products
```bash
# Filter by price range
curl "http://localhost:8082/products?min_price=100&max_price=500"

# Filter by category and in-stock items
curl "http://localhost:8082/products?category=Electronics&in_stock=true"

# Multiple filters
curl "http://localhost:8082/products?category=Footwear&min_price=50&max_price=200&in_stock=true"
```

## 🛒 Order Service API (Port 8083)

### Create Order
```bash
# First, get a valid user ID and product ID
USER_ID="user-id-from-user-service"
PRODUCT_ID="product-id-from-product-service"

curl -X POST http://localhost:8083/orders \
  -H "Content-Type: application/json" \
  -d "{
    \"user_id\": \"$USER_ID\",
    \"items\": [
      {
        \"product_id\": \"$PRODUCT_ID\",
        \"quantity\": 2
      }
    ]
  }"
```

### Get Order by ID
```bash
curl http://localhost:8083/orders/ORDER_ID
```

### Get Orders by User
```bash
curl http://localhost:8083/orders/user/USER_ID
```

### Update Order Status
```bash
curl -X PATCH http://localhost:8083/orders/ORDER_ID/status \
  -H "Content-Type: application/json" \
  -d '{
    "status": "confirmed"
  }'
```

Valid statuses: `pending`, `confirmed`, `shipped`, `delivered`, `cancelled`

### List All Orders
```bash
curl http://localhost:8083/orders
```

## 🔄 Complete Integration Test

Here's a complete workflow that tests all services working together:

```bash
#!/bin/bash

echo "🧪 Complete Integration Test"

# 1. Create a user
echo "1. Creating user..."
USER_RESPONSE=$(curl -s -X POST http://localhost:8081/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Integration Test User",
    "email": "integration@example.com",
    "password": "testpassword123"
  }')

USER_ID=$(echo $USER_RESPONSE | grep -o '"id":"[^"]*"' | cut -d'"' -f4)
echo "   User ID: $USER_ID"

# 2. Get available products
echo "2. Getting products..."
PRODUCTS_RESPONSE=$(curl -s http://localhost:8082/products)
PRODUCT_ID=$(echo $PRODUCTS_RESPONSE | grep -o '"id":"[^"]*"' | head -1 | cut -d'"' -f4)
echo "   Product ID: $PRODUCT_ID"

# 3. Create an order
echo "3. Creating order..."
ORDER_RESPONSE=$(curl -s -X POST http://localhost:8083/orders \
  -H "Content-Type: application/json" \
  -d "{
    \"user_id\": \"$USER_ID\",
    \"items\": [
      {
        \"product_id\": \"$PRODUCT_ID\",
        \"quantity\": 1
      }
    ]
  }")

ORDER_ID=$(echo $ORDER_RESPONSE | grep -o '"id":"[^"]*"' | cut -d'"' -f4)
echo "   Order ID: $ORDER_ID"

# 4. Update order status
echo "4. Updating order status..."
curl -s -X PATCH http://localhost:8083/orders/$ORDER_ID/status \
  -H "Content-Type: application/json" \
  -d '{"status": "confirmed"}' > /dev/null

# 5. Verify order
echo "5. Verifying order..."
curl -s http://localhost:8083/orders/$ORDER_ID | jq .

echo "✅ Integration test complete!"
```

## 🐛 Error Testing

### Test Invalid Requests

#### Invalid User Creation
```bash
# Missing required fields
curl -X POST http://localhost:8081/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John"}'  # Missing email and password
```

#### Invalid Product ID
```bash
curl http://localhost:8082/products/invalid-id
```

#### Invalid Order Creation
```bash
# Non-existent user
curl -X POST http://localhost:8083/orders \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "non-existent-user",
    "items": [{"product_id": "some-product", "quantity": 1}]
  }'
```

## 📋 Postman Collection

You can import these examples into Postman:

1. Create a new collection called "Go Microservices"
2. Add the following requests:

### User Service Collection
- GET Health Check: `http://localhost:8081/health`
- POST Create User: `http://localhost:8081/users`
- GET List Users: `http://localhost:8081/users`
- GET User by ID: `http://localhost:8081/users/{{user_id}}`
- POST Login: `http://localhost:8081/auth/login`

### Product Service Collection
- GET Health Check: `http://localhost:8082/health`
- GET List Products: `http://localhost:8082/products`
- GET Product by ID: `http://localhost:8082/products/{{product_id}}`
- POST Create Product: `http://localhost:8082/products`
- PUT Update Product: `http://localhost:8082/products/{{product_id}}`

### Order Service Collection
- GET Health Check: `http://localhost:8083/health`
- POST Create Order: `http://localhost:8083/orders`
- GET Order by ID: `http://localhost:8083/orders/{{order_id}}`
- GET User Orders: `http://localhost:8083/orders/user/{{user_id}}`
- PATCH Update Status: `http://localhost:8083/orders/{{order_id}}/status`

## 🔧 Advanced Testing

### Load Testing with Apache Bench
```bash
# Test user service health endpoint
ab -n 1000 -c 10 http://localhost:8081/health

# Test product listing
ab -n 500 -c 5 http://localhost:8082/products
```

### Using HTTPie (Alternative to curl)
```bash
# Install httpie: pip install httpie

# Create user
http POST localhost:8081/users name="Jane Doe" email="jane@example.com" password="password123"

# Get products
http GET localhost:8082/products

# Create order
http POST localhost:8083/orders user_id="USER_ID" items:='[{"product_id":"PRODUCT_ID","quantity":1}]'
```

### Testing with jq for JSON Processing
```bash
# Get all product names
curl -s http://localhost:8082/products | jq '.data[].name'

# Get user email from login response
curl -s -X POST http://localhost:8081/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"password123"}' | jq '.data.user.email'
```

---

## 🚨 Common Response Formats

### Success Response
```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": { /* actual data */ }
}
```

### Error Response
```json
{
  "success": false,
  "error": "Error message description"
}
```

### HTTP Status Codes Used
- `200` - OK (successful GET, PUT, PATCH)
- `201` - Created (successful POST)
- `400` - Bad Request (validation errors)
- `401` - Unauthorized (authentication failed)
- `404` - Not Found (resource doesn't exist)
- `409` - Conflict (duplicate resource)
- `500` - Internal Server Error (server issues)

Remember to check the logs if any test fails: `tail -f logs/*.log`
