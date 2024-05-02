package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/teacher"
	"github.com/tnnz20/godemy-be/internal/apps/teacher/entities"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) teacher.Repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateCourse(ctx context.Context, course entities.Course) (err error) {
	query := `
	INSERT INTO courses (course_name, course_code, teacher_id)
	VALUES ($1, $2, $3)
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, course.CourseName, course.CourseCode, course.TeacherId)
	if err != nil {
		return err
	}

	return
}

func (r repository) FindTeacherIdByUserId(ctx context.Context, userId uuid.UUID) (teacher entities.Teacher, err error) {
	query := `
	SELECT id
	FROM teacher
	WHERE users_id = $1
	`

	err = r.db.QueryRowContext(ctx, query, userId).Scan(
		&teacher.ID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		return
	}

	return
}
