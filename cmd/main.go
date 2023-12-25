package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tnnz20/godemy-be/config"
	"github.com/tnnz20/godemy-be/db"
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

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep, &jwtKey)
	userHandler := user.NewHandler(userSvc, validate)

	app := fiber.New()
	app.Use(cors.New())

	router.SetupRoutes(app, userHandler)
	log.Fatal(app.Listen(":5000"))

}
