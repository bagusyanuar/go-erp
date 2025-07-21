package request

type QueryPagination struct {
	Page    int `json:"page" query:"page" validate:"required"`
	PerPage int ` json:"per_page" query:"per_page" validate:"required"`
}

type QueryParams[T any] struct {
	Query T
}
