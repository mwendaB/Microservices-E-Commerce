# API Testing Examples

This document provides comprehensive examples for testing the Go microservices API endpoints.

## 👩‍💻 How to Read These Examples (Intern Guide)
Each example shows: (a) The curl command, (b) Expected shape of response (envelope), (c) What to check if it fails.
If a command fails:
1. Re-run with `-v` for verbose output.
2. Hit the service `/health` endpoint.
3. Inspect the corresponding service log in `logs/`.
4. Validate JSON payload (use https://jqplay.org for quick syntax check).

Prompt Pattern to Extend Examples:
"Generate curl examples for updating a product including invalid JSON and not found edge cases using my existing response envelope." — Use results to enrich below.

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

### Create User (Beginner Notes)
Goal: Persist a new user in in-memory store. Validation: name/email/password required.
Failure Modes:
- Missing field -> error envelope with message
- Duplicate email -> conflict-style error envelope
- Malformed JSON -> JSON decode error
Recovery Checklist: Correct JSON, ensure unique email (add +timestamp), retry.

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

## 📦 Standard Response Envelope (Added)
All services return JSON in the following envelope for consistency.

Design Intent (Explain Like I'm New):
- success (bool): Fast path check; if false, client can branch immediately.
- message/error: Human friendly summary; `error` only present on failures.
- data: Actual payload; omit sensitive internals.
- details: Optional map for field-level or contextual error info.

When Adding New Endpoints:
1. Always wrap success path in envelope.
2. Provide specific, stable error messages (avoid leaking stack traces).
3. Include contextual IDs (order_id, user_id) where useful for client correlation.

Success:
```json
{
  "success": true,
  "message": "<human readable summary>",
  "data": { }
}
```
Error:
```json
{
  "success": false,
  "error": "<error message>",
  "details": { }
}
```
Refer to `docs/TOOLKIT.md` Section 6 & 7 for rationale and testing approach.

## 🧪 Extending Tests with Prompts (Intern Section)
Use AI to scaffold new negative test flows.
Example Prompt:
"List 5 edge case curl tests for order creation including invalid status transitions and explain expected HTTP codes using the existing envelope." — Add chosen tests to a personal scratch markdown then implement.

Track Added Tests Table:
| Date | Endpoint | New Case | Why Valuable | Implemented (Y/N) |
|------|----------|----------|--------------|-------------------|

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

## 🐛 Error Testing (How to Systematically Explore)
1. Start with valid request.
2. Remove one required field (expect validation error).
3. Corrupt JSON (missing quote) (expect decode error).
4. Use non-existent ID (expect not found error).
5. Repeat with high concurrency (later load testing).
Log each finding in your prompt/testing journal.

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

## 🧭 Troubleshooting Flow (If Any Example Fails)
| Step | Action | Purpose |
|------|--------|---------|
| 1 | Add `-v` to curl | See HTTP status + headers |
| 2 | Hit `/health` | Confirm service alive |
| 3 | Tail log | Spot panic or validation failure |
| 4 | Re-run with minimal JSON | Isolate payload issue |
| 5 | Compare vs working example | Spot drift |
| 6 | Form hypothesis & adjust | Intentional learning |

---
Intern additions complete; continue refining with your prompt log.
