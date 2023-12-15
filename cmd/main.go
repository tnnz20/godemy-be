package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/tnnz20/godemy-be/internal/user"
	"github.com/tnnz20/godemy-be/router"

	"github.com/tnnz20/godemy-be/config"
	"github.com/tnnz20/godemy-be/db"
)

func main() {

	// Load env
	env, err := config.LoadEnv(".")
	if err != nil {
		fmt.Println(err)
	}

	// Db connection
	dbConn, err := db.NewDatabase(&env)
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}

	validate := validator.New()

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc, validate)

	app := fiber.New()
	app.Use(cors.New())

	router.SetupRoutes(app, userHandler)
	log.Fatal(app.Listen(":5000"))
}
