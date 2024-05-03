package teacher

import (
	"context"

	"github.com/tnnz20/godemy-be/internal/apps/teacher/entities"
)

type Service interface {
	GetTeacherIdByUserId(ctx context.Context, req entities.GetTeacherIdByUserIdPayload) (res entities.GetTeacherIdByUserIdResponse, err error)
	CreateCourse(ctx context.Context, req entities.CreateCoursePayload) (err error)
	GetCourseByCourseCode(ctx context.Context, req entities.GetCourseByCourseCodePayload) (res entities.GetCourseByCourseCodeResponse, err error)
	GetCourseByTeacherId(ctx context.Context, req entities.GetCourseByTeacherIdPayload) (res entities.GetCourseByTeacherIdResponse, err error)
}
