package courses

import (
	"context"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/courses/entities"
)

// Repository is the interface that provides the repository methods for courses.
type Repository interface {
	CreateCourse(ctx context.Context, course entities.Courses) (err error)
	FindCourseByCourseCode(ctx context.Context, courseCode string) (course entities.Courses, err error)
	FindCoursesByUsersIdWithPagination(ctx context.Context, usersId uuid.UUID, courseName string, model entities.CoursesPagination) (courses []entities.Courses, err error)
	FindCoursesByUsersId(ctx context.Context, usersId uuid.UUID, courseName string) (courses []entities.Courses, err error)
	FindTotalCoursesByUsersId(ctx context.Context, usersId uuid.UUID, courseName string) (total int, err error)
	InsertCourseEnrollment(ctx context.Context, enrollment entities.Enrollment) (err error)
	FindCourseEnrollmentByUsersId(ctx context.Context, usersId uuid.UUID) (enrollments entities.Enrollment, err error)
	UpdateEnrollmentProgress(ctx context.Context, enrollment entities.Enrollment) (err error)
	FindEnrolledUsersByCourseId(ctx context.Context, courseId uuid.UUID, name string, model entities.CoursesPagination) (courses []entities.EnrolledUsersResponse, err error)
	FindTotalEnrolledUsersByCourseId(ctx context.Context, courseId uuid.UUID, name string) (total int, err error)
}
