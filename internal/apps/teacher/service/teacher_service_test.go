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
var ValidUserId = "dbb6bcda-67b7-454d-b28b-7e14448dd0c9"

func TestTeacherService(t *testing.T) {
	t.Run("Success get teacher id", func(t *testing.T) {
		userId, err := uuid.Parse(ValidUserId)
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
		require.Empty(t, teacher)
		log.Println(teacher)
	})
}

func TestCreateCourseService(t *testing.T) {
	t.Run("Success create course", func(t *testing.T) {
		userId, err := uuid.Parse(ValidUserId)
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}
		require.Nil(t, err)

		req := entities.CreateCourseRequest{
			UserId: userId,
		}

		course := svc.CreateCourse(context.Background(), req)
		require.Nil(t, course)
	})
}

func TestGetCourseService(t *testing.T) {
	t.Run("Success get course by course code", func(t *testing.T) {
		courseCode := "go-muiYlpb"

		req := entities.GetCourseByCourseCodeRequest{
			CourseCode: courseCode,
		}

		course, err := svc.GetCourseByCourseCode(context.Background(), req)
		require.Nil(t, err)
		require.NotNil(t, course)
		log.Println(course)
	})

	t.Run("Failed get course by course code", func(t *testing.T) {
		courseCode := "go-muiYlp3"

		req := entities.GetCourseByCourseCodeRequest{
			CourseCode: courseCode,
		}

		course, err := svc.GetCourseByCourseCode(context.Background(), req)
		require.NotNil(t, err)
		require.Empty(t, course)
		log.Println(course)
	})

	t.Run("Success get course by teacher id", func(t *testing.T) {
		userId, err := uuid.Parse(ValidUserId)
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}

		req := entities.GetCourseByTeacherIdRequest{
			UserId: userId,
		}

		course, err := svc.GetCourseByTeacherId(context.Background(), req)
		require.Nil(t, err)
		require.NotNil(t, course)
		log.Println(course)
	})
}
