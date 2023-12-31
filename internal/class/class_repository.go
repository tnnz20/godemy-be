package class

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

func (r *repository) CreateClass(ctx context.Context, class *Class) (*Class, error) {
	var lastInsertedID string

	query := "INSERT INTO class (teacher_id, class_name) VALUES($1, $2) returning id"

	if err := r.db.QueryRowContext(ctx, query, class.TeacherId,
		class.ClassName).Scan(&lastInsertedID); err != nil {
		return nil, err
	}

	parseUUID, _ := uuid.Parse(lastInsertedID)
	class.ID = parseUUID

	return class, nil
}

func (r *repository) GetAllClass(ctx context.Context) (*[]Class, error) {
	query := "SELECT id, teacher_id, class_name FROM class"

	rows, err := r.db.QueryContext(ctx, query)
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

func (r *repository) UpdateStudentClass(ctx context.Context, classId *uuid.UUID, studentId *uuid.UUID) error {
	query := "UPDATE student SET class_id = $1 WHERE id = $2"
	if _, err := r.db.ExecContext(ctx, query, classId, studentId); err != nil {
		return err
	}

	return nil
}
