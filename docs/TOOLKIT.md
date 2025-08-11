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

### Intern Learning Journey (Guided Path)
This subsection is written as if mentoring an intern developer. Follow phases sequentially; each phase lists: Objective, Why It Matters, Example Prompt, What to Observe, Success Signal, Next Step.

Phase 1: Orientation & Fundamentals
- Objective: Understand the problem domain and structure.
- Why: Prevents writing code without a mental model.
- Example Prompt: "Describe the responsibilities of user, product, and order services in simple terms."
- Observe: Do you see clear boundaries? Any overlaps? Clarify before coding.
- Success Signal: You can explain each service in one sentence.
- Next Step: Start with data models.

Phase 2: Data Layer (Repositories)
- Objective: Implement in-memory repositories safely.
- Why: Foundation for handlers; concurrency mistakes early are costly.
- Example Prompt: "Show a thread-safe in-memory repository pattern for users in Go."
- Observe: Presence of sync.RWMutex; methods not exporting internal state.
- Success Signal: Tests can Create/Get/List/Delete without data races (run `go test -race`).
- Next Step: Expose HTTP endpoints.

Phase 3: API Layer (Handlers + Routing)
- Objective: Add HTTP handlers with clean separation from storage.
- Why: Keeps business logic testable outside HTTP concerns.
- Example Prompt: "Generate a basic gorilla/mux handler for creating a product with JSON input validation."
- Observe: Decoding errors handled; input validated; repository injected.
- Success Signal: curl POST returns 201 (or success envelope) with expected JSON.
- Next Step: Standardize responses.

Phase 4: Consistency & Contracts
- Objective: Introduce a uniform response envelope (data / error / meta).
- Why: Simplifies client error handling & logging.
- Example Prompt: "Suggest a JSON envelope structure for success and error responses for Go APIs."
- Observe: Error path includes machine-readable code and message.
- Success Signal: All handlers return the envelope; docs updated.
- Next Step: Cross-service communication.

Phase 5: Inter-Service Calls
- Objective: Order service validates user & product via HTTP.
- Why: Demonstrates service composition while preserving autonomy.
- Example Prompt: "Outline a simple Go HTTP client wrapper with retry for calling user-service." (We later added retry/backoff.)
- Observe: Context usage, error wrapping, retry with backoff.
- Success Signal: Order creation fails cleanly when user or product missing.
- Next Step: Testing depth.

Phase 6: Testing Strategy
- Objective: Cover core business logic with unit & integration tests.
- Why: Enables safe iteration and refactoring.
- Example Prompt: "Provide a table-driven test example for a repository Create method including duplicate case." 
- Observe: Table structure: name, input, expected error.
- Success Signal: All repo tests green; integration script passes end-to-end.
- Next Step: Operational robustness.

Phase 7: Orchestration & Reliability
- Objective: Harden startup + cross-service reliability.
- Why: Race conditions & flaky tests erode confidence.
- Example Prompt: "How to wait for multiple HTTP health endpoints before running integration tests in bash?"
- Observe: Poll loops with timeouts; clear exit codes; unique test data generation.
- Success Signal: Consistent green integration runs without manual delays.
- Next Step: Observability & structured logging (future roadmap).

Phase 8: Reflection & Documentation
- Objective: Capture decisions & lessons to accelerate new team member ramp-up.
- Why: Institutional memory reduces repeated mistakes.
- Example Prompt: "Summarize key pitfalls a beginner hits when starting Go microservices."
- Observe: Are docs actionable (steps) vs vague? Improve iteratively.
- Success Signal: A new intern can run system + tests in <15 minutes using only docs.

Intern Self-Check Questions
- Can I trace a request from HTTP handler to repository and back? (If not, diagram it.)
- Do I know where to add logging to diagnose a failing order creation? (Client wrapper + handler.)
- Can I intentionally break a test and explain the failure output? (If not, experiment.)

Prompt Crafting Tips for Interns
- Start with intent: "I want to..." (e.g., add filtering) then ask for structure, not full code.
- Ask for edge cases explicitly: "List failure modes for creating an order."
- Iterate: Provide current code snippet when asking for refinements to anchor context.

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

### Detailed Issue Guides (Beginner Friendly)
Below each guide: What it Means, How to Recognize, Why it Happens (Go / OS angle), Step-by-Step Fix, Prevention Checklist.

1. Port in Use
- Meaning: Your service tries to listen on a TCP port already taken by another process.
- Recognize: Error like: `listen tcp :8081: bind: address already in use` during `go run` or binary start.
- Why: Previous instance not stopped OR another app (e.g., AirPlay, Docker) occupies that port.
- Fix Steps: (a) Identify process: `lsof -i :8081` (b) Stop that process (c) Re-run service.
- Prevention: Use `stop.sh` before restart; document port map; consider making port configurable via env var.

2. Duplicate User Email
- Meaning: Attempt to create a second user with same unique email.
- Recognize: Our API returns standardized error envelope referencing duplicate.
- Why: Repository checks existing users by email to maintain uniqueness.
- Fix Steps: Use a new email (integration script now generates timestamp). For manual tests, append +tag in Gmail style.
- Prevention: Add uniqueness index when moving to a real DB; include constraint tests.

3. Product Not Found
- Meaning: Order references a product ID not stored.
- Recognize: 404-like error envelope from product-service or order-service validation failure.
- Why: Client passes stale/typo ID; test data not seeded.
- Fix Steps: List products first (`GET /products`); copy `id` exactly; re-run order creation.
- Prevention: Centralize creation of test fixtures; validate IDs at client side earlier.

4. Module Not Found (go build fails)
- Meaning: Go compiler cannot resolve an imported module.
- Recognize: Errors like `cannot find module providing package ...`.
- Why: New dependency added in code but `go mod tidy` not run; or editing wrong service folder.
- Fix Steps: cd into service root (e.g., `services/user-service`) then run `go mod tidy` and rebuild.
- Prevention: Run tidy after adding imports; keep each service's `go.mod` isolated.

5. CORS Blocked
- Meaning: Browser prevents frontend from calling API due to missing CORS headers.
- Recognize: Browser console: `CORS policy: No 'Access-Control-Allow-Origin' header`.
- Why: Missing CORS middleware configuration.
- Fix Steps: Add middleware setting `Access-Control-Allow-Origin: *` (or restricted domain), methods, headers.
- Prevention: Define a shared CORS function and apply consistently across services.

6. Service Communication Fail (Connection Refused)
- Meaning: One service calls another before target is listening or wrong host/port.
- Recognize: `dial tcp 127.0.0.1:8082: connect: connection refused` in logs.
- Why: Startup race (caller faster), incorrect base URL, service crashed.
- Fix Steps: Use health polling (implemented); verify env base URL; inspect target logs.
- Prevention: Always add a readiness/health check poll before issuing dependent calls in tests.

7. Order Service Health 000 / Early Test Failure
- Meaning: Integration script tests before order-service responds.
- Recognize: Test script logs status 000 or curl failure.
- Why: Parallel startup; no wait mechanism originally.
- Fix Steps: Added `wait_for` loop (retries with sleep) and `--no-monitor` to allow earlier handoff.
- Prevention: Treat startup ordering explicitly; design scripts idempotent.

8. Bash 'bad substitution' in run.sh
- Meaning: Shell rejects `${var,,}` lowercase operator.
- Recognize: `run.sh: line X: ${service_name,,}: bad substitution`.
- Why: Using a bash-specific feature under a shell not supporting that expansion.
- Fix Steps: Replace with portable pipeline: `echo "$service_name" | tr '[:upper:]' '[:lower:]'`.
- Prevention: Keep scripts POSIX-friendly unless shebang explicitly requires bash.

9. Missing order-service.pid / False Down Status
- Meaning: Monitor script cannot find PID file, assumes crash.
- Recognize: Logs show service running but monitor prints "stopped".
- Why: Inconsistent naming (spaces vs dashes) or race writing PID.
- Fix Steps: Standardize filenames (`order-service.pid`); write PID atomically.
- Prevention: Use consistent slug function for names; verify existence in tests.

10. Duplicate Integration User Email (Flaky Tests)
- Meaning: Repeated test runs collide on same email.
- Recognize: Second run returns conflict; first succeeded.
- Why: Hard-coded email in script.
- Fix Steps: Generate timestamp or UUID email each run (implemented).
- Prevention: Always randomize unique fields in integration tests.

Issue Diagnosis Mental Model for Interns
- Reproduce: Run the exact failing command again to confirm consistency.
- Isolate Layer: Network (ports), Service code (logs), Data (repository state), Script (parameters).
- Inspect Evidence: Logs + exit codes + curl verbose (-v) if needed.
- Form Hypothesis: Single sentence ("Order service not yet bound to port 8083").
- Apply Fix & Re-test: Did the root symptom vanish? If yes, document; if not, refine hypothesis.

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
Short Term: Structured logging fields (request_id), timeout context on outbound calls, expand handler test coverage for edge cases (some base handler tests already added; deepen scenarios).
Mid Term: Introduce persistence (PostgreSQL), JWT auth, retry/circuit breaker pattern.
Long Term: Observability stack (Prometheus + OpenTelemetry), message broker for async workflows, gateway + rate limiting.

---
Consolidated toolkit complete (minimal example intentionally omitted).
