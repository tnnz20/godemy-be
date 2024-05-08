package courses

import (
	"context"

	"github.com/tnnz20/godemy-be/internal/apps/courses/entities"
)

// TODO: update progress enrollment

// Service is the interface that provides the service methods for courses.
type Service interface {
	CreateCourse(ctx context.Context, req entities.CreateCoursePayload) (err error)
	GetCourseByCourseCode(ctx context.Context, req entities.GetCourseByCourseCodePayload) (res entities.CourseResponse, err error)
	GetCoursesByUsersIdWithPagination(ctx context.Context, req entities.GetCoursesByUsersIdWithPaginationPayload) (res []entities.CourseResponse, err error)
	EnrollCourse(ctx context.Context, req entities.EnrollCoursePayload) (err error)
	GetCourseEnrollmentByUsersId(ctx context.Context, req entities.GetCourseEnrollmentByUsersIdPayload) (res entities.CourseEnrollmentResponse, err error)
}
