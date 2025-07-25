package request

type MaterialSchema struct {
	Name       string   `json:"name" validate:"required"`
	Categories []string `json:"categories" validate:"required,dive,uuid4"`
}

type MaterialQuery struct {
	Param string `json:"param" query:"param"`
	QueryPagination
}
