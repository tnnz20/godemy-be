package teacher

import (
	"context"

	"github.com/tnnz20/godemy-be/internal/apps/teacher/entities"
)

type Service interface {
	GetTeacherIdByUserId(ctx context.Context, req entities.GetTeacherIdByUserIdRequest) (res entities.GetTeacherIdByUserIdResponse, err error)
	CreateCourse(ctx context.Context, req entities.CreateCourseRequest) (err error)
}
