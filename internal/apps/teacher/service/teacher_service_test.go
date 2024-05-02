package service

import (
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/tnnz20/godemy-be/config"
	"github.com/tnnz20/godemy-be/internal/apps/teacher"
	"github.com/tnnz20/godemy-be/internal/apps/teacher/entities"
	"github.com/tnnz20/godemy-be/internal/apps/teacher/repository"
	"github.com/tnnz20/godemy-be/internal/storage/postgres"
)

var svc teacher.Service

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
	svc = NewService(repo)
}

var ErrParsingUUID = "Error Parsing UUID: "

func TestTeacherService(t *testing.T) {
	t.Run("Success get teacher id", func(t *testing.T) {
		userId, err := uuid.Parse("dbb6bcda-67b7-454d-b28b-7e14448dd0c9")
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}

		req := entities.GetTeacherIdByUserIdRequest{
			UserId: userId,
		}

		teacher, err := svc.GetTeacherIdByUserId(context.Background(), req)

		require.Nil(t, err)
		require.NotNil(t, teacher)
		log.Println(teacher)
	})

	t.Run("Failed get teacher id", func(t *testing.T) {
		userId, err := uuid.Parse("dbb6bcda-67b7-454d-b28b-7e14428dd0c9")
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}

		req := entities.GetTeacherIdByUserIdRequest{
			UserId: userId,
		}

		teacher, err := svc.GetTeacherIdByUserId(context.Background(), req)

		require.NotNil(t, err)
		log.Println(teacher)
	})
}

func TestTeacherCoursesService(t *testing.T) {
	t.Run("Success get teacher id", func(t *testing.T) {
		userId, err := uuid.Parse("dbb6bcda-67b7-454d-b28b-7e14448dd0c9")
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}

		req := entities.CreateCourseRequest{
			UserId: userId,
		}

		course := svc.CreateCourse(context.Background(), req)
		require.Nil(t, err)
		log.Println(course)
	})
}
