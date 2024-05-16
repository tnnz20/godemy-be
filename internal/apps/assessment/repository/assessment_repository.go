package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/assessment"
	"github.com/tnnz20/godemy-be/internal/apps/assessment/entities"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) assessment.Repository {
	return &repository{
		db: db,
	}
}

// CreateAssessment is a function to create a new assessment
func (r *repository) CreateAssessment(ctx context.Context, assessment entities.Assessment) (err error) {
	query := `
	INSERT INTO assessment (
		id,
		users_id,
		courses_id,
		assessment_value,
		assessment_code,
		created_at,
		updated_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		assessment.ID,
		assessment.UsersId,
		assessment.CoursesId,
		assessment.AssessmentValue,
		assessment.AssessmentCode,
		assessment.CreatedAt,
		assessment.UpdatedAt,
	)

	if err != nil {
		return err
	}
	return
}

// FindAssessments is a function to get all assessments by user id
func (r *repository) FindAssessmentsFiltered(ctx context.Context, usersId uuid.UUID) (assessments []entities.Assessment, err error) {
	query := `
	SELECT 
		a.id, 
		a.users_id, 
		a.courses_id, 
		a.assessment_value, 
		a.assessment_code, 
		a.created_at, 
		a.updated_at
	FROM assessment AS a
	INNER JOIN (
		SELECT assessment_code, MAX(created_at) AS max_created_at
		FROM assessment
		GROUP BY assessment_code
	) b ON a.assessment_code = b.assessment_code AND a.created_at = b.max_created_at
	WHERE users_id = $1
	ORDER BY a.created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, usersId)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var assessment entities.Assessment
		err = rows.Scan(
			&assessment.ID,
			&assessment.UsersId,
			&assessment.CoursesId,
			&assessment.AssessmentValue,
			&assessment.AssessmentCode,
			&assessment.CreatedAt,
			&assessment.UpdatedAt,
		)

		if err != nil {
			return
		}

		assessments = append(assessments, assessment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return
}

// FindAssessmentByAssessmentCode is a function to get assessment by user id and assessment code
func (r *repository) FindAssessmentByAssessmentCode(ctx context.Context, usersId uuid.UUID, assessmentCode string) (assessment entities.Assessment, err error) {
	query := `
	SELECT 
		id, 
		users_id, 
		courses_id, 
		assessment_value, 
		assessment_code, 
		created_at, 
		updated_at
	FROM assessment
	WHERE users_id = $1 AND assessment_code = $2
	`

	err = r.db.QueryRowContext(ctx, query, usersId, assessmentCode).Scan(
		&assessment.ID,
		&assessment.UsersId,
		&assessment.CoursesId,
		&assessment.AssessmentValue,
		&assessment.AssessmentCode,
		&assessment.CreatedAt,
		&assessment.UpdatedAt,
	)

	if err != nil {
		return
	}

	return
}

// FindCoursesEnrollment is a function to get course by user id
func (r *repository) FindCoursesEnrollment(ctx context.Context, usersId uuid.UUID) (enrollment entities.Enrollment, err error) {
	query := `
	SELECT id, users_id, courses_id, progress, created_at, updated_at
	FROM course_enrollment
	WHERE users_id = $1
	`

	err = r.db.QueryRowContext(ctx, query, usersId).Scan(
		&enrollment.ID,
		&enrollment.UsersId,
		&enrollment.CoursesId,
		&enrollment.Progress,
		&enrollment.CreatedAt,
		&enrollment.UpdatedAt,
	)

	if err != nil {
		return
	}

	return
}
