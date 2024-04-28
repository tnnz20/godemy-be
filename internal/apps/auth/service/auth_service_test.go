package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tnnz20/godemy-be/config"
	"github.com/tnnz20/godemy-be/internal/apps/auth"
	"github.com/tnnz20/godemy-be/internal/apps/auth/entities"
	"github.com/tnnz20/godemy-be/internal/apps/auth/repository"
	"github.com/tnnz20/godemy-be/internal/storage/postgres"
	"github.com/tnnz20/godemy-be/pkg/errs"
	"github.com/tnnz20/godemy-be/pkg/helpers"
)

var svc auth.Service

func init() {
	err := config.Load("../../../../config/config-local.yaml")
	if err != nil {
		panic(err)
	}

	db, err := postgres.NewConnection(config.Cfg.Database.Postgres)
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(db.GetDB())
	svc = newService(repo)
}

func TestServiceRegisterAuth(t *testing.T) {
	t.Run("Success Register", func(t *testing.T) {
		randString := helpers.GenerateRandomString(5)
		email := fmt.Sprintf("jhon%v@gmail.com", randString)
		req := entities.RegisterPayload{
			Name:     "Jhon",
			Email:    email,
			Password: "jhonpassword",
			Role:     "student",
		}

		err := svc.Register(context.Background(), req)
		require.Nil(t, err)
	})

	t.Run("Failed Register Email Already Exists", func(t *testing.T) {
		randString := helpers.GenerateRandomString(5)
		email := fmt.Sprintf("jhon%v@gmail.com", randString)
		req := entities.RegisterPayload{
			Name:     "Jhon",
			Email:    email,
			Password: "jhonpassword",
			Role:     "student",
		}

		err := svc.Register(context.Background(), req)
		require.Nil(t, err)

		// second register with same email
		err = svc.Register(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrEmailAlreadyExists, err)
	})

}
