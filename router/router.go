package router

import (
	"github.com/tnnz20/godemy-be/internal/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, userHandler *user.Handler) {
	user := app.Group("/user", logger.New())
	user.Post("/", userHandler.CreateUser)
}
