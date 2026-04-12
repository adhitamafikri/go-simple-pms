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
