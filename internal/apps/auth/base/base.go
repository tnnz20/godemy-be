package base

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	handlers "github.com/tnnz20/godemy-be/internal/apps/auth/handler"
	"github.com/tnnz20/godemy-be/internal/apps/auth/repository"
	"github.com/tnnz20/godemy-be/internal/apps/auth/service"
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
