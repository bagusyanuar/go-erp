package request

type UserSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8,symbol"`
}

type UserQuery struct {
	Param string `json:"param" query:"param" validate:"required"`
	QueryPagination
}
