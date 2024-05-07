package service

import (
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/tnnz20/godemy-be/config"
	"github.com/tnnz20/godemy-be/internal/apps/courses"
	"github.com/tnnz20/godemy-be/internal/apps/courses/entities"
	"github.com/tnnz20/godemy-be/internal/apps/courses/repository"
	"github.com/tnnz20/godemy-be/internal/storage/postgres"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

var svc courses.Service

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

var (
	ValidUserId     = "b17903ef-0de1-4155-9b9e-98e01e5ff894"
	ValidCourseCode = "go-C7CRGmU"
)

var ErrParsingUUID = "Error Parsing UUID: "

func TestCreateCoursesService(t *testing.T) {
	t.Run("Success create course", func(t *testing.T) {
		userId, err := uuid.Parse(ValidUserId)
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}
		req := entities.CreateCoursePayload{
			UsersId: userId,
		}

		err = svc.CreateCourse(context.Background(), req)
		require.Nil(t, err)
	})
}

func TestGetCoursesService(t *testing.T) {
	t.Run("Success get course by course code", func(t *testing.T) {

		req := entities.GetCourseByCourseCodePayload{
			CourseCode: ValidCourseCode,
		}
		course, err := svc.GetCourseByCourseCode(context.Background(), req)
		require.Nil(t, err)
		require.NotNil(t, course)
		log.Println(course)
	})

	t.Run("Failed get course by course code, must be required", func(t *testing.T) {
		req := entities.GetCourseByCourseCodePayload{
			CourseCode: "",
		}

		_, err := svc.GetCourseByCourseCode(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrCourseCodeRequired, err)
	})

	t.Run("Failed get course by course code, must be 10 characters", func(t *testing.T) {
		req := entities.GetCourseByCourseCodePayload{
			CourseCode: "go-123",
		}

		_, err := svc.GetCourseByCourseCode(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidCourseCodeLength, err)
	})

	t.Run("Failed get course by course code, course not found", func(t *testing.T) {
		req := entities.GetCourseByCourseCodePayload{
			CourseCode: "go-1234567",
		}

		_, err := svc.GetCourseByCourseCode(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrCourseNotFound, err)
	})

	t.Run("Success get courses by user id", func(t *testing.T) {
		userId, err := uuid.Parse(ValidUserId)
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}

		var reqPagination entities.GetCoursesByUsersIdWithPaginationPayload
		reqPagination.UsersId = userId
		req := reqPagination.GenerateDefaultValue()

		log.Println(req)

		courses, err := svc.GetCoursesByUsersIdWithPagination(context.Background(), req)
		require.Nil(t, err)
		require.NotNil(t, courses)
		log.Println(courses)
	})
}
