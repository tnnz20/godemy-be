package assessment

import (
	"context"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/assessment/entities"
)

type Repository interface {
	CreateAssessmentResult(ctx context.Context, assessment entities.AssessmentResult) (err error)
	FindAssessmentsFiltered(ctx context.Context, usersId uuid.UUID) (assessments []entities.AssessmentResult, err error)
	FindAssessmentByAssessmentCode(ctx context.Context, usersId uuid.UUID, assessmentCode string) (assessment entities.AssessmentResult, err error)
	FindCoursesEnrollment(ctx context.Context, usersId uuid.UUID) (enrollment entities.Enrollment, err error)
	CreateUsersAssessment(ctx context.Context, userAssessment entities.AssessmentUser) (err error)
	FindUsersAssessment(ctx context.Context, usersId uuid.UUID, assessmentCode string) (userAssessment entities.AssessmentUser, err error)
	UpdateUsersAssessmentStatus(ctx context.Context, usersId uuid.UUID, assessmentCode string, status string) (err error)
}
