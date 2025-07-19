package usecase

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/pkg/response"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) response.ServiceResponse[any]
}

type UserService interface {
	Create(ctx context.Context, request *request.UserRequest) response.ServiceResponse[any]
}
