package middleware

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func VerifyJWT(cfg *config.AppConfig) fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
				Code:    fiber.StatusUnauthorized,
				Message: "unauthorized",
			})
		},
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    []byte(cfg.JWT.Secret),
		},
	})
}
