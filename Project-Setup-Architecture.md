# Project Setup & Architecture: Zambone-One CLI

## 1. Executive Summary

**Project Name:** Zambone-One CLI (zambone-cli)

**Mission:** To build a "Golden Path" CLI tool that automates the creation, provisioning, and monitoring of macro services.

**Methodology:** strict adherence to Test-Driven Development (TDD) and Clean Architecture (Ports & Adapters) to ensure testability and decoupling of infrastructure concerns.

## 2. Technical Stack & Architecture Decisions

### 2.1 Core Application Stack

**Language:** Go (Golang) 1.25+

**Architecture:** Hexagonal / Clean Architecture

**Testing:** Go testing package + testify (Assertions) + mockery (Mock generation)

**CLI Framework:** Cobra (Acts as the Primary Driver/Adapter)

### 2.2 The "Zero-Cost" Local Simulation Environment

| Domain        | Local Replacement        |
|---------------|--------------------------|
| Cloud API     | LocalStack               |
| Orchestration | Kind                     |
| Observability | OpenTelemetry + Jaeger   |
| FinOps        | Infracost                |

## 3. Repository Structure (Clean Architecture)

We strictly separate Business Logic (Core) from Implementation Details (Adapters).

```
/
├── cmd/
│   └── zambone/
│       └── main.go            # Composition Root: Dependency Injection wire-up
│
├── internal/
│   ├── core/                  # THE DOMAIN (Pure Go, No external deps)
│   │   ├── domain/            # Entities (Structs: Service, Budget, Config)
│   │   ├── ports/             # Interfaces (Input/Output definitions)
│   │   └── services/          # Use Cases (The application logic)
│   │
│   └── adapters/              # THE INFRASTRUCTURE (External implementations)
│       ├── cli/               # Cobra commands (The Entrypoint)
│       ├── filesystem/        # Disk I/O implementation of FS Port
│       ├── aws/               # AWS SDK v2 implementation of Cloud Port
│       └── k8s/               # Client-Go implementation of Orchestrator Port
│
├── templates/                 # Embedded assets (Rails boilerplate, Terraform)
├── deploy/                    # LocalStack/Kind configurations
└── Makefile                   # Build & Test automation
```

## 4. Development Standards (The "Paved Way")

### 4.1 TDD Workflow (Mandatory)

Every feature must follow the Red-Green-Refactor cycle:

1. **Red:** Write a test for a Use Case in `internal/core/services`. Mock all dependencies (FS, AWS) using the interfaces defined in `internal/core/ports`.
2. **Green:** Write the minimal implementation in `internal/core/services` to pass the test.
3. **Refactor:** Clean up code.
4. **Wire:** Implement the actual adapter in `internal/adapters` and wire it in `main.go`.

### 4.2 Interfaces (Ports)

Do not import "heavy" libraries (AWS SDK, Cobra) into the core folder. Define interfaces instead.

**Bad:**
```go
func CreateBucket(svc *s3.Client)
```

**Good:**
```go
type CloudProvider interface { CreateBucket(name string) error }
```

## 5. GitHub Project Management

### 5.1 Milestones

- **Milestone 1:** Foundation & TDD Setup - Scaffolding logic with 100% unit test coverage using mocks.
- **Milestone 2:** FinOps Guardrails - Domain logic for cost estimation.
- **Milestone 3:** Kubernetes Operator - K8s Adapter implementation.
- **Milestone 4:** Observability - Telemetry Adapter implementation.

### 5.2 Labels

| Label           | Description                      |
|-----------------|----------------------------------|
| `area/core`     | Domain logic                     |
| `area/adapter`  | AWS/K8s/CLI implementations      |
| `kind/test`     | Test coverage improvements       |
