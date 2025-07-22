package request

type UnitSchema struct {
	Name string `json:"name" validate:"required"`
}

type UnitQuery struct {
	Param string `json:"param" query:"param"`
	QueryPagination
}
