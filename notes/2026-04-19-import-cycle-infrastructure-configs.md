# Import Cycle: Infrastructure Must Not Import App-Layer Packages

## Problem

`configs` imported `infrastructures/database`, and `infrastructures/database` imported `configs` — causing a circular import error.

```
configs → infrastructures/database → configs  ✗
```

## Root Cause

`postgres.go` accepted `*configs.AppConfig` as a parameter, which forced it to import the `configs` package — the same package that wires up the DB connection.

## Fix

Infrastructure packages should receive plain primitives, not app-layer structs. Build the DSN in `bootstrap.go` (the composition root) and pass just the string.

**Before (`postgres.go`):**
```go
func NewPostgresConn(appConfig *configs.AppConfig) *sqlx.DB {
    dsn := "postgres://..."  // hardcoded or derived from appConfig
    db, err := sqlx.Connect("postgres", dsn)
```

**After (`postgres.go`):**
```go
func NewPostgresConn(dsn string) *sqlx.DB {
    db, err := sqlx.Connect("postgres", dsn)
```

**In `bootstrap.go`:**
```go
dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
    appConfig.POSTGRES_USER,
    appConfig.POSTGRES_PASSWORD,
    appConfig.APP_HOST,
    appConfig.POSTGRES_PORT,
    appConfig.POSTGRES_DB,
    appConfig.POSTGRES_SSL_MODE,
)
db := infrastructures.NewPostgresConn(dsn)
```

## Rule

**Infrastructure packages must never import application-layer packages** (`configs`, `usecases`, etc.).

They receive what they need as plain primitives. `bootstrap.go` is the composition root — it is the only place that knows about all layers and wires them together.

```
configs (composition root) → infrastructures/database  ✓
infrastructures/database   → configs                   ✗
```
