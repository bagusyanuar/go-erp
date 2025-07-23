package service

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/dto"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		FindAll(ctx context.Context, queryParams *request.UserQuery) lib.ServiceResponse[*[]dto.UserDTO]
		FinByID(ctx context.Context, id string) lib.ServiceResponse[*dto.UserDTO]
		Create(ctx context.Context, request *request.UserSchema) lib.ServiceResponse[any]
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

// FinByID implements UserService.
func (service *userServiceImpl) FinByID(ctx context.Context, id string) lib.ServiceResponse[*dto.UserDTO] {
	repositoryResponse := service.UserRepository.FindByID(ctx, id)
	if repositoryResponse.Error != nil {
		return lib.ServiceInternalServerError(lib.ServiceResponseOptions[*dto.UserDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return lib.ServiceOK(lib.ServiceResponseOptions[*dto.UserDTO]{
		Message: "successfully get user",
		Data:    dto.ToUser(repositoryResponse.Data),
	})
}

// FindAll implements UserService.
func (service *userServiceImpl) FindAll(ctx context.Context, queryParams *request.UserQuery) lib.ServiceResponse[*[]dto.UserDTO] {
	repositoryResponse := service.UserRepository.FindAll(ctx)
	if repositoryResponse.Error != nil {
		return lib.ServiceInternalServerError(lib.ServiceResponseOptions[*[]dto.UserDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}

	data := dto.ToUsers(repositoryResponse.Data)
	return lib.ServiceOK(lib.ServiceResponseOptions[*[]dto.UserDTO]{
		Message: "successfully get users",
		Data:    &data,
	})
}

// Create implements UserService.
func (service *userServiceImpl) Create(ctx context.Context, request *request.UserSchema) lib.ServiceResponse[any] {

	email := request.Email
	username := request.Username
	password := request.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return lib.ServiceInternalServerError(lib.ServiceResponseOptions[any]{
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
		return lib.ServiceInternalServerError(lib.ServiceResponseOptions[any]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return lib.ServiceCreated(lib.ServiceResponseOptions[any]{
		Message: "successfully create user",
	})
}
