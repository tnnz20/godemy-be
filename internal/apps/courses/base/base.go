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

	courses.Get("/",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Teacher}),
		handler.GetCoursesByUsersIdWithPagination)

	courses.Get("/list",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Teacher}),
		handler.GetCoursesByUsersId)

	courses.Get("/total",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Teacher}),
		handler.GetTotalCourses)

	courses.Post("/course/create",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Teacher}),
		handler.CreateCourse)

	courses.Post("/course/enroll",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Student}),
		handler.EnrollCourse)

	courses.Get("/course/enroll",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Student}),
		handler.GetCourseEnrollmentDetail)

	courses.Patch("/course/enroll/progress",
		middleware.Protected(),
		middleware.CheckRoles([]string{entities.ROLE_Student}),
		handler.UpdateProgressCourseEnrollment)

	courses.Get("/course/:courseId/enrolled", handler.GetEnrolledUsers)
	courses.Get("/course/:courseId/enrolled/total", handler.GetTotalEnrolledUsers)

}
