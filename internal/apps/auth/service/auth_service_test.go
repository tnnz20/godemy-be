package service

import (
	"context"
	"fmt"
	"log"
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

var randString string = helpers.GenerateRandomString(5)

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

		email := fmt.Sprintf("jhon%v@godemy.com", randString)
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

func TestServiceLoginAuth(t *testing.T) {
	t.Run("Success Login", func(t *testing.T) {
		email := fmt.Sprintf("jhon%v@gmail.com", randString)
		pass := "mysecretpassword"
		req := entities.RegisterPayload{
			Email:    email,
			Password: pass,
			Role:     "student",
			Name:     "Jhon",
		}
		err := svc.Register(context.Background(), req)
		require.Nil(t, err)

		reqLogin := entities.LoginPayload{
			Email:    email,
			Password: pass,
		}

		token, err := svc.Login(context.Background(), reqLogin)
		require.Nil(t, err)
		require.NotEmpty(t, token)
		log.Println(token)
	})

	t.Run("Failed Login Email Not Found", func(t *testing.T) {
		email := fmt.Sprintf("jhon123%v@gmail.com", randString)
		pass := "mysecretpassword"
		req := entities.RegisterPayload{
			Email:    email,
			Password: pass,
			Role:     "student",
			Name:     "Jhon",
		}
		err := svc.Register(context.Background(), req)
		require.Nil(t, err)

		reqLogin := entities.LoginPayload{
			Email:    "xasd@gmail.com",
			Password: pass,
		}

		token, err := svc.Login(context.Background(), reqLogin)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrEmailNotFound, err)
		require.Empty(t, token)
	})

	t.Run("Failed Login Wrong Password", func(t *testing.T) {
		email := fmt.Sprintf("jhon51%v@gmail.com", randString)
		pass := "mysecretpassword"
		req := entities.RegisterPayload{
			Email:    email,
			Password: pass,
			Role:     "student",
			Name:     "Jhon",
		}
		err := svc.Register(context.Background(), req)
		require.Nil(t, err)

		reqLogin := entities.LoginPayload{
			Email:    email,
			Password: "wrongpassword",
		}

		token, err := svc.Login(context.Background(), reqLogin)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrWrongPassword, err)
		require.Empty(t, token)
		log.Println(token)
	})

}
