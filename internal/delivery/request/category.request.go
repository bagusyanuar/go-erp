package request

type CategorySchema struct {
	Name string `json:"name" validate:"required"`
}

type CategoryQuery struct {
	Param string `json:"param" query:"param"`
	QueryPagination
}
