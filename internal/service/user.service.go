package service

import (
	"context"
	"errors"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
	"github.com/bagusyanuar/go-erp/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		Create(ctx context.Context, request *request.UserRequest) response.ServiceResponse[any]
	}

	userServiceImpl struct {
		UserRepository repository.UserRepository
	}
)

// Create implements usecase.UserService.
func (service *userServiceImpl) Create(ctx context.Context, request *request.UserRequest) response.ServiceResponse[any] {
	res := response.ServiceResponse[any]{
		Status: response.InternalServerError,
		Error:  errors.New("unknown error"),
	}

	// email := request.Email
	// username := request.Username
	// password := request.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("administrator"), bcrypt.DefaultCost)
	if err != nil {
		res.Error = err
		return res
	}
	data := &entity.User{
		Email:    "admin@gmail.com",
		Username: "administrator",
		Password: string(hashedPassword),
	}

	repositoryResponse := service.UserRepository.Create(ctx, data)
	if repositoryResponse.Error != nil {
		return repositoryResponse
	}

	res.Status = response.Created
	res.Error = nil
	return res
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: userRepository,
	}
}
