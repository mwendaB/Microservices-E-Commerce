#!/bin/bash

# Test script for Go microservices
echo "🧪 Testing Go Microservices..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to run unit tests for a service
test_service() {
    local service_name=$1
    local service_path=$2
    
    echo -e "${YELLOW}Testing ${service_name}...${NC}"
    
    cd "$service_path" || {
        echo -e "${RED}❌ Failed to navigate to ${service_path}${NC}"
        return 1
    }
    
    # Run tests
    go test ./... -v
    local test_result=$?
    
    cd - > /dev/null
    
    if [ $test_result -eq 0 ]; then
        echo -e "${GREEN}✅ ${service_name} tests passed${NC}"
        return 0
    else
        echo -e "${RED}❌ ${service_name} tests failed${NC}"
        return 1
    fi
}

# Function to test API endpoints
test_api() {
    local service_name=$1
    local endpoint=$2
    local expected_status=$3
    
    echo -e "${BLUE}Testing ${service_name} API: ${endpoint}${NC}"
    
    # Make HTTP request and capture status code
    local status_code=$(curl -s -o /dev/null -w "%{http_code}" "$endpoint")
    
    if [ "$status_code" -eq "$expected_status" ]; then
        echo -e "${GREEN}✅ ${endpoint} returned ${status_code}${NC}"
        return 0
    else
        echo -e "${RED}❌ ${endpoint} returned ${status_code}, expected ${expected_status}${NC}"
        return 1
    fi
}

# Run unit tests for all services
echo "🔬 Running unit tests..."
echo ""

# Note: Since we don't have actual test files yet, we'll skip unit tests
# In a real project, you would have *_test.go files
echo -e "${YELLOW}📝 Unit tests not implemented yet (create *_test.go files)${NC}"
echo ""

# Test API endpoints if services are running
echo "🌐 Testing API endpoints..."
echo ""

# Check if services are running by testing health endpoints
api_test_passed=true

# Test User Service
if test_api "User Service" "http://localhost:8081/health" 200; then
    # Test additional User Service endpoints
    echo "  Testing User Service endpoints..."
    
    # Test creating a user
    echo "  Creating test user..."
    curl -s -X POST http://localhost:8081/users \
         -H "Content-Type: application/json" \
         -d '{"name":"Test User","email":"test@example.com","password":"password123"}' > /dev/null
    
    # Test getting users
    test_api "User Service" "http://localhost:8081/users" 200 || api_test_passed=false
else
    echo -e "${RED}❌ User Service is not running${NC}"
    api_test_passed=false
fi

echo ""

# Test Product Service
if test_api "Product Service" "http://localhost:8082/health" 200; then
    # Test additional Product Service endpoints
    echo "  Testing Product Service endpoints..."
    test_api "Product Service" "http://localhost:8082/products" 200 || api_test_passed=false
else
    echo -e "${RED}❌ Product Service is not running${NC}"
    api_test_passed=false
fi

echo ""

# Test Order Service
if test_api "Order Service" "http://localhost:8083/health" 200; then
    # Test additional Order Service endpoints
    echo "  Testing Order Service endpoints..."
    test_api "Order Service" "http://localhost:8083/orders" 200 || api_test_passed=false
else
    echo -e "${RED}❌ Order Service is not running${NC}"
    api_test_passed=false
fi

echo ""

# Integration test - Create an order
echo "🔗 Running integration test..."
echo "  Creating an order that requires all services..."

# First, create a user and get their ID
echo "  1. Creating user..."
user_response=$(curl -s -X POST http://localhost:8081/users \
                     -H "Content-Type: application/json" \
                     -d '{"name":"Integration Test User","email":"integration@example.com","password":"password123"}')

user_id=$(echo "$user_response" | grep -o '"id":"[^"]*"' | cut -d'"' -f4)

if [ -n "$user_id" ]; then
    echo -e "${GREEN}  ✅ User created with ID: ${user_id}${NC}"
    
    # Get a product ID
    echo "  2. Getting product..."
    products_response=$(curl -s http://localhost:8082/products)
    product_id=$(echo "$products_response" | grep -o '"id":"[^"]*"' | head -1 | cut -d'"' -f4)
    
    if [ -n "$product_id" ]; then
        echo -e "${GREEN}  ✅ Found product with ID: ${product_id}${NC}"
        
        # Create an order
        echo "  3. Creating order..."
        order_response=$(curl -s -X POST http://localhost:8083/orders \
                             -H "Content-Type: application/json" \
                             -d "{\"user_id\":\"${user_id}\",\"items\":[{\"product_id\":\"${product_id}\",\"quantity\":1}]}")
        
        order_id=$(echo "$order_response" | grep -o '"id":"[^"]*"' | cut -d'"' -f4)
        
        if [ -n "$order_id" ]; then
            echo -e "${GREEN}  ✅ Order created with ID: ${order_id}${NC}"
            echo -e "${GREEN}🎉 Integration test passed!${NC}"
        else
            echo -e "${RED}  ❌ Failed to create order${NC}"
            api_test_passed=false
        fi
    else
        echo -e "${RED}  ❌ No products found${NC}"
        api_test_passed=false
    fi
else
    echo -e "${RED}  ❌ Failed to create user${NC}"
    api_test_passed=false
fi

echo ""

# Summary
echo "📊 Test Summary:"
echo "---"
if [ "$api_test_passed" = true ]; then
    echo -e "${GREEN}✅ All API tests passed${NC}"
    echo -e "${GREEN}🎉 Microservices are working correctly!${NC}"
else
    echo -e "${RED}❌ Some API tests failed${NC}"
    echo -e "${YELLOW}💡 Make sure all services are running with './scripts/run.sh'${NC}"
fi

echo ""
echo "🔧 To run manual tests:"
echo "  • Use Postman with the provided collection"
echo "  • Use curl commands from the README.md"
echo "  • Check service logs in the 'logs/' directory"
