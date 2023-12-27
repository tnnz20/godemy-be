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
	ID uuid.UUID `json:"id" validate:"required,uuid4"`
}

type GetTeacherIdByUserIdResponse struct {
	ID     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
}

type CreateClassRequest struct {
	TeacherId uuid.UUID `json:"teacher_id" validate:"required,uuid4"`
	ClassName string    `json:"class_name" validate:"required,min=3,max=50"`
}

type CreateClassResponse struct {
	ID        uuid.UUID `json:"id"`
	TeacherId uuid.UUID `json:"teacher_id"`
	ClassName string    `json:"class_name"`
}

// TODO: create repository for get class by teacherId
type Repository interface {
	GetTeacherIdByUserId(ctx context.Context, id *uuid.UUID) (*Teacher, error)
	CreateClass(ctx context.Context, class *Class) (*Class, error)
}

// TODO: create service for get class by teacherId
type Service interface {
	GetTeacherIdByUserId(ctx context.Context, req *GetTeacherIdByUserIdRequest) (*GetTeacherIdByUserIdResponse, error)
	CreateClass(ctx context.Context, req *CreateClassRequest) (*CreateClassResponse, error)
}
