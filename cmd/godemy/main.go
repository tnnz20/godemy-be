package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tnnz20/godemy-be/config"
	assessment "github.com/tnnz20/godemy-be/internal/apps/assessment/base"
	courses "github.com/tnnz20/godemy-be/internal/apps/courses/base"
	users "github.com/tnnz20/godemy-be/internal/apps/users/base"
	"github.com/tnnz20/godemy-be/internal/storage/postgres"
)

func main() {

	filename := "config/config-local.yaml"
	if err := config.Load(filename); err != nil {
		panic(err)
	}

	cfg := config.Cfg

	db, err := postgres.NewConnection(cfg.Database.Postgres)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err := db.Migrate(); err != nil {
		panic(err)
	}

	router := fiber.New(fiber.Config{
		AppName: cfg.App.Name,
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Init router
	users.Init(router, db.GetDB(), cfg.App.Encryption.JWTSecret)
	courses.Init(router, db.GetDB())
	assessment.Init(router, db.GetDB())

	port := fmt.Sprintf(":%s", cfg.App.Port)
	router.Listen(port)

}
