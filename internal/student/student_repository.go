package student

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type DBTX interface {
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
