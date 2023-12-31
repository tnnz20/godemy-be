package router

import (
	"github.com/tnnz20/godemy-be/internal/auth"
	"github.com/tnnz20/godemy-be/internal/class"
	"github.com/tnnz20/godemy-be/internal/student"
	"github.com/tnnz20/godemy-be/internal/teacher"
	"github.com/tnnz20/godemy-be/internal/user"
	"github.com/tnnz20/godemy-be/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func apiRoutes(app *fiber.App, route string) (api fiber.Router) {
	api = app.Group("/api/"+route, logger.New())
	return
}

func UserRoutes(app *fiber.App, userHandler *user.Handler) {
	user := apiRoutes(app, "user")
	user.Post("/sign-up", userHandler.CreateUser)
	user.Get("/profile", middleware.Protected(), userHandler.GetUserProfileById)
}

func AuthRoutes(app *fiber.App, authHandler *auth.Handler) {
	auth := apiRoutes(app, "auth")
	auth.Post("/sign-in", authHandler.SignIn)
}

func TeacherRoutes(app *fiber.App, teacherHandler *teacher.Handler) {
	teacher := apiRoutes(app, "teachers")
	teacher.Get("/teacher", middleware.Protected(), teacherHandler.GetTeacherIdByUserId)
	teacher.Get("/teacher/classes", middleware.Protected(), teacherHandler.GetAllClassByTeacherId)
	// teacher.Get("/teacher/classes/class", middleware.Protected(), teacherHandler.GetListStudentByClassName)
}

func StudentRoutes(app *fiber.App, studentHandler *student.Handler) {
	student := apiRoutes(app, "students")
	student.Get("/student", middleware.Protected(), studentHandler.GetStudentByUserId)
	student.Patch("/student/threshold", middleware.Protected(), studentHandler.IncrementThreshold)
	student.Post("/student/assessment", middleware.Protected(), studentHandler.InsertAssessment)
}

func ClassRoutes(app *fiber.App, classHandler *class.Handler) {
	class := apiRoutes(app, "classes")
	class.Get("/", classHandler.GetAllClass)
	class.Post("/", middleware.Protected(), classHandler.CreateClass)
	class.Patch("/class/student", middleware.Protected(), classHandler.UpdateStudentClass)
}
