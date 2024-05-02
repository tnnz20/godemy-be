package teacher

import (
	"context"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/teacher/entities"
)

type Repository interface {
	FindTeacherIdByUserId(ctx context.Context, userId uuid.UUID) (teacher entities.Teacher, err error)
	CreateCourse(ctx context.Context, course entities.Course) (err error)

	// FindClassByTeacherId(ctx context.Context, teacherId uuid.UUID) (class entities.Class, err error)
}
