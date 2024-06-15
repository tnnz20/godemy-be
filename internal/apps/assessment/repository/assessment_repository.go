package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
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

// CreateAssessmentResult is a function to create a new assessment result
func (r *repository) CreateAssessmentResult(ctx context.Context, assessment entities.AssessmentResult) (err error) {
	query := `
	INSERT INTO users_assessment_result (
		id,
		users_id,
		courses_id,
		assessment_value,
		assessment_code,
		status,
		created_at,
		updated_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
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
		assessment.Status,
		assessment.CreatedAt,
		assessment.UpdatedAt,
	)

	if err != nil {
		return err
	}
	return
}

// FindAssessments is a function to get all assessments by user id
func (r *repository) FindAssessments(ctx context.Context, usersId uuid.UUID) (assessments []entities.AssessmentResult, err error) {
	query := `
	SELECT 
		a.id, 
		a.users_id, 
		a.courses_id, 
		a.assessment_value, 
		a.assessment_code, 
		a.status,
		a.created_at, 
		a.updated_at
	FROM users_assessment_result AS a
	INNER JOIN (
		SELECT assessment_code, MAX(created_at) AS max_created_at
		FROM users_assessment_result
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
		var assessment entities.AssessmentResult
		err = rows.Scan(
			&assessment.ID,
			&assessment.UsersId,
			&assessment.CoursesId,
			&assessment.AssessmentValue,
			&assessment.AssessmentCode,
			&assessment.Status,
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

// FindAssessmentsFilteredByCode is a function to get assessment by user id and assessment code
func (r *repository) FindAssessmentsFilteredByCode(ctx context.Context, usersId uuid.UUID, assessmentCode string, model entities.AssessmentPagination) (assessments []entities.AssessmentResult, err error) {
	query := `
	SELECT 
		id, 
		users_id, 
		courses_id, 
		assessment_value, 
		assessment_code, 
		status,
		created_at, 
		updated_at
	FROM 
		users_assessment_result
	WHERE 
		users_id = $1 AND 
		assessment_code = $2
	ORDER BY 
		created_at DESC
	LIMIT $3 OFFSET $4
	`

	rows, err := r.db.QueryContext(ctx, query, usersId, assessmentCode, model.Limit, model.Offset)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var assessment entities.AssessmentResult
		err = rows.Scan(
			&assessment.ID,
			&assessment.UsersId,
			&assessment.CoursesId,
			&assessment.AssessmentValue,
			&assessment.AssessmentCode,
			&assessment.Status,
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

func (r *repository) FindTotalAssessmentsFilteredByCode(ctx context.Context, usersId uuid.UUID, assessmentCode string) (total int, err error) {
	query := `
	SELECT 
		COUNT(id)
	FROM 
		users_assessment_result
	WHERE 
		users_id = $1 AND 
		assessment_code = $2
	`

	err = r.db.QueryRowContext(ctx, query, usersId, assessmentCode).Scan(&total)
	if err != nil {
		return
	}

	return total, nil
}

func (r *repository) FindAssessmentsByCourseId(ctx context.Context, courseId uuid.UUID, name, assessmentCode, sort string, status uint8, model entities.AssessmentPagination) (assessments []entities.AssessmentUsersResult, err error) {

	var whereClause string

	if name != "" {
		whereClause = "u.name ILIKE $2"
	} else {
		whereClause = "ar.assessment_code = $2"
	}

	query := fmt.Sprintf(`
	SELECT
		ar.id,
		u.id,
		u.name,
		ar.courses_id,
		ar.assessment_value,
		ar.assessment_code,
		ar.status,
		ar.created_at
	FROM 
		users_assessment_result ar
	JOIN 
		users u ON ar.users_id = u.id
	WHERE 
		ar.courses_id = $1 AND
		%s
	ORDER BY 
		CASE
			WHEN ar.status = $3 THEN 1 
			ELSE 2 
		END,
		ar.created_at %s
	LIMIT $4 OFFSET $5
	`, whereClause, sort)

	wildName := "%" + name + "%"

	var rows *sql.Rows

	if name != "" {
		rows, err = r.db.QueryContext(ctx, query, courseId, wildName, status, model.Limit, model.Offset)
	} else {
		rows, err = r.db.QueryContext(ctx, query, courseId, assessmentCode, status, model.Limit, model.Offset)
	}

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var assessmentUsersResult entities.AssessmentUsersResult
		err = rows.Scan(
			&assessmentUsersResult.Id,
			&assessmentUsersResult.UsersId,
			&assessmentUsersResult.Name,
			&assessmentUsersResult.CoursesId,
			&assessmentUsersResult.AssessmentValue,
			&assessmentUsersResult.AssessmentCode,
			&assessmentUsersResult.Status,
			&assessmentUsersResult.CreatedAt,
		)

		if err != nil {
			return
		}

		assessments = append(assessments, assessmentUsersResult)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return
}

func (r *repository) FindTotalAssessmentsByCourseId(ctx context.Context, courseId uuid.UUID, name, assessmentCode string) (total int, err error) {

	// Construct the base query
	query := `
		SELECT 
			COUNT(ar.id)
		FROM 
			users_assessment_result ar
		JOIN 
			users u ON ar.users_id = u.id
		WHERE 
			ar.courses_id = $1 AND
		`

	// Append the appropriate WHERE clause based on the presence of the name parameter
	var args []interface{}
	args = append(args, courseId)

	if name != "" {
		query += "u.name ILIKE $2"
		args = append(args, "%"+name+"%")
	} else {
		query += "ar.assessment_code = $2"
		args = append(args, assessmentCode)
	}

	// Execute the query with the appropriate parameters
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return
	}

	return total, nil
}

// FindCoursesEnrollment is a function to get course by user id
func (r *repository) FindCoursesEnrollment(ctx context.Context, usersId uuid.UUID) (enrollment entities.Enrollment, err error) {
	query := `
	SELECT 
		id, 
		users_id, 
		courses_id, 
		progress, 
		created_at, 
		updated_at
	FROM 
		course_enrollment
	WHERE 
		users_id = $1
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

// CreateAssessmentUser is a function to create a new assessment task
func (r *repository) CreateUsersAssessment(ctx context.Context, userAssessment entities.AssessmentUser) (err error) {
	query := `
	INSERT INTO users_assessment (
		id,
		users_id,
		assessment_code,
		random_array_id,
		status,
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
		userAssessment.ID,
		userAssessment.UsersId,
		userAssessment.AssessmentCode,
		pq.Array(userAssessment.RandomArrayId),
		userAssessment.Status,
		userAssessment.CreatedAt,
		userAssessment.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return
}

// FindUserAssessment is a function to get assessment task by user id and assessment code
func (r *repository) FindUsersAssessment(ctx context.Context, usersId uuid.UUID, assessmentCode string) (userAssessment entities.AssessmentUser, err error) {
	query := `
	SELECT
		id,
		users_id,
		assessment_code,
		random_array_id,
		status,
		created_at,
		updated_at

	FROM users_assessment
	WHERE users_id = $1 
		AND assessment_code = $2 
		AND status != $3
	`

	err = r.db.QueryRowContext(ctx, query, usersId, assessmentCode, entities.AssessmentStatusDone).Scan(
		&userAssessment.ID,
		&userAssessment.UsersId,
		&userAssessment.AssessmentCode,
		&userAssessment.RandomArrayId,
		&userAssessment.Status,
		&userAssessment.CreatedAt,
		&userAssessment.UpdatedAt,
	)

	if err != nil {
		return
	}

	return
}

// UpdateUserAssessmentStatus is a function to update assessment task status by user id and assessment code
func (r *repository) UpdateUsersAssessmentStatus(ctx context.Context, usersId uuid.UUID, assessmentCode string, status string) (err error) {
	query := `
	UPDATE users_assessment
	SET status = $1
	WHERE users_id = $2 AND assessment_code = $3
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, status, usersId, assessmentCode)
	if err != nil {
		return err
	}

	return
}
