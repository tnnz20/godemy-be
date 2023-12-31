package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func JwtClaim(c *fiber.Ctx) (id uuid.UUID, role string) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	claimId := claims["id"].(string)
	claimRole := claims["role"].(string)

	id, _ = uuid.Parse(claimId)
	role = claimRole
	return
}
