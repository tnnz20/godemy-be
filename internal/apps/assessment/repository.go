package assessment

import (
	"context"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/assessment/entities"
)

type Repository interface {
	CreateAssessment(ctx context.Context, assignment entities.Assessment) (err error)
	FindAssessmentsFiltered(ctx context.Context, usersId uuid.UUID) (assessments []entities.Assessment, err error)
	FindAssessmentByAssessmentCode(ctx context.Context, usersId uuid.UUID, assessmentCode string) (assessment entities.Assessment, err error)
	FindCoursesEnrollment(ctx context.Context, usersId uuid.UUID) (enrollment entities.Enrollment, err error)
}
