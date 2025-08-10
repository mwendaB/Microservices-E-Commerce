#!/bin/bash

# Run script for Go microservices
echo "🚀 Starting Go Microservices..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Check if binaries exist
check_binary() {
    local service_name=$1
    local binary_path=$2
    
    if [ ! -f "$binary_path" ]; then
        echo -e "${RED}❌ Binary not found for ${service_name}: ${binary_path}${NC}"
        echo -e "${YELLOW}💡 Run './scripts/build.sh' first to build all services${NC}"
        return 1
    fi
    return 0
}

# Function to start a service in background
start_service() {
    local service_name=$1
    local binary_path=$2
    local port=$3
    
    echo -e "${BLUE}Starting ${service_name} on port ${port}...${NC}"
    
    # Start service in background and capture PID
    local log_name=$(echo "$service_name" | tr '[:upper:]' '[:lower:]')
    $binary_path > "logs/${log_name}.log" 2>&1 &
    local pid=$!
    echo $pid > "logs/${log_name}.pid"
    
    # Wait a moment and check if process is still running
    sleep 2
    if kill -0 $pid 2>/dev/null; then
        echo -e "${GREEN}✅ ${service_name} started successfully (PID: ${pid})${NC}"
        return 0
    else
        echo -e "${RED}❌ Failed to start ${service_name}${NC}"
        return 1
    fi
}

# Create logs directory
mkdir -p logs

# Check all binaries exist
echo "🔍 Checking if services are built..."
check_binary "User Service" "services/user-service/bin/main" || exit 1
check_binary "Product Service" "services/product-service/bin/main" || exit 1
check_binary "Order Service" "services/order-service/bin/main" || exit 1

echo -e "${GREEN}✅ All binaries found${NC}"
echo ""

# Kill any existing processes
echo "🧹 Cleaning up any existing processes..."
pkill -f "user-service/bin/main" 2>/dev/null || true
pkill -f "product-service/bin/main" 2>/dev/null || true
pkill -f "order-service/bin/main" 2>/dev/null || true

# Wait a moment for processes to terminate
sleep 2

# Start all services
echo "🚀 Starting all services..."
echo ""

# Start User Service (port 8081)
start_service "User Service" "./services/user-service/bin/main" "8081"
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Failed to start User Service${NC}"
    exit 1
fi

# Start Product Service (port 8082)
start_service "Product Service" "./services/product-service/bin/main" "8082"
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Failed to start Product Service${NC}"
    exit 1
fi

# Wait a moment before starting Order Service (it depends on the other services)
sleep 3

# Start Order Service (port 8083)
start_service "Order Service" "./services/order-service/bin/main" "8083"
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Failed to start Order Service${NC}"
    exit 1
fi

echo ""
echo -e "${GREEN}🎉 All services are running!${NC}"
echo ""
echo "📊 Service Status:"
echo -e "${BLUE}  • User Service:    http://localhost:8081${NC}"
echo -e "${BLUE}  • Product Service: http://localhost:8082${NC}"
echo -e "${BLUE}  • Order Service:   http://localhost:8083${NC}"
echo ""
echo "📋 Quick Health Checks:"
echo "  curl http://localhost:8081/health"
echo "  curl http://localhost:8082/health"
echo "  curl http://localhost:8083/health"
echo ""
echo "📄 Logs are available in the 'logs/' directory"
echo "🛑 Run './scripts/stop.sh' to stop all services"
echo ""
echo -e "${YELLOW}💡 Press Ctrl+C to stop monitoring (services will continue running)${NC}"

# Monitor services (optional)
while true; do
    sleep 10
    
    # Check if all services are still running
    if ! kill -0 $(cat logs/user\ service.pid 2>/dev/null) 2>/dev/null; then
        echo -e "${RED}❌ User Service stopped unexpectedly${NC}"
        break
    fi
    
    if ! kill -0 $(cat logs/product\ service.pid 2>/dev/null) 2>/dev/null; then
        echo -e "${RED}❌ Product Service stopped unexpectedly${NC}"
        break
    fi
    
    if ! kill -0 $(cat logs/order\ service.pid 2>/dev/null) 2>/dev/null; then
        echo -e "${RED}❌ Order Service stopped unexpectedly${NC}"
        break
    fi
done
