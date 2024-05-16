package service

import (
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/tnnz20/godemy-be/config"
	"github.com/tnnz20/godemy-be/internal/apps/assessment"
	"github.com/tnnz20/godemy-be/internal/apps/assessment/entities"
	"github.com/tnnz20/godemy-be/internal/apps/assessment/repository"
	"github.com/tnnz20/godemy-be/internal/storage/postgres"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

var svc assessment.Service

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

const (
	validUserId                string = "19f75097-e613-4769-a6dd-fe97ccb8e54b"
	validAssessmentChapterCode string = "chap-3"
)

func TestCreateAssessmentResult(t *testing.T) {
	t.Run("Success create assessment", func(t *testing.T) {
		userId, err := uuid.Parse(validUserId)
		if err != nil {
			t.Error(err)
		}

		req := entities.CreateAssessmentRequest{
			UsersId:         userId,
			AssessmentValue: 7,
			AssessmentCode:  validAssessmentChapterCode,
		}

		err = svc.CreateAssessmentResult(context.Background(), req)
		require.Nil(t, err)
	})

	t.Run("Failed create assessment, course enrollment not found", func(t *testing.T) {
		userId, err := uuid.Parse("f8739934-a08a-494c-a0da-3f66553819f2")
		if err != nil {
			t.Error(err)
		}

		req := entities.CreateAssessmentRequest{
			UsersId:         userId,
			AssessmentValue: 80,
			AssessmentCode:  validAssessmentChapterCode,
		}

		err = svc.CreateAssessmentResult(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrCourseEnrollmentNotFound, err)
		log.Print(err)
	})
}

func TestGetAssessmentsResult(t *testing.T) {
	t.Run("Success get assessments", func(t *testing.T) {
		userId, err := uuid.Parse(validUserId)
		if err != nil {
			t.Error(err)
		}

		req := entities.GetAssessmentRequest{
			UsersId: userId,
		}

		assessments, err := svc.GetAssessmentsResult(context.Background(), req)
		require.Nil(t, err)
		require.NotEmpty(t, assessments)
		log.Print(assessments)
	})

	t.Run("Failed get assessments, assessment not found", func(t *testing.T) {
		userId, err := uuid.Parse("6286637a-3d6c-460a-b68a-956fd9553058")
		if err != nil {
			t.Error(err)
		}

		req := entities.GetAssessmentRequest{
			UsersId: userId,
		}

		_, err = svc.GetAssessmentsResult(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrAssessmentNotFound, err)
		log.Print(err)
	})

	t.Run("Success get assessment by assessment code", func(t *testing.T) {
		userId, err := uuid.Parse(validUserId)
		if err != nil {
			t.Error(err)
		}

		req := entities.GetAssessmentByAssessmentCodeRequest{
			UsersId:        userId,
			AssessmentCode: validAssessmentChapterCode,
		}

		assessment, err := svc.GetAssessmentResultByAssessmentCode(context.Background(), req)
		require.Nil(t, err)
		require.NotEmpty(t, assessment)
		log.Print(assessment)
	})

	t.Run("Failed get assessment by assessment code, assessment not found", func(t *testing.T) {
		userId, err := uuid.Parse(validUserId)
		if err != nil {
			t.Error(err)
		}

		req := entities.GetAssessmentByAssessmentCodeRequest{
			UsersId:        userId,
			AssessmentCode: "chap-10",
		}

		_, err = svc.GetAssessmentResultByAssessmentCode(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrAssessmentNotFound, err)
		log.Print(err)
	})
}

func TestCreateUsersAssessment(t *testing.T) {
	t.Run("Success create users assessment", func(t *testing.T) {
		userId, err := uuid.Parse(validUserId)
		if err != nil {
			t.Error(err)
		}

		req := entities.CreateUsersAssessmentRequest{
			UsersId:        userId,
			AssessmentCode: validAssessmentChapterCode,
			RandomArrayId:  []int{1, 2, 3, 4, 5},
		}

		err = svc.CreateUsersAssessment(context.Background(), req)
		require.Nil(t, err)
	})

	t.Run("Failed create users assessment, assessment code required", func(t *testing.T) {
		userId, err := uuid.Parse(validUserId)
		if err != nil {
			t.Error(err)
		}

		req := entities.CreateUsersAssessmentRequest{
			UsersId:        userId,
			AssessmentCode: "",
			RandomArrayId:  []int{1, 2, 3, 4, 5},
		}

		err = svc.CreateUsersAssessment(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrAssessmentCodeRequired, err)
		log.Print(err)
	})
}
