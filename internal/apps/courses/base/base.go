package base

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	handlers "github.com/tnnz20/godemy-be/internal/apps/courses/handler"
	"github.com/tnnz20/godemy-be/internal/apps/courses/repository"
	"github.com/tnnz20/godemy-be/internal/apps/courses/service"
	"github.com/tnnz20/godemy-be/internal/apps/users/entities"
	"github.com/tnnz20/godemy-be/internal/middleware"
)

func Init(router fiber.Router, db *sql.DB) {
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	handler := handlers.NewHandler(svc)

	_ = handler

	courses := router.Group("/courses", logger.New())

	courses.Post("/create",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Teacher}),
		handler.CreateCourse)

	courses.Get("/course",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Teacher}),
		handler.GetCoursesByUsersId)

	courses.Post("/course/enroll",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Student}),
		handler.EnrollCourse)

	courses.Get("/course/enroll",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Student}),
		handler.GetCourseEnrollmentDetail)
}
