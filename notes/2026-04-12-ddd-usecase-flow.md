# DDD Usecase Flow — Users Service Architecture

## Layer Overview

```
router/handlers  →  usecases  →  repository interface  ←  infrastructure/repositories
                                    (defined in entities)        (implements interface)
                                          ↑
                                    infrastructure/db
                                    (connection + migration)
```

**Key DDD principle**: inner layers define interfaces, outer layers implement them. Usecases depend on a repository *interface*, not a concrete implementation. The concrete Postgres implementation lives in `infrastructure`.

---

## What Goes Where

### `entities/` — domain layer
- Domain model structs (mirrors DB table shape, but no DB/HTTP knowledge)
- Repository interfaces (the contract for DB access)
- **No DTOs here** — those are a transport concern

### `usecases/` — business logic
- Depends on repository interfaces only (never on concrete implementations)
- Validates business rules, orchestrates operations across repos

### `infrastructure/` — all external concerns
- `infrastructure/db/` — DB connection setup
- `infrastructure/repositories/` — concrete SQL implementations of repo interfaces
- `infrastructure/sql/migrations/` — golang-migrate SQL files

### `router/` — HTTP concerns only
- Handler functions (call usecases, never repositories directly)
- Request/Response DTOs live here (or a `router/dto/` sub-package)

---

## Feature Implementation Flow

When adding a new feature, follow this order:

1. **Write the migration SQL** → `infrastructure/sql/migrations/000001_create_users_table.up.sql`
2. **Define the entity struct** → `entities/user.go`
3. **Define the repository interface** → `entities/user_repository.go`
4. **Set up DB connection** → `infrastructure/db/postgres.go`
5. **Implement the repository** → `infrastructure/repositories/user_repository.go`
6. **Write the usecase** → `usecases/user_usecase.go` (receives the interface, not the impl)
7. **Wire everything in bootstrap** → `configs/bootstrap.go` injects DB → repo → usecase → handler
8. **Add handler + DTOs** → `router/user_handler.go`

---

## Target Directory Structure

```
services/users/
├── entities/
│   ├── user.go              # domain struct (User, Role)
│   └── user_repository.go   # UserRepository interface
├── usecases/
│   └── user_usecase.go      # RegisterUser, LoginUser, etc.
├── infrastructure/
│   ├── db/
│   │   └── postgres.go      # *sqlx.DB connection + migrate runner
│   ├── repositories/
│   │   └── user_repository.go  # UserRepository impl (SQL queries)
│   └── sql/
│       └── migrations/
│           ├── 000001_create_users_table.up.sql
│           └── 000001_create_users_table.down.sql
└── router/
    ├── rest.go              # wires up route groups
    └── user_handler.go      # handlers + request/response DTOs
```

---

## DTOs Belong in the Router Layer

Keep DTOs out of `entities`. The entity is a pure domain model — no JSON tags, no validation tags, no HTTP-specific field names.

```go
// entities/user.go — domain model (no transport concerns)
type User struct {
    ID           uuid.UUID
    Email        string
    PasswordHash string
    RoleID       uuid.UUID
    CreatedAt    time.Time
}

// router/user_handler.go — transport DTOs
type RegisterRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}

type UserResponse struct {
    ID    string `json:"id"`
    Email string `json:"email"`
    Role  string `json:"role"`
}
```

The usecase returns a domain `User`. The handler maps it to `UserResponse` before writing JSON. Mapping logic lives where the format decision lives.

---

## Known Bug in `configs/bootstrap.go`

Lines 43–44 accidentally read `POSTGRES_DB` for both user and password:

```go
// WRONG
postgresUser := os.Getenv("POSTGRES_DB")
postgresPassword := os.Getenv("POSTGRES_DB")

// CORRECT
postgresUser := os.Getenv("POSTGRES_USER")
postgresPassword := os.Getenv("POSTGRES_PASSWORD")
```

Fix this before wiring up the DB connection or it will silently connect with wrong credentials.
