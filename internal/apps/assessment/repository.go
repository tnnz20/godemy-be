package assessment

import (
	"context"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/assessment/entities"
)

type Repository interface {
	CreateAssessmentResult(ctx context.Context, assessment entities.AssessmentResult) (err error)
	FindAssessments(ctx context.Context, usersId uuid.UUID) (assessments []entities.AssessmentResult, err error)
	FindAssessmentsFilteredByCode(ctx context.Context, usersId uuid.UUID, assessmentCode string, model entities.AssessmentPagination) (assessments []entities.AssessmentResult, err error)
	FindTotalAssessmentsFilteredByCode(ctx context.Context, usersId uuid.UUID, assessmentCode string) (total int, err error)
	FindAssessmentsUsersByCode(ctx context.Context, courseId uuid.UUID, assessmentCode string, model entities.AssessmentPagination) (assessments []entities.AssessmentUsersResult, err error)
	FindCoursesEnrollment(ctx context.Context, usersId uuid.UUID) (enrollment entities.Enrollment, err error)
	CreateUsersAssessment(ctx context.Context, userAssessment entities.AssessmentUser) (err error)
	FindUsersAssessment(ctx context.Context, usersId uuid.UUID, assessmentCode string) (userAssessment entities.AssessmentUser, err error)
	UpdateUsersAssessmentStatus(ctx context.Context, usersId uuid.UUID, assessmentCode string, status string) (err error)
}
