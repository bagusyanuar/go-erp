package response

type StatusCode int

const (
	Success             StatusCode = 200
	Created             StatusCode = 201
	BadRequest          StatusCode = 400
	Unauthorized        StatusCode = 401
	Forbidden           StatusCode = 403
	NotFound            StatusCode = 404
	InternalServerError StatusCode = 500
)
