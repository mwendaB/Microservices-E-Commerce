# Go Microservices Capstone Toolkit (Consolidated)

NOTE: The "Minimal Working Example" section is intentionally omitted per current instruction.

## 1. Title & Objective
Title: Prompt-Powered Kickstart: Building a Beginner’s Go Microservices E‑Commerce System
Objective: Provide a guided, reproducible learning toolkit showing how to design, run, test, and iterate on a basic microservices architecture in Go using AI-assisted workflows.
Scope:
- Three services (user, product, order) with independent modules
- RESTful APIs, in-memory storage, structured layering (handlers / repository / models)
- Scripts for build, run, test, integration verification

## 2. Technology Summary
Technology: Go (Golang) for microservices.
Why Go:
- Built-in concurrency (goroutines, channels)
- Batteries-included HTTP server
- Simple deployment (static binaries)
- Strong tooling (testing, profiling, formatting)
Microservices Pattern Highlighted:
- Service per business capability
- Independent build & deployment
- HTTP REST for synchronous communication
Real-World Parallel: Companies like Uber, Netflix, and Alibaba use similar decomposition patterns to scale teams and functionality independently.

## 3. System Requirements
See README for detailed commands.
- OS: macOS / Linux / Windows
- Go >= 1.21
- Git
- Docker & Docker Compose (optional container execution)
- curl / Postman (API testing)
Recommended VS Code Extensions: Go, Docker, REST Client, GitLens.
Environment Variables (future extensibility): `.env` sample in README.

## 4. Installation & Setup
Steps (summary):
1. Clone repo: `git clone <repo>`
2. Make scripts executable: `chmod +x scripts/*.sh`
3. Build services: `./scripts/build.sh`
4. Run services: `./scripts/run.sh`
5. Health checks: `curl localhost:8081/health` etc.
6. Integration test (automatic flow): use `scripts/test.sh` (now executes unit + API checks).
Troubleshooting: Refer to `docs/TROUBLESHOOTING.md`.

## 5. (Intentionally Skipped)
Minimal Working Example requirement excluded as per instruction.

## 6. AI Prompt Journal
Captured representative prompts actually applied during build & refinement.
| Day | Prompt (Condensed) | Goal | Action Taken | Outcome | Reflection |
|-----|--------------------|------|--------------|---------|------------|
| 1 | "Explain microservices vs monolith using my 3 services" | Concept clarity | Adjusted README architecture section | Clearer onboarding text | High value; accelerated documentation drafting |
| 1 | "Generate repository pattern for in-memory user store with concurrency safety" | Scaffold storage | Adopted RWMutex + interface | Working thread-safe repo | Prompt produced 90% of code; minor tweaks needed |
| 2 | "Add product filtering (category, price range, in-stock)" | Enhance product list | Implemented `ProductFilter` struct + filtering logic | Flexible querying works | Iterative prompt refinement produced concise solution |
| 3 | "How to structure inter-service calls from order service" | Cross-service logic | Created simple HTTP client wrapper | Order creation validates dependencies | Next step: add retries/timeouts |
| 3 | "Design consistent API response envelope" | Consistency | Unified success/error JSON format | Easier debugging & testing | Standardization reduced confusion |
| 4 | "Suggest table-driven tests for Go repository layer" | Testing | Added *_test.go files (see section 7) | Automated verification | AI helped enumerate negative cases quickly |
| 5 | "Summarize common setup issues for beginners" | Troubleshooting docs | Consolidated into TROUBLESHOOTING.md | Fewer repetitive setup questions | Good for learner empathy |
Evaluation: AI sped up boilerplate & documentation; required human judgment for naming, clarity, and test edge cases.

### Prompt Refinement Pattern
1. Start broad (scaffold) 2. Run code / identify gaps 3. Ask targeted improvement prompt 4. Integrate & refactor 5. Document decisions.

### Productivity Reflection
- Estimated time saved: 40–50% on boilerplate & docs.
- Main benefit: Rapid iteration of repository + handler scaffolds.
- Risk mitigated: Over-reliance avoided by manual validation & test creation.

## 7. Testing & Iteration
Testing Layers Implemented:
- Unit Tests: Repositories (user, product, order) verifying create, duplicate constraints, filtering, stock update, retrieval, negative cases.
- Scripted API Checks: `scripts/test.sh` exercises health endpoints + integration order flow.
- Integration Scenario: User creation -> product fetch -> order placement -> status check.
Iteration Notes:
- After first run, added duplicate email / product name validation tests.
- Added negative quantity/stock scenario tests (repository rejects invalid updates).
Peer Feedback (Placeholder):
- Feedback: "Need clearer indication of consistent JSON envelope" -> Action: Added explicit response format section in API docs.
- Feedback: "Add tests to prove repository logic" -> Action: Added table-driven tests in section 7.
Future Test Enhancements:
- Add HTTP handler tests with httptest
- Add contract tests for inter-service calls with mock HTTP servers
- Include load benchmarks (ab / k6) for critical endpoints

## 8. Common Issues & Fixes (Curated)
Selected high-impact entries (full list: TROUBLESHOOTING.md):
| Issue | Symptom | Fix |
|-------|---------|-----|
| Port in use | bind error | Free port or adjust server Addr |
| Duplicate user email | 409-like error message | Ensure unique email before Create |
| Product not found | 404-like error | Validate ID before order creation |
| Module not found | go build fails | Run `go mod tidy` inside service |
| CORS blocked | Browser reject | Confirm middleware present |
| Service comms fail | connection refused | Ensure target service up & correct URL |

### Recently Resolved (Execution / Orchestration)
| Issue | Symptom | Root Cause | Fix Applied |
|-------|---------|------------|------------|
| Order service health 000 during tests | test script marks order service down | Tests ran before service fully started | Added wait_for health polling + --no-monitor option to run.sh to exit promptly |
| Integration order creation failed | User & product OK, order create 400/connection error | Order service not yet listening | Same wait_for + shorter startup race window |
| bad substitution in run.sh | Script aborted at logs/${service_name,,}.log | Bash version without ${var,,} lowercase | Replaced with portable tr transformation |
| Missing order-service.pid | Monitor loop missing PID; reported service stopped | Previous naming with spaces vs dash & early exit | Standardized dashed log/pid names (user-service etc.) |
| Duplicate integration user email conflict | User create returned conflict or ambiguous | Reusing static email across runs | Generate unique timestamp-based email in test script |

## 9. References
Official:
- Go: https://go.dev/doc/
- gorilla/mux: https://github.com/gorilla/mux
- Effective Go: https://go.dev/doc/effective_go
- Go Testing: https://pkg.go.dev/testing
Supplemental:
- Microservices Patterns (Fowler): https://martinfowler.com/articles/microservices.html
- Twelve-Factor App: https://12factor.net/
- Docker Docs: https://docs.docker.com/
- UUID Library: https://github.com/google/uuid
Learning Aids:
- Prompt Engineering (internal prompt collection `docs/AI_LEARNING_PROMPTS.md`)

## 10. Evaluation Criteria Mapping
| Criterion | Evidence |
|-----------|----------|
| Clarity & completeness (30%) | This toolkit + README + API_EXAMPLES + Troubleshooting |
| GenAI usage (20%) | Section 6 journal & reflections |
| Functionality (20%) | Running 3-service system, integration script |
| Testing & iteration (20%) | Repository unit tests + test script + iteration notes |
| Creativity (10%) | Microservices e-commerce domain, consistent patterns, extensibility roadmap |

## 11. Next Steps (Roadmap)
Short Term: Add handler tests, structured logging fields (request_id), timeout context on outbound calls.
Mid Term: Introduce persistence (PostgreSQL), JWT auth, retry/circuit breaker pattern.
Long Term: Observability stack (Prometheus + OpenTelemetry), message broker for async workflows, gateway + rate limiting.

---
Consolidated toolkit complete (minimal example intentionally omitted).
