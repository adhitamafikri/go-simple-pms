# Database Setup: sqlx, Migrations, and Repositories

## Architecture Pattern

This project follows **Clean Architecture** (layered). The dependency flow is:

```
router (delivery) → usecases (business logic) → repository interface (port)
                                                        ↑
                                           infrastructure/database (adapter/impl)
```

---

## 1. sqlx Connection

Create the connection in `infrastructure/database/postgres.go`, wire it in `configs/bootstrap.go`.

```
infrastructure/
  database/
    postgres.go   ← Connect(), returns *sqlx.DB
```

`postgres.go` builds the DSN from `AppConfig` and calls `sqlx.Connect`. The `*sqlx.DB` instance is then passed down via dependency injection through constructors.

---

## 2. Migrations

Store SQL migration files in `infrastructure/database/migrations/`. Use `golang-migrate/migrate` or `goose`.

```
infrastructure/
  database/
    migrations/
      000001_create_users_table.up.sql
      000001_create_users_table.down.sql
    postgres.go
```

Run migrations via a Makefile target (`make migrate-up`), not automatically at startup — keeps it explicit.

---

## 3. Repositories

- **Interface (port)** lives in `usecases/repository.go` — keeps business logic decoupled from sqlx.
- **Implementation (adapter)** lives in `infrastructure/database/repositories/`.

```
usecases/
  usecase.go          ← UseCase interface
  repository.go       ← UserRepository interface (port)

infrastructure/
  database/
    repositories/
      user_repository.go   ← implements UserRepository using *sqlx.DB
    migrations/
    postgres.go
```

The `clientUseCase` struct holds a `UserRepository` interface, not a concrete type. Replace `statement any` with `repo UserRepository`.

---

## 4. Wiring in bootstrap.go

```go
db   := database.Connect(appConfig)       // infrastructure
repo := repositories.NewUserRepo(db)      // infrastructure
uc   := usecase.New(repo)                 // usecases
r    := router.New(uc)                    // router
```

This matches the existing `// Connect DB` comment placeholder in `configs/bootstrap.go`.
