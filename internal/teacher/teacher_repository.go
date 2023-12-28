package teacher

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) GetTeacherIdByUserId(ctx context.Context, id *uuid.UUID) (*Teacher, error) {
	var teacher = &Teacher{}

	query := "SELECT id FROM teacher WHERE users_id = $1"
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&teacher.ID); err != nil {
		return nil, err
	}

	teacher.UserId = *id

	return teacher, nil
}

func (r *repository) CreateClass(ctx context.Context, class *Class) (*Class, error) {
	var lastInsertedID string

	query := "INSERT INTO class (teacher_id, class_name) VALUES($1, $2) returning id"

	if err := r.db.QueryRowContext(ctx, query, class.TeacherId, class.ClassName).Scan(&lastInsertedID); err != nil {
		return nil, err
	}

	parseUUID, _ := uuid.Parse(lastInsertedID)

	class.ID = parseUUID

	return class, nil
}

func (r *repository) GetAllClassByTeacherId(ctx context.Context, teacherId *uuid.UUID) (*[]Class, error) {
	query := "SELECT id, teacher_id, class_name FROM class WHERE teacher_id = $1"

	rows, err := r.db.QueryContext(ctx, query, teacherId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []Class

	for rows.Next() {
		var class Class
		if err := rows.Scan(&class.ID, &class.TeacherId, &class.ClassName); err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &classes, nil
}
