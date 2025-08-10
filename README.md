# Go Microservices E-Commerce System

A beginner-friendly toolkit for building microservices using Go. This project demonstrates a simple e-commerce system with three core services.

> **Repository**: https://github.com/mwendaB/Microservices-E-Commerce  
> **Clone with**: `git clone https://github.com/mwendaB/Microservices-E-Commerce.git`

## � Table of Contents

- [Getting Started](#-getting-started)
- [Prerequisites](#-prerequisites)
- [Quick Start Guide](#-quick-start-guide)
- [Alternative Setup Methods](#️-alternative-setup-methods)
- [Development Environment](#-development-environment-setup)
- [API Documentation](#-api-documentation)
- [Project Structure](#️-project-structure)
- [Troubleshooting](#-troubleshooting)

## �🚀 Quick Start: What Are Microservices?

**Microservices** are a software architecture pattern where applications are built as a collection of small, independent services that communicate over well-defined APIs. Instead of building one large monolithic application, you create multiple smaller services that each handle a specific business function.

### Why Go for Microservices?

1. **Performance**: Go is compiled and fast, with excellent concurrency support
2. **Simplicity**: Clean syntax makes it easy to maintain multiple services
3. **Small Binaries**: Go produces small, self-contained executables
4. **Built-in HTTP Server**: Standard library includes everything for web services
5. **Great Tooling**: Excellent testing, profiling, and deployment tools

### Real-World Example
**Uber** uses Go microservices extensively. Their platform consists of hundreds of Go services handling everything from user authentication to ride matching and payment processing.

### Microservices vs Monolithic

| Aspect | Monolithic | Microservices |
|--------|------------|---------------|
| **Deployment** | Single unit | Independent services |
| **Scaling** | Scale entire app | Scale individual services |
| **Technology** | One tech stack | Different tech per service |
| **Failure** | App-wide failure | Isolated failures |
| **Team Size** | Large teams | Small, focused teams |

## � Getting Started

### 1. Clone the Repository

```bash
# Clone the project from GitHub
git clone https://github.com/mwendaB/Microservices-E-Commerce.git

# Navigate to the project directory
cd Microservices-E-Commerce
```

### 2. Verify Project Structure
After cloning, you should see:
```
Microservices-E-Commerce/
├── services/           # Microservices (user, product, order)
├── scripts/           # Build and run scripts
├── docs/             # Documentation
├── docker-compose.yml # Container orchestration
└── README.md         # This file
```

## �📋 Prerequisites

### System Requirements

#### Go Installation
- **Go 1.21+** (latest stable version recommended)
- **Git** for version control
- **Docker** for containerization (optional)
- **IDE**: VS Code with Go extension (recommended)

#### Installation Steps

**macOS:**
```bash
# Install Go using Homebrew
brew install go

# Verify installation
go version
```

**Linux (Ubuntu/Debian):**
```bash
# Download and install Go
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

# Add to PATH (add to ~/.bashrc or ~/.zshrc)
export PATH=$PATH:/usr/local/go/bin

# Verify installation
go version
```

**Windows:**
1. Download installer from https://golang.org/dl/
2. Run the installer
3. Verify with `go version` in Command Prompt

#### Essential Tools
```bash
# Install essential packages
go install github.com/gorilla/mux@latest
go install github.com/golang/mock/mockgen@latest

# Install Docker
# Visit https://docs.docker.com/get-docker/

# Install Postman for API testing
# Visit https://www.postman.com/downloads/
```

## 🏗️ Project Structure

```
E-Commerce/
├── services/
│   ├── user-service/
│   │   ├── cmd/main.go
│   │   ├── internal/
│   │   │   ├── handlers/
│   │   │   ├── models/
│   │   │   └── repository/
│   │   ├── Dockerfile
│   │   └── go.mod
│   ├── product-service/
│   │   ├── cmd/main.go
│   │   ├── internal/
│   │   │   ├── handlers/
│   │   │   ├── models/
│   │   │   └── repository/
│   │   ├── Dockerfile
│   │   └── go.mod
│   └── order-service/
│       ├── cmd/main.go
│       ├── internal/
│       │   ├── handlers/
│       │   ├── models/
│       │   ├── repository/
│       │   └── client/
│       ├── Dockerfile
│       └── go.mod
├── shared/
│   └── utils/
├── docker-compose.yml
├── scripts/
│   ├── build.sh
│   ├── run.sh
│   └── test.sh
├── docs/
│   ├── api/
│   └── setup/
└── README.md
```

## 🔧 Development Environment Setup

### VS Code Extensions (Recommended)
- **Go** (official Go extension)
- **Docker**
- **REST Client** (for API testing)
- **GitLens**

### Environment Variables Setup
Create a `.env` file in the root directory:
```env
# Service Ports
USER_SERVICE_PORT=8081
PRODUCT_SERVICE_PORT=8082
ORDER_SERVICE_PORT=8083

# Database URLs (for future expansion)
USER_DB_URL=memory
PRODUCT_DB_URL=memory
ORDER_DB_URL=memory

# API Keys (for future use)
JWT_SECRET=your-secret-key-here
```

## 🚀 Quick Start Guide

### 1. Set Up the Development Environment

```bash
# After cloning, navigate to the project directory
cd Microservices-E-Commerce

# Make scripts executable (Linux/macOS)
chmod +x scripts/*.sh

# Verify Go installation
go version
# Should show Go 1.21+ 
```

### 2. Install Dependencies and Build Services

```bash
# Build all services (this will also download dependencies)
./scripts/build.sh

# You should see:
# ✅ User Service built successfully
# ✅ Product Service built successfully  
# ✅ Order Service built successfully
```

### 3. Start All Services

```bash
# Start all microservices
./scripts/run.sh

# Services will start on:
# - User Service: http://localhost:8081
# - Product Service: http://localhost:8082
# - Order Service: http://localhost:8083
```

### 4. Verify Everything is Working

```bash
# Run the test suite
./scripts/test.sh

# Or test individual services
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
```

### 5. Test the Complete System

```bash
# Create a user
curl -X POST http://localhost:8081/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","password":"password123"}'

# List available products (comes with sample data)
curl http://localhost:8082/products

# Create an order (replace USER_ID and PRODUCT_ID with actual values)
curl -X POST http://localhost:8083/orders \
  -H "Content-Type: application/json" \
  -d '{"user_id":"USER_ID","product_id":"PRODUCT_ID","quantity":2}'
```

## 🛠️ Alternative Setup Methods

### Option 1: Manual Setup (if scripts don't work)

```bash
# Build each service individually
cd services/user-service && go build -o bin/main ./cmd/main.go && cd ../..
cd services/product-service && go build -o bin/main ./cmd/main.go && cd ../..
cd services/order-service && go build -o bin/main ./cmd/main.go && cd ../..

# Run each service in separate terminals
./services/user-service/bin/main &     # Terminal 1
./services/product-service/bin/main &  # Terminal 2  
./services/order-service/bin/main &    # Terminal 3
```

### Option 2: Using Docker (Alternative)

```bash
# Build and run with Docker Compose
docker-compose up --build

# Services will be available on the same ports
```

## 🔧 Development Environment Setup

### VS Code Extensions (Recommended)
- **Go** (official Go extension)
- **Docker**
- **REST Client** (for API testing)
- **GitLens**

### Environment Variables Setup
Create a `.env` file in the root directory:
```env
# Service Ports
USER_SERVICE_PORT=8081
PRODUCT_SERVICE_PORT=8082
ORDER_SERVICE_PORT=8083

# Database URLs (for future expansion)
USER_DB_URL=memory
PRODUCT_DB_URL=memory
ORDER_DB_URL=memory

# API Keys (for future use)
JWT_SECRET=your-secret-key-here
```

## 🧹 Managing Services

### Stop All Services
```bash
./scripts/stop.sh
```

### View Logs
```bash
# Service logs are saved in the logs/ directory
tail -f logs/userservice.log
tail -f logs/productservice.log
tail -f logs/orderservice.log
```

### Troubleshooting
If you encounter issues:
1. Check `docs/TROUBLESHOOTING.md` for common problems
2. Verify Go version: `go version` (needs 1.21+)
3. Check if ports are available: `lsof -i :8081,8082,8083`
4. Review service logs in the `logs/` directory

## 📚 API Documentation

### User Service (Port 8081)
- `POST /users` - Create user
- `GET /users/{id}` - Get user by ID
- `POST /auth/login` - User authentication
- `GET /health` - Health check

### Product Service (Port 8082)
- `GET /products` - List all products
- `GET /products/{id}` - Get product by ID
- `POST /products` - Create product (admin)
- `GET /health` - Health check

### Order Service (Port 8083)
- `POST /orders` - Create order
- `GET /orders/{id}` - Get order by ID
- `GET /orders/user/{user_id}` - Get user orders
- `GET /health` - Health check

## 🧪 Testing

### Unit Tests
```bash
# Run tests for all services
./scripts/test.sh

# Run tests for specific service
cd services/user-service && go test ./...
```

### Integration Testing
```bash
# Start all services
./scripts/run.sh

# Run integration tests
cd tests && go test -tags=integration ./...
```

## 🐳 Docker Deployment

### Single Service
```bash
# Build and run user service
cd services/user-service
docker build -t user-service .
docker run -p 8081:8081 user-service
```

### All Services with Docker Compose
```bash
# Start all services
docker-compose up

# Start in background
docker-compose up -d

# Stop services
docker-compose down
```

## 🔍 Common Issues & Solutions

### Issue 1: Port Already in Use
**Error**: `bind: address already in use`
**Solution**: 
```bash
# Find process using port
lsof -i :8081

# Kill process
kill -9 <PID>
```

### Issue 2: Module Not Found
**Error**: `package not found`
**Solution**:
```bash
# Run in service directory
go mod tidy
go mod download
```

### Issue 3: CORS Issues
**Error**: Frontend can't connect to API
**Solution**: Services include CORS headers by default

### Issue 4: Service Communication Fails
**Error**: Order service can't reach user/product services
**Solution**: Check service URLs in configuration

## 📈 Next Steps

After mastering this basic example:

1. **Database Integration**: Replace in-memory storage with PostgreSQL/MongoDB
2. **Authentication**: Implement JWT tokens and middleware
3. **Message Queues**: Add RabbitMQ or Kafka for async communication
4. **API Gateway**: Implement routing and load balancing
5. **Monitoring**: Add logging, metrics, and health checks
6. **CI/CD**: Set up automated testing and deployment
7. **Service Discovery**: Implement service registry (Consul, etcd)
8. **Caching**: Add Redis for improved performance

## 🤝 Contributing

This is a learning project. Feel free to:
- Add new features
- Improve documentation
- Create additional examples
- Fix bugs and issues

## 📄 License

This project is for educational purposes. Feel free to use and modify as needed.

---

## Troubleshooting Guide

### Go Environment Issues
```bash
# Check Go installation
go version
go env GOPATH
go env GOROOT

# Reset Go modules
go clean -modcache
```

### Service Startup Issues
```bash
# Check if ports are available
netstat -tulpn | grep :808

# Check logs
docker-compose logs user-service
```

### API Testing
Use the provided Postman collection in `docs/api/` or the REST Client files in each service directory.

---

**Happy Coding!** 🎉

This toolkit provides everything you need to understand and build Go microservices. Start with the User Service and gradually add complexity as you learn.
