Below is a **fully written, ready-to-create GitHub issue** for **Sub-issue 1.2**.

It is grounded in the project documents’ requirement for a **fully functional codebase following clean architecture principles**, with **microservices**, an **API Gateway**, **containerized services**, **secure APIs**, **testing**, and later **observability/deployment** support. It also follows your approved Phase 1 architecture, which fixes the runtime services as **Auth, Users, Posts, Feed, and Notifications**, with Redis used for sessions/feed caching, Kafka for async communication, and PostgreSQL per service where applicable.  

---

# GitHub Issue — Sub-issue 1.2

**Title:** Scaffold Go service templates with clean architecture layers for gateway and all microservices

**Parent Issue:** Establish Phase 2 implementation foundation

**Type:** Sub-issue

**Priority:** P0

**Complexity:** Large

**Labels:** `backend`, `architecture`, `golang`, `clean-architecture`, `setup`

---

## Description

Create the **initial runnable Go application skeleton** for all runtime units in the system:

* API Gateway
* Auth Service
* Users Service
* Posts Service
* Feed Service
* Notification Service

Each runtime unit must follow a **consistent clean-architecture-oriented internal structure** so later implementation work can be added without restructuring the codebase. The scaffold should provide:

* a standard Go application bootstrap,
* layered package boundaries,
* config loading hook,
* router/server initialization,
* health endpoint(s),
* graceful shutdown,
* placeholder route groups,
* dependency wiring points,
* logging middleware hooks,
* readiness for future DB/Redis/Kafka integration.

This issue is about creating the **engineering skeleton**, not implementing business functionality yet. It should leave the project in a state where each service can **compile, start, and respond to `/health`** while preserving architectural separation for future development. This directly supports the Phase 2 requirement for a **clean architecture codebase**, and it matches the Phase 1 design where services are independently deployable and own distinct responsibilities.  

---

## Why this issue matters

The project brief requires the final implementation to be:

* microservices-based,
* containerized,
* secure,
* testable,
* observable,
* documented,
* aligned with the approved Phase 1 design.

If the services start with inconsistent structures, business logic will end up mixed into handlers, testing will become harder, and later work such as logging, tracing, retries, validation, Kafka consumers, Redis integration, and database repositories will be messy to add. A clean scaffold ensures that:

* business logic can be unit tested,
* adapters like PostgreSQL/Redis/Kafka can be added cleanly,
* every service follows the same engineering pattern,
* the final implementation remains consistent with the architecture grading criteria. 

---

## Scope

This issue includes:

* standardizing the internal Go service structure,
* creating the initial `main.go` bootstrap for each service,
* creating clean-architecture folders and placeholder files,
* adding `/health` and optional `/ready`,
* initializing router/server setup,
* adding placeholder route registration,
* defining where future domain/service/repository/handler code will live,
* adding graceful shutdown,
* creating minimal service README notes.

This issue does **not** include:

* implementing real OAuth2,
* implementing JWT logic,
* implementing profile/post/feed/notification business logic,
* connecting to PostgreSQL/Redis/Kafka for real operations,
* writing migrations,
* implementing OpenAPI docs,
* implementing security middleware,
* implementing monitoring,
* implementing Docker or Compose.

Those are later issues.

---

## Architectural constraints from the documents

The scaffolding must reflect the approved architecture and not drift into a generic monolith.

### Fixed service responsibilities from Phase 1

* **Auth Service** → Google sign-in, token/session lifecycle
* **Users Service** → profiles and follow relationships
* **Posts Service** → post CRUD
* **Feed Service** → personalized feed
* **Notification Service** → follow/interaction notifications
* **API Gateway** → public entry point and JWT/session validation path later on. 

### Data/runtime expectations

* PostgreSQL is used per service where applicable.
* Feed Service uses Redis and **no persistent database**.
* Kafka is used for async communication.
* Redis is used for sessions and feed caching. 

Therefore, the scaffolding must leave clean extension points for:

* HTTP handlers,
* PostgreSQL repositories,
* Redis repositories,
* Kafka producers/consumers,
* middleware,
* config,
* domain models.

---

## Clean architecture structure to use

Use the following **standard service structure** for every runtime unit:

```text
service-name/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── app/
│   ├── config/
│   ├── domain/
│   ├── handler/
│   │   └── http/
│   ├── middleware/
│   ├── repository/
│   ├── service/
│   ├── transport/
│   │   └── http/
│   └── bootstrap/
├── pkg/
├── configs/
├── .env.example
├── go.mod
└── README.md
```

### Purpose of each folder

**`cmd/server/`**
Application entrypoint. Contains `main.go` only. Responsible for starting the service.

**`internal/app/`**
Optional place for shared app-level types or application startup coordination.

**`internal/config/`**
Typed configuration loading from environment variables.

**`internal/domain/`**
Core business entities and domain models.

**`internal/handler/http/`**
HTTP handlers that translate requests/responses only.

**`internal/middleware/`**
Cross-cutting HTTP middleware such as request ID, logging, recovery, auth hooks, rate limiting hooks.

**`internal/repository/`**
Persistence or infrastructure adapters:

* PostgreSQL repos
* Redis repos
* Kafka producers/consumers

**`internal/service/`**
Business logic layer.

**`internal/transport/http/`**
Router setup, route registration, server-related HTTP wiring.

**`internal/bootstrap/`**
Application wiring and dependency assembly.

**`pkg/`**
Reusable code safe to share if needed, though overuse should be avoided.

This structure supports the project’s clean-architecture requirement and keeps later testing focused on core business logic rather than handler-heavy code.

---

## Required implementation by service

## 1. API Gateway scaffold

### Purpose

The gateway is required by the course and Phase 1 architecture. It will later route requests, enforce authentication, validate sessions through Redis, and apply rate limiting.  

### Required initial files

```text
api-gateway/
├── cmd/server/main.go
├── internal/config/config.go
├── internal/handler/http/health_handler.go
├── internal/middleware/request_id.go
├── internal/middleware/logging.go
├── internal/middleware/recovery.go
├── internal/transport/http/router.go
├── internal/bootstrap/app.go
├── pkg/
├── configs/
├── .env.example
├── go.mod
└── README.md
```

### Initial behavior

* starts an HTTP server,
* loads config,
* registers `/health`,
* registers placeholder route groups:

  * `/api/v1/auth`
  * `/api/v1/users`
  * `/api/v1/posts`
  * `/api/v1/feed`
  * `/api/v1/notifications`
* applies request ID, logging, and recovery middleware,
* shuts down gracefully.

---

## 2. Auth Service scaffold

### Purpose

Auth handles Google OAuth2 login, JWT issuance, and Redis-backed sessions later. Phase 1 explicitly chooses Google OAuth2 + JWT + Redis session storage. 

### Required initial files

```text
auth-service/
├── cmd/server/main.go
├── internal/config/config.go
├── internal/domain/session.go
├── internal/domain/auth_user.go
├── internal/handler/http/health_handler.go
├── internal/handler/http/auth_handler.go
├── internal/service/auth_service.go
├── internal/repository/redis/session_repository.go
├── internal/middleware/request_id.go
├── internal/middleware/logging.go
├── internal/middleware/recovery.go
├── internal/transport/http/router.go
├── internal/bootstrap/app.go
├── pkg/
├── configs/
├── .env.example
├── go.mod
└── README.md
```

### Initial behavior

* starts service,
* exposes `/health`,
* creates placeholder route group:

  * `/api/v1/auth/login`
  * `/api/v1/auth/callback`
  * `/api/v1/auth/logout`
* contains stub auth service and stub Redis session repository interfaces,
* no real OAuth yet.

---

## 3. Users Service scaffold

### Purpose

Users service handles profiles and follow relationships. It later owns a PostgreSQL database and publishes follow events to Kafka. 

### Required initial files

```text
users-service/
├── cmd/server/main.go
├── internal/config/config.go
├── internal/domain/user.go
├── internal/domain/follow.go
├── internal/handler/http/health_handler.go
├── internal/handler/http/user_handler.go
├── internal/service/user_service.go
├── internal/repository/postgres/user_repository.go
├── internal/repository/postgres/follow_repository.go
├── internal/repository/kafka/follow_producer.go
├── internal/middleware/request_id.go
├── internal/middleware/logging.go
├── internal/middleware/recovery.go
├── internal/transport/http/router.go
├── internal/bootstrap/app.go
├── pkg/
├── configs/
├── migrations/
├── .env.example
├── go.mod
└── README.md
```

### Initial behavior

* starts service,
* exposes `/health`,
* creates placeholder route group:

  * `GET /api/v1/users/me`
  * `PATCH /api/v1/users/me`
  * `GET /api/v1/users/{id}`
  * `POST /api/v1/users/{id}/follow`
  * `DELETE /api/v1/users/{id}/follow`

---

## 4. Posts Service scaffold

### Purpose

Posts service implements the project’s required CRUD core feature and later publishes `post.created` events to Kafka.  

### Required initial files

```text
posts-service/
├── cmd/server/main.go
├── internal/config/config.go
├── internal/domain/post.go
├── internal/handler/http/health_handler.go
├── internal/handler/http/post_handler.go
├── internal/service/post_service.go
├── internal/repository/postgres/post_repository.go
├── internal/repository/kafka/post_producer.go
├── internal/middleware/request_id.go
├── internal/middleware/logging.go
├── internal/middleware/recovery.go
├── internal/transport/http/router.go
├── internal/bootstrap/app.go
├── pkg/
├── configs/
├── migrations/
├── .env.example
├── go.mod
└── README.md
```

### Initial behavior

* starts service,
* exposes `/health`,
* creates placeholder route group:

  * `POST /api/v1/posts`
  * `GET /api/v1/posts/{id}`
  * `PATCH /api/v1/posts/{id}`
  * `DELETE /api/v1/posts/{id}`

---

## 5. Feed Service scaffold

### Purpose

Feed service is performance-sensitive and later consumes Kafka events and writes to Redis using the fan-out-on-write approach. Phase 1 explicitly states Feed has **no persistent database**. 

### Required initial files

```text
feed-service/
├── cmd/server/main.go
├── internal/config/config.go
├── internal/domain/feed_item.go
├── internal/handler/http/health_handler.go
├── internal/handler/http/feed_handler.go
├── internal/service/feed_service.go
├── internal/repository/redis/feed_repository.go
├── internal/repository/kafka/post_consumer.go
├── internal/repository/kafka/follow_consumer.go
├── internal/middleware/request_id.go
├── internal/middleware/logging.go
├── internal/middleware/recovery.go
├── internal/transport/http/router.go
├── internal/bootstrap/app.go
├── pkg/
├── configs/
├── .env.example
├── go.mod
└── README.md
```

### Initial behavior

* starts service,
* exposes `/health`,
* creates placeholder route:

  * `GET /api/v1/feed`

---

## 6. Notification Service scaffold

### Purpose

Notification service later consumes follow and interaction events and exposes notification retrieval endpoints. 

### Required initial files

```text
notification-service/
├── cmd/server/main.go
├── internal/config/config.go
├── internal/domain/notification.go
├── internal/handler/http/health_handler.go
├── internal/handler/http/notification_handler.go
├── internal/service/notification_service.go
├── internal/repository/postgres/notification_repository.go
├── internal/repository/kafka/follow_consumer.go
├── internal/repository/kafka/interaction_consumer.go
├── internal/middleware/request_id.go
├── internal/middleware/logging.go
├── internal/middleware/recovery.go
├── internal/transport/http/router.go
├── internal/bootstrap/app.go
├── pkg/
├── configs/
├── migrations/
├── .env.example
├── go.mod
└── README.md
```

### Initial behavior

* starts service,
* exposes `/health`,
* creates placeholder route:

  * `GET /api/v1/notifications`

---

# Detailed implementation requirements

## A. `main.go` responsibilities for every service

Each service’s `cmd/server/main.go` must:

1. load configuration from env using `internal/config`,
2. initialize the app via `internal/bootstrap`,
3. create an HTTP server,
4. start listening on configured port,
5. log startup,
6. listen for OS termination signals,
7. gracefully shut down with timeout.

### Pseudocode shape

```go
func main() {
    cfg := config.Load()

    app, err := bootstrap.NewApp(cfg)
    if err != nil {
        log.Fatal(err)
    }

    srv := &http.Server{
        Addr:         ":" + cfg.Port,
        Handler:      app.Router,
        ReadTimeout:  cfg.HTTP.ReadTimeout,
        WriteTimeout: cfg.HTTP.WriteTimeout,
        IdleTimeout:  cfg.HTTP.IdleTimeout,
    }

    go func() {
        log.Printf("starting %s on port %s", cfg.ServiceName, cfg.Port)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal(err)
        }
    }()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
    <-stop

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal(err)
    }
}
```

This should be consistent across all services.

---

## B. `config.go` responsibilities

Every service must have `internal/config/config.go` with:

* a typed `Config` struct,
* `Load()` function,
* basic validation for required env vars,
* defaults for non-sensitive values.

At scaffold time, config can be minimal, but must include:

* service name,
* port,
* environment,
* log level,
* timeouts.

Service-specific placeholders:

* gateway/auth: Redis + JWT + service URLs
* auth: Google OAuth env placeholders
* users/posts/notifications: DB placeholders
* feed: Redis + Kafka placeholders
* users/posts/feed/notifications: Kafka placeholders

This aligns with the Phase 1 architecture decisions around Redis, JWT, Kafka, and DB per service. 

---

## C. Router requirements

Each service must expose at least:

* `GET /health`

Optional but recommended:

* `GET /ready`

The router layer must:

* register middleware,
* register health route,
* register placeholder API groups,
* not contain business logic.

Use a simple HTTP router consistently across services.

### Example route grouping

For Users Service:

```go
r.Route("/api/v1/users", func(r chi.Router) {
    r.Get("/me", userHandler.GetMe)
    r.Patch("/me", userHandler.UpdateMe)
    r.Get("/{id}", userHandler.GetByID)
    r.Post("/{id}/follow", userHandler.FollowUser)
    r.Delete("/{id}/follow", userHandler.UnfollowUser)
})
```

Handlers can return `501 Not Implemented` initially except `/health`.

---

## D. Middleware requirements

Each service must include these three middleware files in scaffold form:

### `request_id.go`

* reads `X-Request-ID` from incoming headers,
* generates one if missing,
* stores it in context,
* sets it in response header.

### `logging.go`

* logs method, path, status, duration, request ID,
* structured logging preferred.

### `recovery.go`

* catches panics,
* logs them,
* returns `500`.

This prepares the codebase for the later observability/logging requirements in Phase 2 and supports the Phase 1 target that debugging should be fast and centralized.  

---

## E. Health endpoint requirements

Every service must expose:

* `GET /health`

Response should be simple and consistent, for example:

```json
{
  "status": "ok",
  "service": "users-service"
}
```

Optional `GET /ready` may later check DB/Redis/Kafka readiness, but for now it can return a placeholder healthy response.

This is important for:

* container readiness later,
* deployment validation,
* monitoring,
* CI smoke tests.

---

## F. Dependency wiring via `bootstrap/app.go`

Each service should have a `bootstrap.NewApp(cfg)` function that:

* initializes the router,
* constructs placeholder repositories/services/handlers,
* returns an app object with the router.

This creates one clean place for dependency injection and avoids manually wiring everything in `main.go`.

### Example shape

```go
type App struct {
    Router http.Handler
}

func NewApp(cfg config.Config) (*App, error) {
    router := httptransport.NewRouter(cfg)
    return &App{Router: router}, nil
}
```

Later, this function will be expanded to create:

* DB pools,
* Redis clients,
* Kafka producers/consumers,
* services,
* handlers.

---

## G. Placeholder repository/service/handler contracts

Do not leave `internal/service` and `internal/repository` empty. Add starter interfaces or stubs so the architecture is visible.

### Example for Posts Service

`internal/service/post_service.go`

```go
package service

type PostService interface {
    CreatePost() error
    GetPost() error
    UpdatePost() error
    DeletePost() error
}
```

`internal/repository/postgres/post_repository.go`

```go
package postgres

type PostRepository interface {
    // placeholder for CRUD methods
}
```

`internal/handler/http/post_handler.go`

```go
package http

import "net/http"

type PostHandler struct{}

func NewPostHandler() *PostHandler { return &PostHandler{} }

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}
```

These are intentionally minimal, but they establish the service layering clearly.

---

## H. Required placeholder route matrix

Create the following route placeholders now so the API surface already matches the project plan.

### API Gateway

* `/health`
* `/api/v1/auth/*`
* `/api/v1/users/*`
* `/api/v1/posts/*`
* `/api/v1/feed`
* `/api/v1/notifications`

### Auth Service

* `GET /health`
* `GET /api/v1/auth/login`
* `GET /api/v1/auth/callback`
* `POST /api/v1/auth/logout`

### Users Service

* `GET /health`
* `GET /api/v1/users/me`
* `PATCH /api/v1/users/me`
* `GET /api/v1/users/{id}`
* `POST /api/v1/users/{id}/follow`
* `DELETE /api/v1/users/{id}/follow`

### Posts Service

* `GET /health`
* `POST /api/v1/posts`
* `GET /api/v1/posts/{id}`
* `PATCH /api/v1/posts/{id}`
* `DELETE /api/v1/posts/{id}`

### Feed Service

* `GET /health`
* `GET /api/v1/feed`

### Notification Service

* `GET /health`
* `GET /api/v1/notifications`

These placeholder routes map directly to the documented functional requirements and user stories. 

---

## I. Service-specific scaffold notes

## API Gateway

Do not add DB code here.
Add placeholder config for:

* service upstream URLs,
* Redis,
* JWT secret/session validation mode.

## Auth Service

Do not add PostgreSQL migrations unless your actual implementation later requires them.
At scaffold level, focus on:

* auth domain model,
* session repository placeholder,
* OAuth config placeholders.

This matches the Phase 1 ADR focused on Google OAuth2 + Redis-backed sessions. 

## Users / Posts / Notification Services

These should include `migrations/` and repository placeholders because they are DB-backed service types under the database-per-service decision. 

## Feed Service

Do not create persistent SQL repository placeholders as primary storage.
Use Redis and Kafka placeholder repos only, because the Phase 1 ADR explicitly says feed has no persistent database. 

---

## Exact deliverables

This issue must produce:

1. Runnable Go skeleton for all six runtime units
2. Consistent clean-architecture folder structure in each unit
3. `main.go` entrypoint in every service
4. `config.go` in every service
5. `router.go` in every service
6. health endpoint handler in every service
7. middleware scaffolding for request ID, logging, recovery
8. bootstrap wiring point in every service
9. placeholder domain/service/repository/handler files
10. service READMEs updated to reflect each service’s role

---

## Acceptance criteria

This issue is complete when:

* each of the six runtime units compiles as a Go application,
* each service has a consistent internal structure,
* each service starts successfully on its configured port,
* each service exposes `GET /health`,
* route placeholders matching the project architecture are registered,
* middleware skeletons for request ID, logging, and recovery exist,
* `main.go` handles graceful shutdown,
* business logic is not yet implemented beyond placeholders,
* Feed Service scaffold does **not** assume a persistent DB,
* Users/Posts/Notification services include DB-ready repository structure,
* the scaffold is ready for later config, Docker, database, Redis, Kafka, and API implementation.

---

## Definition of done

A reviewer should be able to:

1. open any service folder,
2. immediately understand where:

   * config lives,
   * domain models live,
   * HTTP handlers live,
   * business logic lives,
   * repositories live,
   * middleware lives,
   * server/router wiring lives,
3. run the service,
4. hit `/health`,
5. see that the structure supports the documented architecture and Phase 2 implementation requirements.  

---

## Dependencies

Depends on **Sub-issue 1.1 — Create monorepo structure for API Gateway and all microservices**.

---

## Risks / notes

* Do not place business logic directly in handlers.
* Do not make a single shared `internal/` folder at repo root that collapses all service boundaries.
* Do not create a shared database abstraction that suggests services share a DB, because Phase 1 explicitly chose **database per service**. 
* Do not scaffold Feed as a SQL-first service, because Phase 1 explicitly states Feed uses Redis and no persistent DB. 
* Keep placeholders small and clear; this issue is a scaffold issue, not a feature implementation issue.

---

## Suggested implementation order inside this issue

1. create `main.go` for all services
2. add `config.go` for all services
3. add middleware skeletons
4. add router registration
5. add health handlers
6. add placeholder domain/service/repository/handler files
7. test that all services compile and start

---

If you want, I can do the next step and turn this into a **practical implementation pack** with:

* exact file contents for every scaffolded file,
* recommended router library and module names,
* and PowerShell commands to generate the entire service skeleton automatically.
