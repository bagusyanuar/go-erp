package request

type QueryPagination struct {
	Page     int `json:"page" query:"page" validate:"required"`
	PageSize int ` json:"page_size" query:"page_size" validate:"required"`
}

type QueryParams[T any] struct {
	Query T
}
