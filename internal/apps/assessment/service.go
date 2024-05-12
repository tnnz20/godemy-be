package assessment

import (
	"context"

	"github.com/tnnz20/godemy-be/internal/apps/assessment/entities"
)

type Service interface {
	CreateAssessment(ctx context.Context, req entities.CreateAssessmentRequest) (err error)
	GetAssessments(ctx context.Context, req entities.GetAssessmentRequest) (res []entities.AssessmentResponse, err error)
	GetAssessmentByAssessmentCode(ctx context.Context, req entities.GetAssessmentByAssessmentCodeRequest) (res entities.AssessmentResponse, err error)
}
