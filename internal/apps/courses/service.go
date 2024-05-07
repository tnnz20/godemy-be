package courses

import (
	"context"

	"github.com/tnnz20/godemy-be/internal/apps/courses/entities"
)

// TODO: add insert course enrollment service

// Service is the interface that provides the service methods for courses.
type Service interface {
	CreateCourse(ctx context.Context, req entities.CreateCoursePayload) (err error)
	GetCourseByCourseCode(ctx context.Context, req entities.GetCourseByCourseCodePayload) (res entities.CourseResponse, err error)
	GetCoursesByUsersIdWithPagination(ctx context.Context, req entities.GetCoursesByUsersIdWithPaginationPayload) (res []entities.CourseResponse, err error)
}
