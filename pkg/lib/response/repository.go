package response

type RepositoryResponse[T any] struct {
	Message string
	Error   error
	Data    T
	Meta    any
}

func MakeRepositoryError[T any](err error) RepositoryResponse[T] {
	return RepositoryResponse[T]{
		Message: err.Error(),
		Error:   err,
	}
}

func MakeRepositorySuccess[T any](data T, meta any) RepositoryResponse[T] {
	return RepositoryResponse[T]{
		Message: "successfull",
		Error:   nil,
		Data:    data,
		Meta:    meta,
	}
}
