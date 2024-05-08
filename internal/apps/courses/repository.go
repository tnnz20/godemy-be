package courses

import (
	"context"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/courses/entities"
)

// TODO: update progress enrollment
// Repository is the interface that provides the repository methods for courses.
type Repository interface {
	CreateCourse(ctx context.Context, course entities.Courses) (err error)
	FindCourseByCourseCode(ctx context.Context, courseCode string) (course entities.Courses, err error)
	FindCoursesByUsersIdWithPagination(ctx context.Context, usersId uuid.UUID, model entities.CoursesPagination) (courses []entities.Courses, err error)
	InsertCourseEnrollment(ctx context.Context, enrollment entities.Enrollment) (err error)
	FindCourseEnrollmentByUsersId(ctx context.Context, usersId uuid.UUID) (enrollments entities.Enrollment, err error)
}
