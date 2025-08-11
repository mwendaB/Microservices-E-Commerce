#!/bin/bash

# Stop script for Go microservices
echo "🛑 Stopping Go Microservices..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to stop a service
stop_service() {
    local service_name=$1
    local pid_file=$2
    
    if [ -f "$pid_file" ]; then
        local pid=$(cat "$pid_file")
        if kill -0 "$pid" 2>/dev/null; then
            echo -e "${YELLOW}Stopping ${service_name} (PID: ${pid})...${NC}"
            kill "$pid"
            
            # Wait for graceful shutdown
            local count=0
            while kill -0 "$pid" 2>/dev/null && [ $count -lt 10 ]; do
                sleep 1
                count=$((count + 1))
            done
            
            # Force kill if still running
            if kill -0 "$pid" 2>/dev/null; then
                echo -e "${RED}Force killing ${service_name}...${NC}"
                kill -9 "$pid"
            fi
            
            echo -e "${GREEN}✅ ${service_name} stopped${NC}"
        else
            echo -e "${YELLOW}${service_name} was not running${NC}"
        fi
        rm -f "$pid_file"
    else
        echo -e "${YELLOW}No PID file found for ${service_name}${NC}"
    fi
}

# Stop all services
stop_service "User Service" "logs/user-service.pid"
stop_service "Product Service" "logs/product-service.pid"
stop_service "Order Service" "logs/order-service.pid"

# Also kill any processes that might be running without PID files
echo "🧹 Cleaning up any remaining processes..."
pkill -f "user-service/bin/main" 2>/dev/null || true
pkill -f "product-service/bin/main" 2>/dev/null || true
pkill -f "order-service/bin/main" 2>/dev/null || true

echo ""
echo -e "${GREEN}🎉 All services stopped successfully!${NC}"
echo ""
echo "📄 Logs are still available in the 'logs/' directory"
