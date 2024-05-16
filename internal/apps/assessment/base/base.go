package base

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	handlers "github.com/tnnz20/godemy-be/internal/apps/assessment/handler"
	"github.com/tnnz20/godemy-be/internal/apps/assessment/repository"
	"github.com/tnnz20/godemy-be/internal/apps/assessment/service"
	"github.com/tnnz20/godemy-be/internal/apps/users/entities"
	"github.com/tnnz20/godemy-be/internal/middleware"
)

func Init(router fiber.Router, db *sql.DB) {
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	handler := handlers.NewHandler(svc)

	_ = handler

	assessment := router.Group("/assessments", logger.New())
	assessment.Post("/assessment/create",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Student}),
		handler.CreateAssessment,
	)
	assessment.Get("/assessment/users",
		middleware.Protected(),
		handler.GetAssessmentsFiltered,
	)
	assessment.Get("/assessment/code",
		middleware.Protected(),
		handler.GetAssessmentByAssessmentCode,
	)
}
