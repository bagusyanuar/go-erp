package dto

type StatusCode int

const (
	OK                  StatusCode = 200
	Created             StatusCode = 201
	BadRequest          StatusCode = 400
	Unauthorized        StatusCode = 401
	Forbidden           StatusCode = 403
	NotFound            StatusCode = 404
	UnprocessableEntity StatusCode = 422
	InternalServerError StatusCode = 500
)

type ServiceResponse[T any] struct {
	Status  StatusCode `json:"status"`
	Message string     `json:"message"`
	Data    T          `json:"data,omitempty"`
	Error   error      `json:"error"`
}

type APIResponse[T any] struct {
	Code    StatusCode `json:"code"`
	Message string     `json:"message"`
	Data    T          `json:"data,omitempty"`
	Meta    any        `json:"meta,omitempty"`
}

func FromService[T any](service ServiceResponse[T], meta any) APIResponse[T] {

	return APIResponse[T]{
		Code:    service.Status,
		Message: service.Message,
		Data:    service.Data,
		Meta:    meta,
	}
}
