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

	assessment.Get("/",
		middleware.Protected(),
		handler.GetAssessments,
	)

	assessment.Post("/assessment",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Student}),
		handler.CreateAssessment,
	)

	assessment.Get("/assessment",
		middleware.Protected(),
		handler.GetFilteredAssessmentResult,
	)

	assessment.Get("/assessment/total",
		middleware.Protected(),
		handler.GetTotalFilteredAssessmentResult,
	)

	assessment.Get("/assessment/users",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Student}),
		handler.GetUsersAssessment,
	)

	assessment.Post("/assessment/users",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Student}),
		handler.CreateUsersAssessment,
	)

	assessment.Patch("/assessment/users/status",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Student}),
		handler.UpdateUsersAssessmentStatus,
	)

}
