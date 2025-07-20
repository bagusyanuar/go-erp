package service

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/dto"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
	"github.com/bagusyanuar/go-erp/internal/pkg/myexception"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		Create(ctx context.Context, request *request.UserRequest) dto.ServiceResponse[any]
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
func (service *userServiceImpl) Create(ctx context.Context, request *request.UserRequest) dto.ServiceResponse[any] {
	res := dto.ServiceResponse[any]{
		Status: dto.InternalServerError,
		Error:  myexception.ErrUnknown,
	}

	email := request.Email
	username := request.Username
	password := request.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		res.Error = err
		return res
	}
	data := &entity.User{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
	}

	repositoryResponse := service.UserRepository.Create(ctx, data)
	if repositoryResponse.Error != nil {
		return repositoryResponse
	}

	res.Status = dto.Created
	res.Error = nil
	return res
}
