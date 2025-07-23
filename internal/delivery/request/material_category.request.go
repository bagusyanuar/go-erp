package request

type MaterialCategorySchema struct {
	Name string `json:"name" validate:"required"`
}

type MaterialCategoryQuery struct {
	Param string `json:"param" query:"param"`
	QueryPagination
}
