package service

import (
	"context"
	"errors"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/dto"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	UserService interface {
		FindAll(ctx context.Context, queryParams *request.UserQuery) response.ServiceResponse[*[]dto.UserDTO]
		FindByID(ctx context.Context, id string) response.ServiceResponse[*dto.UserDTO]
		Create(ctx context.Context, request *request.UserSchema) response.ServiceResponse[any]
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
func (service *userServiceImpl) Create(ctx context.Context, request *request.UserSchema) response.ServiceResponse[any] {

	email := request.Email
	username := request.Username
	password := request.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return response.ServiceInternalServerError(response.ServiceResponseOptions[any]{
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
		return response.ServiceInternalServerError(response.ServiceResponseOptions[any]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return response.ServiceCreated(response.ServiceResponseOptions[any]{
		Message: "successfully create user",
	})
}

// FindAll implements UserService.
func (service *userServiceImpl) FindAll(ctx context.Context, queryParams *request.UserQuery) response.ServiceResponse[*[]dto.UserDTO] {
	repositoryResponse := service.UserRepository.FindAll(ctx, queryParams)
	if repositoryResponse.Error != nil {
		return response.ServiceInternalServerError(response.ServiceResponseOptions[*[]dto.UserDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}

	data := dto.ToUsers(repositoryResponse.Data)
	return response.ServiceOK(response.ServiceResponseOptions[*[]dto.UserDTO]{
		Message: "successfully get users",
		Data:    &data,
		Meta:    repositoryResponse.Meta,
	})
}

// FinByID implements UserService.
func (service *userServiceImpl) FindByID(ctx context.Context, id string) response.ServiceResponse[*dto.UserDTO] {
	repositoryResponse := service.UserRepository.FindByID(ctx, id)
	if repositoryResponse.Error != nil {
		if errors.Is(repositoryResponse.Error, gorm.ErrRecordNotFound) {
			return response.ServiceNotFound(response.ServiceResponseOptions[*dto.UserDTO]{
				Error:   repositoryResponse.Error,
				Message: repositoryResponse.Message,
			})
		}
		return response.ServiceInternalServerError(response.ServiceResponseOptions[*dto.UserDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return response.ServiceOK(response.ServiceResponseOptions[*dto.UserDTO]{
		Message: "successfully get user",
		Data:    dto.ToUser(repositoryResponse.Data),
	})
}
