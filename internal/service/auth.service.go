package service

import (
	"context"
	"time"

	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/dto"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"github.com/bagusyanuar/go-erp/pkg/myexception"
	"github.com/golang-jwt/jwt/v5"
)

type (
	AuthService interface {
		Login(ctx context.Context, schema *request.LoginSchema) lib.ServiceResponse[*dto.LoginDTO]
	}

	authServiceImpl struct {
		AuthRepository repository.AuthRepository
		Config         *config.AppConfig
	}
)

func NewAuthService(authRepository repository.AuthRepository, cfg *config.AppConfig) AuthService {
	return &authServiceImpl{
		AuthRepository: authRepository,
		Config:         cfg,
	}
}

// Login implements AuthService.
func (service *authServiceImpl) Login(ctx context.Context, schema *request.LoginSchema) lib.ServiceResponse[*dto.LoginDTO] {
	email := schema.Email
	repositoryResponse := service.AuthRepository.Login(ctx, email)
	if repositoryResponse.Error != nil {
		return lib.ServiceInternalServerError(lib.ServiceResponseOptions[*dto.LoginDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}

	user := repositoryResponse.Data
	accessToken, err := service.createToken(user)
	if err != nil {
		return lib.ServiceInternalServerError(lib.ServiceResponseOptions[*dto.LoginDTO]{
			Error:   myexception.ErrGenerateToken,
			Message: myexception.ErrGenerateToken.Error(),
		})
	}

	refreshToken, err := service.createRefreshToken(user)
	if err != nil {
		return lib.ServiceInternalServerError(lib.ServiceResponseOptions[*dto.LoginDTO]{
			Error:   myexception.ErrGenerateToken,
			Message: myexception.ErrGenerateToken.Error(),
		})
	}
	return lib.ServiceOK(lib.ServiceResponseOptions[*dto.LoginDTO]{
		Message: "successfully login",
		Data: &dto.LoginDTO{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}

func (service *authServiceImpl) createToken(user *entity.User) (string, error) {
	JWTSignInMethod := jwt.SigningMethodHS256
	exp := time.Now().Add(time.Minute * time.Duration(service.Config.JWT.Expiration))
	claims := config.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    service.Config.JWT.Issuer,
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		Email:    user.Email,
		Username: user.Username,
	}
	token := jwt.NewWithClaims(JWTSignInMethod, claims)
	return token.SignedString([]byte(service.Config.JWT.Secret))
}

func (service *authServiceImpl) createRefreshToken(user *entity.User) (string, error) {
	JWTSignInMethod := jwt.SigningMethodHS256
	exp := time.Now().Add(time.Hour * 24 * time.Duration(service.Config.JWT.ExpirationRefreh))
	claims := jwt.RegisteredClaims{
		Issuer:    service.Config.JWT.Issuer,
		ExpiresAt: jwt.NewNumericDate(exp),
		Subject:   user.ID.String(),
	}
	token := jwt.NewWithClaims(JWTSignInMethod, claims)
	return token.SignedString([]byte(service.Config.JWT.SecretRefresh))
}
