package middleware

import (
	"context"
	"errors"
	"strings"

	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/pkg/constant"
	"github.com/bagusyanuar/go-erp/pkg/exception"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals("user").(*jwt.Token) // casting ke *jwt.Token
			claims := token.Claims.(jwt.MapClaims) // baru ambil claim-nya
			userID, ok := claims["sub"].(string)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
					Code:    fiber.StatusUnauthorized,
					Message: "Invalid token subject",
				})
			}

			// Masukkan userID ke context.Context
			ctx := context.WithValue(c.UserContext(), constant.UserIDKey, userID)
			c.SetUserContext(ctx)
			return c.Next()
		},
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    []byte(cfg.JWT.Secret),
		},
	})
}
