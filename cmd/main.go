package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tnnz20/godemy-be/config"
	"github.com/tnnz20/godemy-be/db"

	"github.com/tnnz20/godemy-be/internal/auth"
	"github.com/tnnz20/godemy-be/internal/teacher"
	"github.com/tnnz20/godemy-be/internal/user"
	"github.com/tnnz20/godemy-be/router"
)

func main() {

	// Load DB env
	env, err := config.LoadEnv(".")
	if err != nil {
		log.Fatalf("Failed to load env: %s", err)
	}

	// JWT Key
	jwtKey, err := config.LoadJWTKey(".")
	if err != nil {
		log.Fatalf("Failed to load env: %s", err)
	}

	// Db connection
	dbConn, err := db.NewDatabase(&env)
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}

	validate := validator.New()

	app := fiber.New()
	app.Use(cors.New())

	// User
	userRepo := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRepo)
	userHandler := user.NewHandler(userSvc, validate)

	router.UserRoutes(app, userHandler)

	// Auth
	authSvc := auth.NewService(userRepo, &jwtKey)
	authHandler := auth.NewHandler(authSvc, validate)

	router.AuthRoutes(app, authHandler)

	// Teacher
	teacherRepo := teacher.NewRepository(dbConn.GetDB())
	teacherSvc := teacher.NewService(teacherRepo)
	teacherHandler := teacher.NewHandler(teacherSvc, validate)

	router.TeacherRoutes(app, teacherHandler)

	log.Fatal(app.Listen(":5000"))

}
