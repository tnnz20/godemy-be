package assessment

import (
	"context"

	"github.com/tnnz20/godemy-be/internal/apps/assessment/entities"
)

type Service interface {
	CreateAssessmentResult(ctx context.Context, req entities.CreateAssessmentPayload) (err error)
	GetAssessmentsResult(ctx context.Context, req entities.GetAssessmentPayload) (res []entities.AssessmentResponse, err error)
	GetFilteredAssessmentResult(ctx context.Context, req entities.GetAssessmentResultWithPaginationPayload) (res []entities.AssessmentResponse, err error)
	GetTotalFilteredAssessmentResult(ctx context.Context, req entities.GetAssessmentResultWithPaginationPayload) (res entities.AssessmentTotalResponse, err error)
	GetAssessmentsResultUsers(ctx context.Context, req entities.GetAssessmentResultsByCourseIdPayload) (res []entities.AssessmentResultUsersResponse, err error)
	GetTotalAssessmentsResultUsers(ctx context.Context, req entities.GetAssessmentResultsByCourseIdPayload) (res entities.AssessmentTotalResponse, err error)
	CreateUsersAssessment(ctx context.Context, req entities.CreateUsersAssessmentPayload) (err error)
	GetUsersAssessment(ctx context.Context, req entities.GetUsersAssessmentPayload) (res entities.AssessmentUserResponse, err error)
	UpdateUsersAssessmentStatus(ctx context.Context, req entities.UpdateUsersAssessmentStatusPayload) (err error)
}
