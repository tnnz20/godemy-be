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
	ValidCourseCode    = "go-wlDqAUx"
	validCourseName    = "go-fundamental-1"
	ValidUserIdTeacher = "28126e28-f696-4af5-9bfc-d5c9b7d952d2"
	ValidUserIdStudent = "0c550f03-79d1-4f55-aeac-770f1aefa818"
)

var ErrParsingUUID = "Error Parsing UUID: "

func TestCreateCoursesService(t *testing.T) {
	t.Run("Success create course", func(t *testing.T) {
		userId, err := uuid.Parse(ValidUserIdTeacher)
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}
		req := entities.CreateCoursePayload{
			UsersId:    userId,
			CourseName: validCourseName,
		}

		err = svc.CreateCourse(context.Background(), req)
		require.Nil(t, err)
	})

	t.Run("Failed create course, course name must be required", func(t *testing.T) {
		userId, err := uuid.Parse(ValidUserIdTeacher)
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}
		req := entities.CreateCoursePayload{
			UsersId:    userId,
			CourseName: "",
		}

		err = svc.CreateCourse(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrCourseNameRequired, err)
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
		userId, err := uuid.Parse(ValidUserIdTeacher)
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}

		var reqPagination entities.GetCoursesByUsersIdWithPaginationPayload
		reqPagination.UsersId = userId
		reqPagination.ModelPaginationPayload.GenerateDefaultValue()

		log.Println(reqPagination)

		courses, err := svc.GetCoursesByUsersIdWithPagination(context.Background(), reqPagination)
		require.Nil(t, err)
		require.NotNil(t, courses)
		log.Println(courses)
	})

	t.Run("Success get total courses by user id", func(t *testing.T) {
		userId, err := uuid.Parse(ValidUserIdTeacher)
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}

		req := entities.GetTotalCoursesByUsersIdPayload{
			UsersId: userId,
		}

		total, err := svc.GetTotalCourses(context.Background(), req)
		require.Nil(t, err)
		require.NotNil(t, total)
		log.Println(total)
	})
}

func TestEnrollCoursesService(t *testing.T) {
	t.Run("Success enroll course", func(t *testing.T) {
		userId, err := uuid.Parse(ValidUserIdStudent)
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}
		req := entities.EnrollCoursePayload{
			UsersId:    userId,
			CourseCode: ValidCourseCode,
		}
		err = svc.EnrollCourse(context.Background(), req)
		require.Nil(t, err)
	})
}

func TestGetCourseEnrollmentService(t *testing.T) {
	t.Run("Success get course enrollment by user id", func(t *testing.T) {
		userId, err := uuid.Parse(ValidUserIdStudent)
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}
		req := entities.GetCourseEnrollmentByUsersIdPayload{
			UsersId: userId,
		}
		courseEnrollment, err := svc.GetCourseEnrollmentByUsersId(context.Background(), req)
		require.Nil(t, err)
		require.NotNil(t, courseEnrollment)
		log.Println(courseEnrollment)
	})
}

func TestUpdateProgressCourseEnrollmentService(t *testing.T) {
	t.Run("Success update progress course enrollment", func(t *testing.T) {
		userId, err := uuid.Parse(ValidUserIdStudent)
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}
		req := entities.UpdateEnrollmentProgressPayload{
			UsersId:  userId,
			Progress: 15,
		}
		err = svc.UpdateProgressCourseEnrollment(context.Background(), req)
		require.Nil(t, err)
	})

	t.Run("Failed update progress course enrollment", func(t *testing.T) {
		userId, err := uuid.Parse(ValidUserIdStudent)
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}
		req := entities.UpdateEnrollmentProgressPayload{
			UsersId:  userId,
			Progress: 8,
		}
		err = svc.UpdateProgressCourseEnrollment(context.Background(), req)
		require.NotNil(t, err)
	})
}

func TestGetListCourseByCourseIdService(t *testing.T) {
	t.Run("Success get list course by course id", func(t *testing.T) {
		courseId, err := uuid.Parse("6183b21e-de39-44b7-bb39-4b9cadc26847")
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}
		req := entities.GetEnrolledUsersByCourseIdPayload{
			CourseId: courseId,
		}

		req.GenerateDefaultValue()
		log.Println(req)
		courses, err := svc.GetEnrolledUsersByCourseId(context.Background(), req)
		require.Nil(t, err)
		require.NotNil(t, courses)
		log.Println(courses)
	})
}

func TestGetTotalEnrolledUsersByCourseIdService(t *testing.T) {
	t.Run("Success get total enrolled users by course id", func(t *testing.T) {
		courseId, err := uuid.Parse("6183b21e-de39-44b7-bb39-4b9cadc26847")
		if err != nil {
			log.Fatal(ErrParsingUUID, err)
		}
		req := entities.GetTotalEnrolledUsersByCourseIdPayload{
			CourseId: courseId,
		}

		total, err := svc.GetTotalEnrolledUsersByCourseId(context.Background(), req)
		require.Nil(t, err)
		require.NotNil(t, total)
		log.Println(total)
	})
}
