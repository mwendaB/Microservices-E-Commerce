#!/bin/bash

# Build script for Go microservices
echo "🔨 Building Go Microservices..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to build a service
build_service() {
    local service_name=$1
    local service_path=$2
    
    echo -e "${YELLOW}Building ${service_name}...${NC}"
    
    cd "$service_path" || {
        echo -e "${RED}❌ Failed to navigate to ${service_path}${NC}"
        return 1
    }
    
    # Download dependencies
    echo "📦 Downloading dependencies for ${service_name}..."
    go mod tidy
    go mod download
    
    # Build the service
    echo "🔧 Compiling ${service_name}..."
    go build -o bin/main ./cmd/
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✅ ${service_name} built successfully${NC}"
        cd - > /dev/null
        return 0
    else
        echo -e "${RED}❌ Failed to build ${service_name}${NC}"
        cd - > /dev/null
        return 1
    fi
}

# Build all services
echo "🚀 Starting build process..."

# Create bin directories
mkdir -p services/user-service/bin
mkdir -p services/product-service/bin
mkdir -p services/order-service/bin

# Build User Service
build_service "User Service" "services/user-service"
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Build failed for User Service${NC}"
    exit 1
fi

# Build Product Service
build_service "Product Service" "services/product-service"
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Build failed for Product Service${NC}"
    exit 1
fi

# Build Order Service
build_service "Order Service" "services/order-service"
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Build failed for Order Service${NC}"
    exit 1
fi

echo -e "${GREEN}🎉 All services built successfully!${NC}"
echo ""
echo "📁 Binaries are located in:"
echo "  • services/user-service/bin/main"
echo "  • services/product-service/bin/main"
echo "  • services/order-service/bin/main"
echo ""
echo "🚀 Run './scripts/run.sh' to start all services"
