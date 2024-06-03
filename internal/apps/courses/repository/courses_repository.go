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
		id,
		users_id, 
		course_name, 
		course_code,
		created_at,
		updated_at
	)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		course.ID,
		course.UsersId,
		course.CourseName,
		course.CourseCode,
		course.CreatedAt,
		course.UpdatedAt,
	)
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

func (r *repository) FindTotalCoursesByUsersId(ctx context.Context, usersId uuid.UUID) (total int, err error) {
	query := `
	SELECT COUNT(id)
	FROM courses
	WHERE users_id = $1
	`

	err = r.db.QueryRowContext(ctx, query, usersId).Scan(&total)
	if err != nil {
		return
	}

	return
}

func (r *repository) InsertCourseEnrollment(ctx context.Context, enrollment entities.Enrollment) (err error) {
	query := `
	INSERT INTO course_enrollment (
		id, 
		users_id, 
		courses_id, 
		progress, 
		created_at, 
		updated_at
	) 
	VALUES ($1, $2, $3, $4, $5, $6)`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		enrollment.ID,
		enrollment.UsersId,
		enrollment.CoursesId,
		enrollment.Progress,
		enrollment.CreatedAt,
		enrollment.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return
}

func (r *repository) FindCourseEnrollmentByUsersId(ctx context.Context, usersId uuid.UUID) (enrollments entities.Enrollment, err error) {
	query := `
	SELECT id, users_id, courses_id, progress, created_at, updated_at
	FROM course_enrollment
	WHERE users_id = $1
	`

	err = r.db.QueryRowContext(ctx, query, usersId).Scan(
		&enrollments.ID,
		&enrollments.UsersId,
		&enrollments.CoursesId,
		&enrollments.Progress,
		&enrollments.CreatedAt,
		&enrollments.UpdatedAt,
	)

	if err != nil {
		return
	}

	return
}

func (r *repository) UpdateEnrollmentProgress(ctx context.Context, enrollment entities.Enrollment) (err error) {
	query := `
	UPDATE course_enrollment
	SET progress = $1, updated_at = $2
	WHERE id = $3
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, enrollment.Progress, enrollment.UpdatedAt, enrollment.ID)
	if err != nil {
		return err
	}

	return
}

func (r *repository) FindEnrolledUsersByCourseId(ctx context.Context, courseId uuid.UUID, name string, model entities.CoursesPagination) (courses []entities.EnrolledUsersResponse, err error) {
	query := `
	SELECT 
		u.id, 
		c.course_name,
		u.name,
		ce.progress,
		ce.updated_at
	FROM 
		users AS u
	JOIN 
		course_enrollment AS ce ON u.id = ce.users_id
	JOIN
		courses AS c ON ce.courses_id = c.id
	WHERE
		c.id = $1 AND
		(u.name ILIKE $2)
	LIMIT $3 OFFSET $4
	`

	wildcardName := "%" + name + "%"

	rows, err := r.db.QueryContext(ctx, query, courseId, wildcardName, model.Limit, model.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var c entities.EnrolledUsersResponse

		err = rows.Scan(
			&c.ID,
			&c.CourseName,
			&c.Name,
			&c.Progress,
			&c.UpdateAt,
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
