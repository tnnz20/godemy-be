package service

import (
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/tnnz20/godemy-be/config"
	"github.com/tnnz20/godemy-be/internal/apps/profile"
	"github.com/tnnz20/godemy-be/internal/apps/profile/entities"
	"github.com/tnnz20/godemy-be/internal/apps/profile/repository"
	"github.com/tnnz20/godemy-be/internal/storage/postgres"
)

var svc profile.Service

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
	svc = NewService(repo, config.Cfg.App.Encryption.JWTSecret)
}

func TestProfileService(t *testing.T) {
	t.Run("Success retrieve profile user", func(t *testing.T) {
		userId, err := uuid.Parse("85472c11-eded-450a-88b7-82d5898ddc40")
		if err != nil {
			log.Fatalf("error parsing %+v", err)
		}

		req := entities.GetProfileByUserIdRequest{
			UserId: userId,
		}

		profile, err := svc.GetProfileByUserId(context.Background(), req)
		require.NotNil(t, profile)
		require.Nil(t, err)
		log.Println(profile)
	})

	t.Run("Failed retrieve profile user", func(t *testing.T) {
		userId, err := uuid.Parse("85372c11-eded-450a-88b7-82d5898ddc40")
		if err != nil {
			log.Fatalf("error parsing %+v", err)
		}

		req := entities.GetProfileByUserIdRequest{
			UserId: userId,
		}

		profile, err := svc.GetProfileByUserId(context.Background(), req)
		require.Nil(t, profile)
		require.NotNil(t, err)

		log.Println(err)

	})

}
