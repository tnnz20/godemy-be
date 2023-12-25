package router

import (
	"github.com/tnnz20/godemy-be/internal/user"
	"github.com/tnnz20/godemy-be/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, userHandler *user.Handler) {
	user := app.Group("/users", logger.New())
	user.Post("/register", userHandler.CreateUser)
	user.Post("/sign-in", userHandler.SignIn)

	user.Get("/user/", middleware.Protected(), userHandler.GetUserProfileById)
}
