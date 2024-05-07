package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tnnz20/godemy-be/config"
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

	// Init router
	users.Init(router, db.GetDB(), cfg.App.Encryption.JWTSecret)

	port := fmt.Sprintf(":%s", cfg.App.Port)
	router.Listen(port)

}
