package teacher

import (
	"context"

	"github.com/google/uuid"
)

type Teacher struct {
	ID     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
}

type Class struct {
	ID        uuid.UUID `json:"id"`
	TeacherId uuid.UUID `json:"teacher_id"`
	ClassName string    `json:"class_name"`
}

type GetTeacherIdByUserIdRequest struct {
	UserId uuid.UUID `json:"user_id" validate:"required,uuid4"`
}

type GetTeacherIdByUserIdResponse struct {
	ID     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
}

type GetClassByTeacherIdRequest struct {
	TeacherId uuid.UUID `json:"id" validate:"required,uuid4"`
}

type GetClassByTeacherIdResponse struct {
	ID        uuid.UUID `json:"id"`
	TeacherId uuid.UUID `json:"teacher_id"`
	ClassName string    `json:"class_name"`
}

// TODO: List student by class (class_name query)
type Repository interface {
	GetTeacherIdByUserId(ctx context.Context, userId *uuid.UUID) (*Teacher, error)
	GetAllClassByTeacherId(ctx context.Context, teacherId *uuid.UUID) (*[]Class, error)
}

type Service interface {
	GetTeacherIdByUserId(ctx context.Context, req *GetTeacherIdByUserIdRequest) (*GetTeacherIdByUserIdResponse, error)
	GetAllClassByTeacherId(ctx context.Context, req *GetClassByTeacherIdRequest) (*[]GetClassByTeacherIdResponse, error)
}
