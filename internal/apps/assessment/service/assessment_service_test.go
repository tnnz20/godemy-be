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
}

func TestGetFilteredAssessmentResult(t *testing.T) {
	t.Run("Success get filtered assessment", func(t *testing.T) {
		userId, err := uuid.Parse(validUserId)
		if err != nil {
			t.Error(err)
		}

		req := entities.GetAssessmentResultByAssessmentCodeRequest{
			UsersId:        userId,
			AssessmentCode: "chap-4",
			ModelPaginationPayload: entities.ModelPaginationPayload{
				Limit:  5,
				Offset: 0,
			},
		}

		assessments, err := svc.GetFilteredAssessmentResult(context.Background(), req)
		require.Nil(t, err)
		require.NotEmpty(t, assessments)
		log.Print(assessments)
	})

	t.Run("Failed get filtered assessment, assessment not found", func(t *testing.T) {
		userId, err := uuid.Parse("6286637a-3d6c-460a-b68a-956fd9553058")
		if err != nil {
			t.Error(err)
		}

		req := entities.GetAssessmentResultByAssessmentCodeRequest{
			UsersId:        userId,
			AssessmentCode: validAssessmentChapterCode,
			ModelPaginationPayload: entities.ModelPaginationPayload{
				Limit:  5,
				Offset: 0,
			},
		}

		_, err = svc.GetFilteredAssessmentResult(context.Background(), req)
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
			RandomArrayId:  []uint8{1, 2, 3, 4, 5},
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
			RandomArrayId:  []uint8{1, 2, 3, 4, 5},
		}

		err = svc.CreateUsersAssessment(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrAssessmentCodeRequired, err)
		log.Print(err)
	})

	t.Run("Failed create users assessment, assessment already created", func(t *testing.T) {
		userId, err := uuid.Parse(validUserId)
		if err != nil {
			t.Error(err)
		}

		req := entities.CreateUsersAssessmentRequest{
			UsersId:        userId,
			AssessmentCode: validAssessmentChapterCode,
			RandomArrayId:  []uint8{1, 2, 3, 4, 5},
		}

		err = svc.CreateUsersAssessment(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrAssessmentStatusAlreadyCreated, err)
		log.Print(err)
	})
}

func TestGetUsersAssessment(t *testing.T) {
	t.Run("Success get users assessment", func(t *testing.T) {
		userId, err := uuid.Parse(validUserId)
		if err != nil {
			t.Error(err)
		}

		req := entities.GetUsersAssessmentRequest{
			UsersId:        userId,
			AssessmentCode: validAssessmentChapterCode,
		}

		assessment, err := svc.GetUsersAssessment(context.Background(), req)
		require.Nil(t, err)
		require.NotEmpty(t, assessment)
		log.Print(assessment)
	})

	t.Run("Failed get users assessment, assessment not found", func(t *testing.T) {
		userId, err := uuid.Parse(validUserId)
		if err != nil {
			t.Error(err)
		}

		req := entities.GetUsersAssessmentRequest{
			UsersId:        userId,
			AssessmentCode: "chap-10",
		}

		_, err = svc.GetUsersAssessment(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrAssessmentNotFound, err)
		log.Print(err)
	})
}

func TestUpdateUsersAssessmentStatus(t *testing.T) {
	t.Run("Success update users assessment status", func(t *testing.T) {
		userId, err := uuid.Parse(validUserId)
		if err != nil {
			t.Error(err)
		}

		req := entities.UpdateUsersAssessmentStatusRequest{
			UsersId:        userId,
			AssessmentCode: validAssessmentChapterCode,
			Status:         5,
		}

		err = svc.UpdateUsersAssessmentStatus(context.Background(), req)
		require.Nil(t, err)

		getAssessmentReq := entities.GetUsersAssessmentRequest{
			UsersId:        userId,
			AssessmentCode: validAssessmentChapterCode,
		}
		assessment, err := svc.GetUsersAssessment(context.Background(), getAssessmentReq)
		require.Nil(t, err)
		require.Equal(t, entities.AssessmentStatusOnGoing, assessment.Status)
		log.Print(assessment)
	})

	t.Run("Failed update users assessment status, assessment not found", func(t *testing.T) {
		userId, err := uuid.Parse(validUserId)
		if err != nil {
			t.Error(err)
		}

		req := entities.UpdateUsersAssessmentStatusRequest{
			UsersId:        userId,
			AssessmentCode: "chap-10",
			Status:         5,
		}

		err = svc.UpdateUsersAssessmentStatus(context.Background(), req)

		require.NotNil(t, err)
		require.Equal(t, errs.ErrAssessmentNotFound, err)
	})

	t.Run("Failed update users assessment status, status invalid", func(t *testing.T) {
		userId, err := uuid.Parse(validUserId)
		if err != nil {
			t.Error(err)
		}

		req := entities.UpdateUsersAssessmentStatusRequest{
			UsersId:        userId,
			AssessmentCode: validAssessmentChapterCode,
			Status:         8,
		}

		err = svc.UpdateUsersAssessmentStatus(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidAssessmentStatus, err)
	})
}
