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
func (s *service) CreateAssessmentResult(ctx context.Context, req entities.CreateAssessmentPayload) (err error) {
	courseEnrollment, err := s.Repository.FindCoursesEnrollment(ctx, req.UsersId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrCourseEnrollmentNotFound
			return
		}
		return
	}

	var status uint8
	if req.AssessmentValue >= 80 {
		status = 1
	} else {
		status = 0
	}

	NewAssessmentResult := entities.NewAssessmentResult(req.UsersId, courseEnrollment.CoursesId, req.AssessmentValue, req.AssessmentCode, status)

	if err = NewAssessmentResult.Validate(); err != nil {
		return
	}
	err = s.Repository.CreateAssessmentResult(ctx, NewAssessmentResult)
	if err != nil {
		return err
	}
	return
}

// GetAssessments is a function to get all assessments by user id
func (s *service) GetAssessmentsResult(ctx context.Context, req entities.GetAssessmentPayload) (res []entities.AssessmentResponse, err error) {
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

// GetFilteredAssessmentResult is a function to get assessment by assessment code
func (s *service) GetFilteredAssessmentResult(ctx context.Context, req entities.GetAssessmentResultWithPaginationPayload) (res []entities.AssessmentResponse, err error) {

	NewAssessmentPagination := entities.NewAssessmentPagination(req.Limit, req.Offset)
	assessments, err := s.Repository.FindAssessmentsFilteredByCode(ctx, req.UsersId, req.AssessmentCode, NewAssessmentPagination)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrAssessmentNotFound
			return
		}
		return res, err
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

// GetTotalFilteredAssessmentResult is a function to get total assessment by assessment code
func (s *service) GetTotalFilteredAssessmentResult(ctx context.Context, req entities.GetAssessmentResultWithPaginationPayload) (res entities.AssessmentTotalResponse, err error) {
	total, err := s.Repository.FindTotalAssessmentsFilteredByCode(ctx, req.UsersId, req.AssessmentCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrAssessmentNotFound
			return
		}
		return entities.AssessmentTotalResponse{}, err
	}

	if total == 0 {
		err = errs.ErrAssessmentNotFound
		return
	}

	res = entities.AssessmentTotalResponse{
		Total: total,
	}

	return res, err

}

func (s *service) GetAssessmentsResultUsers(ctx context.Context, req entities.GetAssessmentResultsByCourseIdPayload) (res []entities.AssessmentResultUsersResponse, err error) {

	NewAssessmentPagination := entities.NewAssessmentPagination(req.Limit, req.Offset)
	assessments, err := s.Repository.FindAssessmentsByCourseId(ctx, req.CoursesId, req.Name, req.AssessmentCode, req.Sort, req.Status, NewAssessmentPagination)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrAssessmentNotFound
			return
		}
		return res, err
	}

	if len(assessments) == 0 {
		err = errs.ErrAssessmentNotFound
		return
	}

	for _, assessment := range assessments {
		res = append(res, entities.AssessmentResultUsersResponse{
			ID:              assessment.Id,
			UsersId:         assessment.UsersId,
			Name:            assessment.Name,
			CoursesId:       assessment.CoursesId,
			AssessmentValue: assessment.AssessmentValue,
			AssessmentCode:  assessment.AssessmentCode,
			Status:          assessment.Status,
			CreatedAt:       assessment.CreatedAt,
		})
	}

	return
}

func (s *service) GetTotalAssessmentsResultUsers(ctx context.Context, req entities.GetAssessmentResultsByCourseIdPayload) (res entities.AssessmentTotalResponse, err error) {
	total, err := s.Repository.FindTotalAssessmentsByCourseId(ctx, req.CoursesId, req.Name, req.AssessmentCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrAssessmentNotFound
			return
		}
		return entities.AssessmentTotalResponse{}, err
	}

	if total == 0 {
		err = errs.ErrAssessmentNotFound
		return
	}

	res = entities.AssessmentTotalResponse{
		Total: total,
	}

	return
}

func (s *service) CreateUsersAssessment(ctx context.Context, req entities.CreateUsersAssessmentPayload) (err error) {

	NewAssessmentUser := entities.NewAssessmentUser(req.UsersId, req.AssessmentCode, req.RandomArrayId)

	if req.AssessmentCode == "" {
		return errs.ErrAssessmentCodeRequired
	}

	if err = NewAssessmentUser.ValidateAssessmentUserCode(); err != nil {
		return
	}

	assessment, err := s.Repository.FindUsersAssessment(ctx, NewAssessmentUser.UsersId, NewAssessmentUser.AssessmentCode)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}

	if assessment.IsStatusCreated() {
		return errs.ErrAssessmentStatusAlreadyCreated
	}

	err = s.Repository.CreateUsersAssessment(ctx, NewAssessmentUser)
	if err != nil {
		return err
	}

	return
}

func (s *service) GetUsersAssessment(ctx context.Context, req entities.GetUsersAssessmentPayload) (res entities.AssessmentUserResponse, err error) {
	newUserAssessment := entities.AssessmentUser{
		UsersId:        req.UsersId,
		AssessmentCode: req.AssessmentCode,
	}

	assessment, err := s.Repository.FindUsersAssessment(ctx, newUserAssessment.UsersId, newUserAssessment.AssessmentCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrAssessmentNotFound
			return
		}
		return entities.AssessmentUserResponse{}, err
	}

	// res.RandomArrayId will hold Base64 encoded string
	res = entities.AssessmentUserResponse(assessment)

	return
}

func (s *service) UpdateUsersAssessmentStatus(ctx context.Context, req entities.UpdateUsersAssessmentStatusPayload) (err error) {
	newUserAssessment := entities.AssessmentUser{
		UsersId:        req.UsersId,
		AssessmentCode: req.AssessmentCode,
	}

	if err = newUserAssessment.ValidateAssessmentUserCode(); err != nil {
		return
	}

	if err = newUserAssessment.UpdateStatus(req.Status); err != nil {
		return
	}

	assessment, err := s.Repository.FindUsersAssessment(ctx, newUserAssessment.UsersId, newUserAssessment.AssessmentCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrAssessmentNotFound
		}
		return err
	}

	err = s.Repository.UpdateUsersAssessmentStatus(ctx, assessment.UsersId, assessment.AssessmentCode, newUserAssessment.Status)
	if err != nil {
		return err
	}

	return
}
