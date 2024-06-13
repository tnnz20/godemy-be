package assessment

import (
	"context"

	"github.com/tnnz20/godemy-be/internal/apps/assessment/entities"
)

type Service interface {
	CreateAssessmentResult(ctx context.Context, req entities.CreateAssessmentRequest) (err error)
	GetAssessmentsResult(ctx context.Context, req entities.GetAssessmentRequest) (res []entities.AssessmentResponse, err error)
	GetFilteredAssessmentResult(ctx context.Context, req entities.GetAssessmentResultByAssessmentCodeRequest) (res []entities.AssessmentResponse, err error)
	GetTotalFilteredAssessmentResult(ctx context.Context, req entities.GetAssessmentResultByAssessmentCodePayload) (res entities.AssessmentTotalResponse, err error)
	CreateUsersAssessment(ctx context.Context, req entities.CreateUsersAssessmentRequest) (err error)
	GetUsersAssessment(ctx context.Context, req entities.GetUsersAssessmentRequest) (res entities.AssessmentUserResponse, err error)
	UpdateUsersAssessmentStatus(ctx context.Context, req entities.UpdateUsersAssessmentStatusRequest) (err error)
}
