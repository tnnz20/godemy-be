package student

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) GetStudentByUserId(ctx context.Context, userId *uuid.UUID) (*Student, error) {
	student := Student{}

	query := "SELECT id, users_id, class_id, threshold FROM student where users_id = $1"

	if err := r.db.QueryRowContext(ctx, query, userId).Scan(&student.ID, &student.UsersId,
		&student.ClassId, &student.Threshold); err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *repository) IncrementThreshold(ctx context.Context, userId *uuid.UUID, threshold *int) error {
	query := "UPDATE student SET threshold = $1 WHERE id = $2"

	if _, err := r.db.ExecContext(ctx, query, threshold, userId); err != nil {
		return err
	}

	return nil
}

func (r *repository) InsertAssessment(ctx context.Context, studentId *uuid.UUID, assessment *Assessment) (*Student, *Assessment, error) {
	student := Student{}

	query := "INSERT INTO assessment (student_id, assessment_value, code_assessment) VALUES ($1, $2, $3)"
	if _, err := r.db.ExecContext(ctx, query, studentId, assessment.AssessmentValue,
		assessment.CodeAssessment); err != nil {
		return nil, nil, err
	}

	student.ID = *studentId

	return &student, assessment, nil
}
