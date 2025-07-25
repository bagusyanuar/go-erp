package exception

import "errors"

var (
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInvalidBearer       = errors.New("invalid bearer type")
	ErrMissingOrMalformed  = errors.New("token is missing or malformed")
	ErrInvalidUserFormat   = errors.New("invalid user format")
	ErrInvalidJWTParse     = errors.New("invalid jwt parse")
	ErrUserNotFound        = errors.New("user not found")
	ErrRecordNotFound      = errors.New("record not found")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrUnknown             = errors.New("unknown error")
	ErrValidateRequest     = errors.New("unprocessable request validation")
	ErrGenerateToken       = errors.New("failed to generate access token")
	ErrDBInsertFailed      = errors.New("failed to insert data")
	ErrDBUpdateFailed      = errors.New("failed to update data")
	ErrDBDeleteFailed      = errors.New("failed to delete data")
	ErrDBQueryFailed       = errors.New("failed to query data")
)
