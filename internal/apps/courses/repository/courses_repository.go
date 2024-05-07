package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/courses"
	"github.com/tnnz20/godemy-be/internal/apps/courses/entities"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) courses.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateCourse(ctx context.Context, course entities.Courses) (err error) {
	query := `
	INSERT INTO courses (
		users_id, 
		course_name, 
		course_code
	)
	VALUES ($1, $2, $3)
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, course.UsersId, course.CourseName, course.CourseCode)
	if err != nil {
		return err
	}

	return
}

func (r *repository) FindCourseByCourseCode(ctx context.Context, courseCode string) (course entities.Courses, err error) {
	query := `
	SELECT id, users_id, course_name, course_code, created_at, updated_at
	FROM courses
	WHERE course_code = $1
	`

	err = r.db.QueryRowContext(ctx, query, courseCode).Scan(
		&course.ID,
		&course.UsersId,
		&course.CourseName,
		&course.CourseCode,
		&course.CreatedAt,
		&course.UpdatedAt,
	)

	if err != nil {
		return
	}

	return
}

func (r *repository) FindCoursesByUsersIdWithPagination(ctx context.Context, usersId uuid.UUID, model entities.CoursesPagination) (courses []entities.Courses, err error) {
	query := `
	SELECT id, users_id, course_name, course_code, created_at, updated_at
	FROM courses
	WHERE users_id = $1
	ORDER BY created_at DESC
	LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, usersId, model.Limit, model.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var c entities.Courses

		err = rows.Scan(
			&c.ID,
			&c.UsersId,
			&c.CourseName,
			&c.CourseCode,
			&c.CreatedAt,
			&c.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		courses = append(courses, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return
}
