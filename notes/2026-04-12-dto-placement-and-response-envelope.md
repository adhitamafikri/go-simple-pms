# DTO Placement and Standard Response Envelope

## The Split

| Type | Location | Reason |
|---|---|---|
| Standard response envelope | `pkg/response/` | Shared across all handlers, not a domain concept |
| Endpoint-specific request/response | `router/` (inline in handler file) | Tightly coupled to the handler, changes with business needs |

`pkg/` is the idiomatic Go convention for shared internal utilities — not domain, not infrastructure, just reusable primitives within the service.

---

## Directory Structure

```
services/users/
├── pkg/
│   └── response/
│       └── response.go    # BaseResponse, PaginatedResponse
├── router/
│   └── user_handler.go    # endpoint-specific DTOs + uses pkg/response for envelopes
```

---

## Implementation Using Generics

```go
// pkg/response/response.go
package response

const Version = "1.0.0"

type Base[T any] struct {
    Version string `json:"version"`
    Message string `json:"message"`
    Data    T      `json:"data"`
}

type PaginationMeta struct {
    CurrentPage int `json:"current_page"`
    LastPage    int `json:"last_page"`
    PerPage     int `json:"per_page"`
}

type Paginated[T any] struct {
    Version string         `json:"version"`
    Message string         `json:"message"`
    Data    []T            `json:"data"`
    Meta    PaginationMeta `json:"meta"`
}

func OK[T any](data T, message string) Base[T] {
    return Base[T]{Version: Version, Message: message, Data: data}
}

func Created[T any](data T, message string) Base[T] {
    return Base[T]{Version: Version, Message: message, Data: data}
}

func WithPagination[T any](data []T, message string, meta PaginationMeta) Paginated[T] {
    return Paginated[T]{Version: Version, Message: message, Data: data, Meta: meta}
}
```

---

## Usage in a Handler

```go
// router/user_handler.go

// Endpoint-specific DTO — lives here, not in entities/
type UserResponse struct {
    ID    string `json:"id"`
    Email string `json:"email"`
}

func (h *UserHandler) Register(ctx *gin.Context) {
    // usecase returns a domain User (entities.User)
    // handler maps it to the transport DTO before writing JSON
    ctx.JSON(http.StatusCreated, response.Created(UserResponse{
        ID:    user.ID.String(),
        Email: user.Email,
    }, "User registered successfully"))
}
```

---

## Response Shapes

**Single resource (200 OK / 201 Created)**
```json
{
  "version": "1.0.0",
  "message": "User registered successfully",
  "data": { ... }
}
```

**Paginated list**
```json
{
  "version": "1.0.0",
  "message": "Users fetched successfully",
  "data": [ ... ],
  "meta": {
    "current_page": 1,
    "last_page": 20,
    "per_page": 25
  }
}
```
