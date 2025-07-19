package response

type ServiceResponse[T any] struct {
	Status StatusCode `json:"status"`
	Data   T          `json:"data,omitempty"`
	Error  error      `json:"error"`
}
