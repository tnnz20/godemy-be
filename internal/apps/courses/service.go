package courses

import (
	"context"

	"github.com/tnnz20/godemy-be/internal/apps/courses/entities"
)

// Service is the interface that provides the service methods for courses.
type Service interface {
	CreateCourse(ctx context.Context, req entities.CreateCoursePayload) (err error)
	GetCourseByCourseCode(ctx context.Context, req entities.GetCourseByCourseCodePayload) (res entities.CourseResponse, err error)
	GetCoursesByUsersIdWithPagination(ctx context.Context, req entities.GetCoursesByUsersIdWithPaginationPayload) (res []entities.CourseResponse, err error)
	GetCoursesByUsersId(ctx context.Context, req entities.GetCoursesByUsersIdPayload) (res []entities.CourseResponse, err error)
	GetTotalCourses(ctx context.Context, req entities.GetTotalCoursesByUsersIdPayload) (res entities.CoursesLengthResponse, err error)
	EnrollCourse(ctx context.Context, req entities.EnrollCoursePayload) (err error)
	GetCourseEnrollmentByUsersId(ctx context.Context, req entities.GetCourseEnrollmentByUsersIdPayload) (res entities.CourseEnrollmentResponse, err error)
	UpdateProgressCourseEnrollment(ctx context.Context, req entities.UpdateEnrollmentProgressPayload) (err error)
	GetEnrolledUsersByCourseId(ctx context.Context, req entities.GetEnrolledUsersByCourseIdPayload) (res []entities.EnrolledUsersResponse, err error)
	GetTotalEnrolledUsersByCourseId(ctx context.Context, req entities.GetTotalEnrolledUsersByCourseIdPayload) (res entities.EnrolledUsersLengthResponse, err error)
}
