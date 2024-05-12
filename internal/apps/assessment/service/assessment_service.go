package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tnnz20/godemy-be/internal/apps/assessment"
	"github.com/tnnz20/godemy-be/internal/apps/assessment/entities"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type service struct {
	assessment.Repository
}

func NewService(assessmentRepo assessment.Repository) assessment.Service {
	return &service{
		Repository: assessmentRepo,
	}
}

// CreateAssessment is a function to create a new assessment
func (s *service) CreateAssessment(ctx context.Context, req entities.CreateAssessmentRequest) (err error) {
	courseEnrollment, err := s.Repository.FindCoursesEnrollment(ctx, req.UsersId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrCourseEnrollmentNotFound
			return
		}
		return
	}

	NewAssessment := entities.NewAssessment(req.UsersId, courseEnrollment.CoursesId, req.AssessmentValue, req.AssessmentCode)

	if err = NewAssessment.Validate(); err != nil {
		return
	}
	err = s.Repository.CreateAssessment(ctx, NewAssessment)
	if err != nil {
		return err
	}
	return
}

// GetAssessment is a function to get assessment by user id
func (s *service) GetAssessment(ctx context.Context, req entities.GetAssessmentRequest) (res entities.AssessmentResponse, err error) {
	assessment, err := s.Repository.FindAssessment(ctx, req.UsersId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrAssessmentNotFound
			return
		}
		return res, err
	}

	res = entities.AssessmentResponse(assessment)

	return
}

// GetAssessments is a function to get all assessments by user id
func (s *service) GetAssessments(ctx context.Context, req entities.GetAssessmentRequest) (res []entities.AssessmentResponse, err error) {
	assessments, err := s.Repository.FindAssessments(ctx, req.UsersId)
	if err != nil {
		return []entities.AssessmentResponse{}, err
	}

	if len(assessments) == 0 {
		err = errs.ErrAssessmentNotFound
		return
	}

	for _, assessment := range assessments {
		res = append(res, entities.AssessmentResponse(assessment))
	}

	return
}

// GetAssessmentByAssessmentCode is a function to get assessment by assessment code
func (s *service) GetAssessmentByAssessmentCode(ctx context.Context, req entities.GetAssessmentByAssessmentCodeRequest) (res entities.AssessmentResponse, err error) {
	assessment, err := s.Repository.FindAssessmentByAssessmentCode(ctx, req.UsersId, req.AssessmentCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrAssessmentNotFound
			return
		}
		return res, err
	}

	res = entities.AssessmentResponse(assessment)

	return
}
