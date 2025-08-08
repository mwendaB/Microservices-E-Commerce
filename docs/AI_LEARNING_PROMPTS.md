# Go Microservices Learning Prompt Collection

This document contains the complete set of prompts you can use throughout your 5-day capstone project to maximize your learning with AI assistance.

## 📋 Master Prompt Categories

### Day 1: Foundation & Setup
### Day 2: Service Implementation  
### Day 3: Service Communication
### Day 4: Integration & Testing
### Day 5: Documentation & Refinement

---

## 🚀 Day 1: Foundation & Setup

### Initial Setup Prompt
```
I'm starting a Go microservices project for my capstone. I have the basic structure but need help with:

1. Setting up my development environment properly
2. Understanding the project structure I've been given
3. Troubleshooting any installation issues

Current issue: [DESCRIBE YOUR SPECIFIC ISSUE]

My environment:
- OS: [macOS/Linux/Windows]
- Go version: [run 'go version']
- Current error: [paste error message if any]

Please provide step-by-step guidance to resolve this and verify my setup is correct.
```

### Understanding Microservices Prompt
```
I'm new to microservices and need clarification on:

1. How do the three services (user, product, order) interact?
2. What makes this architecture "microservices" vs monolithic?
3. Why is Go a good choice for this pattern?
4. What are the main benefits and challenges I should be aware of?

Please explain with concrete examples from my codebase.
```

### Environment Troubleshooting Prompt
```
I'm having issues with my Go environment setup:

Error message: [PASTE ERROR HERE]

What I've tried:
- [List what you've already attempted]

My current setup:
- go env GOPATH: [output]
- go env GOROOT: [output]  
- echo $PATH: [relevant parts]

Please help me diagnose and fix this issue step by step.
```

---

## 🔧 Day 2: Service Implementation

### User Service Deep Dive Prompt
```
I want to understand and enhance the User Service. Help me with:

1. Explain the repository pattern used in user-service/internal/repository/
2. How does the handler layer work in user-service/internal/handlers/?
3. I want to add [SPECIFIC FEATURE] to the user service
4. How can I improve error handling and validation?

Current code I'm working on: [PASTE RELEVANT CODE SECTION]

Please show me how to implement this properly with Go best practices.
```

### Product Service Enhancement Prompt
```
I need to modify the Product Service:

Goal: [DESCRIBE WHAT YOU WANT TO ADD/CHANGE]

Current issue: [DESCRIBE PROBLEM IF ANY]

Specific questions:
1. How do I add a new field to the Product model?
2. How should I handle product categories better?
3. How can I add search functionality?
4. What's the best way to handle product images?

Show me the exact code changes needed with explanations.
```

### Order Service Logic Prompt
```
The Order Service is the most complex. I need help understanding:

1. How does it communicate with other services?
2. The client package - how does HTTP communication work?
3. I want to add [SPECIFIC BUSINESS LOGIC]
4. How should I handle inventory management when orders are created?

Current challenge: [DESCRIBE SPECIFIC ISSUE]

Please explain the flow and show me how to implement my enhancement.
```

---

## 🔗 Day 3: Service Communication

### Inter-Service Communication Prompt
```
I need to understand and improve service-to-service communication:

Current situation:
- Order service calls User and Product services
- Sometimes getting [SPECIFIC ERROR]

Questions:
1. What are the different ways services can communicate?
2. How do I handle timeouts and failures gracefully?
3. Should I use HTTP REST, gRPC, or message queues?
4. How do I implement circuit breakers or retry logic?

Show me how to make the communication more robust.
```

### Error Handling Across Services Prompt
```
I'm getting errors when services try to communicate:

Error: [PASTE ERROR MESSAGE]

Service flow:
1. [Describe the request flow]
2. [Where it fails]

I need help with:
1. Proper error propagation between services
2. How to handle partial failures
3. Implementing timeouts and retries
4. Making the system more resilient

Please show me the code patterns to fix this.
```

### API Design Consistency Prompt
```
I want to make my APIs more consistent across all services:

Current issues I notice:
1. [Inconsistent response formats]
2. [Different error handling approaches]  
3. [Inconsistent naming conventions]

Help me:
1. Design a standard API response format
2. Implement consistent error handling
3. Ensure proper HTTP status codes
4. Add request/response logging

Show me how to refactor the existing code to be more consistent.
```

---

## 🧪 Day 4: Integration & Testing

### Testing Strategy Prompt
```
I need to implement comprehensive testing for my microservices:

Current testing gaps:
1. No unit tests yet
2. Need integration tests
3. Want to test service communication
4. How to test error scenarios

Help me create:
1. Unit tests for [SPECIFIC SERVICE/COMPONENT]
2. Integration tests that test multiple services
3. API endpoint testing strategy
4. How to mock external dependencies

Show me examples with the Go testing framework.
```

### Docker & Deployment Prompt
```
I want to containerize and deploy my services:

Current status:
- Have basic Dockerfiles
- Docker-compose setup exists
- [CURRENT ISSUE OR GOAL]

I need help with:
1. Optimizing Docker builds
2. Handling service discovery in containers
3. Managing environment variables
4. Setting up proper logging
5. Health checks and monitoring

Walk me through the deployment process and best practices.
```

### Load Testing & Performance Prompt
```
I want to test how my microservices perform under load:

Goals:
1. Test individual service performance
2. Test the complete order flow under load
3. Identify bottlenecks
4. Implement basic monitoring

Tools available: [curl, ab, etc.]

Show me how to:
1. Create realistic load tests
2. Measure and interpret results
3. Identify performance issues
4. Implement basic improvements
```

---

## 📚 Day 5: Documentation & Refinement

### Code Review & Improvement Prompt
```
I want to review and improve my codebase for the final submission:

Please review my [SPECIFIC FILE/COMPONENT] and help me:

1. Identify code quality issues
2. Suggest Go best practices I'm missing
3. Improve error handling and logging
4. Make the code more maintainable
5. Add proper comments and documentation

Code to review: [PASTE CODE SECTION]

Provide specific, actionable feedback with examples.
```

### API Documentation Prompt
```
I need to create comprehensive API documentation:

Current status:
- Have basic README
- APIs are working
- Need formal documentation

Help me create:
1. OpenAPI/Swagger specification
2. Clear endpoint documentation with examples
3. Error response documentation  
4. Authentication and authorization docs (if applicable)
5. Integration guide for other developers

Show me the best practices for API documentation.
```

### Deployment Guide Prompt
```
I need to create a deployment guide for beginners:

Target audience: Developers new to microservices

Cover:
1. Prerequisites and dependencies
2. Step-by-step installation guide
3. Configuration options
4. Troubleshooting common issues
5. Monitoring and maintenance

Current deployment methods:
- Local development
- Docker containers
- [Any cloud deployment]

Help me create clear, comprehensive documentation.
```

### Common Issues Documentation Prompt
```
Based on my experience building this project, help me document common issues:

Issues I encountered:
1. [LIST YOUR SPECIFIC ISSUES]
2. [HOW YOU SOLVED THEM]

Help me create:
1. Troubleshooting guide with solutions
2. FAQ section
3. Common error messages and fixes
4. Best practices and gotchas
5. Tips for beginners

This will help other students who use this project.
```

---

## 🔄 Continuous Learning Prompts

### Architecture Deep Dive Prompt
```
Now that I have a working system, help me understand advanced concepts:

1. How would this scale to 10+ services?
2. What about database per service pattern?
3. How do I handle distributed transactions?
4. What about service mesh and API gateways?
5. How do I implement proper observability?

Relate your explanations to my current codebase.
```

### Security Enhancement Prompt
```
I want to add security to my microservices:

Current state: Basic authentication simulation

Help me implement:
1. JWT token authentication
2. API rate limiting
3. Input validation and sanitization
4. HTTPS/TLS setup
5. Security headers

Show me step-by-step how to enhance my existing code.
```

### Production Readiness Prompt
```
How do I make my microservices production-ready?

Current limitations I see:
- In-memory storage
- Basic error handling
- No monitoring

Help me understand:
1. Database integration strategies
2. Logging and monitoring setup
3. Configuration management
4. Graceful shutdown handling
5. Health checks and readiness probes

Prioritize the most important improvements.
```

---

## 🎯 Specific Problem-Solving Prompts

### Debug Specific Error Prompt
```
I'm getting this specific error:

Error message: [EXACT ERROR]

Context:
- What I was trying to do: [DESCRIBE ACTION]
- When it happens: [DESCRIBE TIMING]
- Service involved: [WHICH SERVICE]
- Request details: [CURL COMMAND OR REQUEST INFO]

Code that might be relevant: [PASTE CODE]

Please help me debug this step by step.
```

### Performance Issue Prompt
```
My service is running slowly:

Performance issue:
- Service: [WHICH SERVICE]
- Endpoint: [WHICH ENDPOINT]  
- Response time: [ACTUAL TIME]
- Expected time: [DESIRED TIME]

I suspect: [YOUR THEORY]

Help me:
1. Profile and identify the bottleneck
2. Optimize the slow parts
3. Add proper benchmarking
4. Verify the improvements

Show me the tools and techniques to use.
```

### Feature Addition Prompt
```
I want to add a new feature:

Feature description: [DETAILED DESCRIPTION]

Questions:
1. Which service should handle this?
2. What new endpoints do I need?
3. How does this affect existing code?
4. What validation and error handling is needed?
5. How should I test this feature?

Please design this feature with me and show the implementation.
```

---

## 💡 Best Practices for Using These Prompts

### 1. Be Specific
- Always include exact error messages
- Paste relevant code sections
- Describe your environment and current state

### 2. Show Your Work
- Mention what you've already tried
- Explain your current understanding
- Ask for clarification on specific concepts

### 3. Ask for Examples
- Request code examples with explanations
- Ask for step-by-step guidance
- Want to see best practices in action

### 4. Follow Up
- Ask clarifying questions
- Request additional examples
- Validate your understanding

### 5. Document Learning
- Save the responses for your documentation
- Create your own examples based on the guidance
- Share knowledge with peers

---

## 🎓 Reflection Prompts for Learning

### Understanding Check Prompt
```
I've implemented [SPECIFIC FEATURE/CONCEPT]. Test my understanding:

My explanation: [EXPLAIN IN YOUR OWN WORDS]

Questions to verify my understanding:
1. [ASK SPECIFIC QUESTIONS]
2. [ABOUT THE IMPLEMENTATION]
3. [OR THE CONCEPTS INVOLVED]

Please correct any misconceptions and fill in gaps in my knowledge.
```

### Knowledge Gap Identification Prompt
```
As I work on this project, I realize I don't fully understand:

1. [CONCEPT 1]
2. [CONCEPT 2]  
3. [CONCEPT 3]

For each concept:
- Why is it important for microservices?
- How does it apply to my current project?
- What should I learn next to improve?

Help me prioritize my learning and suggest resources.
```

This collection gives you a comprehensive toolkit for AI-assisted learning throughout your capstone project. Use these prompts as starting points and modify them based on your specific needs and challenges!
