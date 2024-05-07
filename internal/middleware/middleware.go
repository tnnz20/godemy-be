package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/tnnz20/godemy-be/config"
	"github.com/tnnz20/godemy-be/pkg/jwt"
	"github.com/tnnz20/godemy-be/pkg/response"
)

// Protected is a middleware to check if the request has a valid token
func Protected() fiber.Handler {
	filename := "config/config-local.yaml"
	err := config.Load(filename)
	if err != nil {
		panic(err)
	}

	JWTSecret := config.Cfg.App.Encryption.JWTSecret

	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" {
			return response.ErrorUnauthorized(c)
		}

		bearer := strings.Split(authorization, "Bearer ")
		if len(bearer) != 2 {
			return response.ErrorInvalidToken(c)
		}

		token := bearer[1]
		id, role, err := jwt.ValidateToken(token, JWTSecret)
		if err != nil {
			return response.ErrorValidateToken(c, err)
		}

		c.Locals("role", role)
		c.Locals("id", id)

		return c.Next()
	}
}

// CheckRoles is a middleware to check if the request has a valid role
func CheckRoles(authorizedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := fmt.Sprintf("%v", c.Locals("role"))

		isExists := false
		for _, authorizedRole := range authorizedRoles {
			if role == authorizedRole {
				isExists = true
				break
			}
		}

		if !isExists {
			return response.ErrorForbiddenAccess(c)
		}

		return c.Next()
	}
}
