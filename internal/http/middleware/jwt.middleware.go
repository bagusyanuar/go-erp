package middleware

import (
	"errors"
	"strings"

	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/pkg/exception"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func VerifyJWT(cfg *config.AppConfig) fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			if errors.Is(err, jwtware.ErrJWTMissingOrMalformed) {
				return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
					Code:    fiber.StatusUnauthorized,
					Message: exception.ErrMissingOrMalformed.Error(),
				})
			}

			if strings.Contains(err.Error(), "token is expired") {
				return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
					Code:    fiber.StatusUnauthorized,
					Message: "Token has expired",
				})
			}

			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
				Code:    fiber.StatusUnauthorized,
				Message: exception.ErrUnauthorized.Error(),
			})
		},
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    []byte(cfg.JWT.Secret),
		},
	})
}
