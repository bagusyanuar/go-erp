package service

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		Create(ctx context.Context, request *request.UserRequest) lib.ServiceResponse[any]
	}

	userServiceImpl struct {
		UserRepository repository.UserRepository
	}
)

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: userRepository,
	}
}

// Create implements UserService.
func (service *userServiceImpl) Create(ctx context.Context, request *request.UserRequest) lib.ServiceResponse[any] {

	email := request.Email
	username := request.Username
	password := request.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return lib.ServiceInternalServerError[any](lib.ServiceResponseOptions[any]{
			Error:   err,
			Message: err.Error(),
		})
	}
	data := &entity.User{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
	}

	repositoryResponse := service.UserRepository.Create(ctx, data)
	if repositoryResponse.Error != nil {
		return lib.ServiceInternalServerError[any](lib.ServiceResponseOptions[any]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return lib.ServiceCreated(lib.ServiceResponseOptions[any]{
		Message: "successfully create user",
	})
}
