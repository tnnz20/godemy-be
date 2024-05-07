package base

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	handlers "github.com/tnnz20/godemy-be/internal/apps/users/handler"
	"github.com/tnnz20/godemy-be/internal/apps/users/repository"
	"github.com/tnnz20/godemy-be/internal/apps/users/service"
)

func Init(router fiber.Router, db *sql.DB, secret string) {
	repo := repository.NewRepository(db)
	svc := service.NewService(repo, secret)
	handler := handlers.NewHandler(svc)

	_ = handler

	auth := router.Group("/auth", logger.New())
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)

}
