# Go Microservices Common Issues & Solutions

This document covers the most common errors beginners face when working with this Go microservices project and their solutions.

## 🧭 Intern Diagnostic Framework (Read First)
When something breaks, avoid random fixes. Apply this sequence:
1. Classify Layer: (Env | Build | Runtime | Network | Data | Test).
2. Reproduce Minimally: Re-run single failing command (copy exact curl or build cmd).
3. Observe Evidence: Error text, exit code, logs, timestamp.
4. Form Single Hypothesis: One-sentence cause guess.
5. Test Fix: Smallest reversible change (config tweak, restart, code revert).
6. Confirm & Document: Did symptom disappear? Add note to personal log.

Prompt to Use if Stuck:
"Given this Go microservice error output: <PASTE>, classify which layer (Env/Build/Runtime/Network/Data/Test) and list top 3 likely root causes plus 1 verification command each."

## 🔧 Setup & Installation Issues

### Issue 1: Go Not Installed or Wrong Version
**Error Messages:**
```
command not found: go
go version go1.xx.x (need go1.21+)
```

**Solution:**
```bash
# Check current Go version
go version

# Install/Update Go on macOS
brew install go
# or brew upgrade go

# Install on Linux
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

# Add to PATH if not already added
export PATH=$PATH:/usr/local/go/bin
```

### Issue 2: GOPATH/GOROOT Issues
**Error Messages:**
```
cannot find package
$GOPATH not set
```

**Solution:**
```bash
# Check Go environment
go env GOPATH
go env GOROOT

# Modern Go (1.11+) uses modules, GOPATH is less important
# But if needed, set it:
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

### Issue 3: Module Download Issues
**Error Messages:**
```
go: cannot find main module
go mod download: connection refused
```

**Solution:**
```bash
# Navigate to service directory first
cd services/user-service

# Initialize module if go.mod missing
go mod init user-service

# Clean module cache if corrupted
go clean -modcache

# Download dependencies
go mod tidy
go mod download
```

## 🏗️ Build Issues

### Issue 4: Build Failures
**Error Messages:**
```
package xxx is not in GOROOT
undefined: SomeFunction
```

**Solution:**
```bash
# Check if you're in the right directory
pwd
# Should be in service directory, not root

# Clean and rebuild
go clean
go mod tidy
go build ./cmd/

# If still failing, check imports in Go files
# Make sure import paths match your module name
```

### Issue 5: Permission Denied on Scripts
**Error Messages:**
```
permission denied: ./scripts/build.sh
```

**Solution:**
```bash
# Make scripts executable
chmod +x scripts/*.sh

# Or individually
chmod +x scripts/build.sh
chmod +x scripts/run.sh
chmod +x scripts/stop.sh
chmod +x scripts/test.sh
```

## 🚀 Runtime Issues

### Issue 6: Port Already in Use
**Error Messages:**
```
bind: address already in use
listen tcp :8081: bind: address already in use
```

**Solution:**
```bash
# Find what's using the port
lsof -i :8081
# or
netstat -tulpn | grep :8081

# Kill the process
kill -9 <PID>

# Or use different ports by modifying the code
# Change the port in cmd/main.go:
# server := &http.Server{Addr: ":8084", ...}
```

### Issue 7: Service Communication Failures
**Error Messages:**
```
connection refused
no such host
dial tcp: lookup failed
```

**Solution:**
```bash
# Check if services are running
curl http://localhost:8081/health
curl http://localhost:8082/health

# Check service URLs in order-service
# In services/order-service/cmd/main.go, verify:
userServiceURL := "http://localhost:8081"
productServiceURL := "http://localhost:8082"

# For Docker, use service names:
userServiceURL := "http://user-service:8081"
```

### Issue 8: JSON Parsing Errors
**Error Messages:**
```
invalid character 'x' looking for beginning of value
EOF
```

**Solution:**
```bash
# Check your JSON payload format
# Valid example:
curl -X POST http://localhost:8081/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John","email":"john@example.com","password":"123456"}'

# Common mistakes:
# - Missing Content-Type header
# - Invalid JSON syntax (missing quotes, trailing commas)
# - Wrong field names (check models in internal/models/)
```

## 🐳 Docker Issues

### Issue 9: Docker Build Failures
**Error Messages:**
```
failed to solve with frontend dockerfile.v0
no such file or directory
```

**Solution:**
```bash
# Make sure you're building from the service directory
cd services/user-service
docker build -t user-service .

# Or from root with proper context
docker build -t user-service ./services/user-service/

# Check Dockerfile exists and is properly formatted
cat Dockerfile
```

### Issue 10: Docker Compose Issues
**Error Messages:**
```
service "xxx" failed to build
unhealthy
```

**Solution:**
```bash
# Check individual service build first
cd services/user-service && docker build -t user-service .

# Check logs
docker-compose logs user-service

# Force rebuild
docker-compose build --no-cache

# Check health endpoints manually
docker exec -it <container_id> wget -O- http://localhost:8081/health
```

## 🧪 Testing Issues

### Issue 11: API Tests Failing
**Error Messages:**
```
curl: (7) Failed to connect
404 Not Found
500 Internal Server Error
```

**Solution:**
```bash
# Check if services are actually running
ps aux | grep "main"

# Check service logs
tail -f logs/user\ service.log

# Verify endpoints exist
# Check internal/handlers/ files for correct routes

# Test with proper HTTP methods
curl -X GET http://localhost:8081/users     # List users
curl -X POST http://localhost:8081/users    # Create user
```

### Issue 12: CORS Issues (Frontend Integration)
**Error Messages:**
```
Access to fetch at 'http://localhost:8081' from origin 'http://localhost:3000' has been blocked by CORS policy
```

**Solution:**
The services already include CORS middleware, but if you face issues:

```go
// In cmd/main.go, check corsMiddleware function:
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}
```

## 📊 Debugging Tips

### General Debugging Approach

1. **Check Service Status:**
   ```bash
   curl http://localhost:8081/health
   curl http://localhost:8082/health
   curl http://localhost:8083/health
   ```

2. **Check Logs:**
   ```bash
   tail -f logs/*.log
   ```

3. **Verify Service Communication:**
   ```bash
   # From order service, test if it can reach others
   curl http://localhost:8081/users
   curl http://localhost:8082/products
   ```

4. **Test Individual Components:**
   ```bash
   # Test user creation
   curl -X POST http://localhost:8081/users \
     -H "Content-Type: application/json" \
     -d '{"name":"Test","email":"test@example.com","password":"123456"}'
   
   # Test product listing
   curl http://localhost:8082/products
   
   # Test order creation (need valid user and product IDs)
   curl -X POST http://localhost:8083/orders \
     -H "Content-Type: application/json" \
     -d '{"user_id":"USER_ID","items":[{"product_id":"PRODUCT_ID","quantity":1}]}'
   ```

### Environment-Specific Issues

**macOS:**
- Use `brew` for package management
- Check firewall settings if ports are blocked

**Linux:**
- Check if ports are available: `ss -tlnp | grep :808`
- Verify user permissions for file operations

**Windows:**
- Use Git Bash or WSL for shell scripts
- Check Windows Defender firewall

## 🔍 Layer Classification Cheat Sheet (Intern)
| Layer | Typical Symptoms | First Command |
|-------|------------------|---------------|
| Env | go not found, wrong version | `go version` |
| Build | cannot find module, undefined symbol | `go mod tidy` |
| Runtime | panic, bind error, nil pointer | `tail -n 50 logs/<svc>.log` |
| Network | connection refused, timeout | `curl -v http://localhost:PORT/health` |
| Data | not found, duplicate errors | Inspect repository logic / test data |
| Test | flaky, race warnings | `go test -race ./...` |

## 🧪 Race Condition Quick Check
Run with race detector when adding concurrency:
```bash
go test -race ./...
```
If race found, identify shared mutable state without lock; add `sync.RWMutex` or channel pattern.

## 🔁 Prompt Patterns for Troubleshooting
| Scenario | Prompt Template |
|----------|-----------------|
| Unknown error text | "Explain this Go error and typical root causes: <ERROR>" |
| Flaky test | "Suggest 3 hypotheses for intermittent failure in test <NAME> and verification steps." |
| Performance concern | "Profile strategy for endpoint <ENDPOINT> under load; which tools and why?" |

## 🛠️ Development Best Practices

### Code Organization
- Keep each service in its own directory
- Use consistent module naming
- Follow Go naming conventions

### Error Handling
- Always check errors in Go
- Use meaningful error messages
- Log errors with context

### Service Communication
- Use proper HTTP status codes
- Implement timeout handling
- Add retry logic for production

### Testing Strategy
- Write unit tests for business logic
- Test API endpoints individually
- Create integration tests for service communication

---

## Quick Reference Commands

```bash
# Start everything
./scripts/build.sh && ./scripts/run.sh

# Stop everything
./scripts/stop.sh

# Test everything
./scripts/test.sh

# Check service health
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health

# View logs
tail -f logs/*.log

# Docker alternative
docker-compose up -d
docker-compose logs -f
docker-compose down
```

## 🔁 Cross-Reference
For a curated subset aligned with evaluation criteria and AI-assisted refinement history, see `docs/TOOLKIT.md` Sections 6, 8, and 10.

## Getting Help

If you encounter issues not covered here:

1. Check the service logs first
2. Verify all dependencies are installed
3. Ensure you're following the exact steps in README.md
4. Test each service individually before testing integration
5. Use the debugging commands provided above

Remember: Most issues are due to incorrect setup, missing dependencies, or typos in configuration. Take your time and verify each step!

---
Intern-focused enhancements added; use prompt templates to accelerate accurate diagnosis.
