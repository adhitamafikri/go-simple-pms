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
