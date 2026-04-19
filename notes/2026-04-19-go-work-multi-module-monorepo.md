# Go Workspace (go.work) for Multi-Module Monorepos

## Problem

In a monorepo with multiple Go modules (e.g. `pkg/logger`, `pkg/auth`, `services/users`), importing one local module from another causes:

```
could not import github.com/adhitamafikri/go-simple-pms/pkg/logger
(missing metadata for import of "...")
```

This happens because each directory with a `go.mod` is an independent module. Without a workspace file, Go has no way to resolve cross-module imports locally — it would need the packages published to a registry (e.g. pkg.go.dev).

## Fix: Create a go.work at the Repo Root

```bash
# Run once from the repo root
go work init ./pkg/logger ./pkg/auth ./services/users
```

This generates a `go.work` file:

```
go 1.26.0

use (
    ./pkg/auth
    ./pkg/logger
    ./services/users
)
```

Now the Go toolchain (and gopls/LSP) can resolve all local module imports without needing them published.

## Adding New Modules Later

```bash
go work use ./services/new-service
```

## Key Rules

- `go.work` lives at the **repo root**, never inside a service or package.
- All modules that need to reference each other must be listed under `use`.
- `go.work` is for **local development**. In CI/production builds (e.g. Docker), each service is built independently from its own `go.mod` — `go.work` is typically gitignored or excluded from Docker context.
- After adding `go.work`, run `go mod tidy` inside each module to clean up dependencies.

## Postgres Driver Note

`sqlx` does not bundle a postgres driver. A blank import is required wherever the DB connection is opened:

```go
import _ "github.com/lib/pq"
```

Without this, `sqlx.Connect("postgres", dsn)` will panic at runtime with "unknown driver".
