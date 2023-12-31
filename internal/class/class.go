package class

import (
	"context"

	"github.com/google/uuid"
)

type Class struct {
	ID        uuid.UUID `json:"id"`
	TeacherId uuid.UUID `json:"teacher_id"`
	ClassName string    `json:"class_name"`
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

type GetAllClassResponse struct {
	ID        uuid.UUID `json:"id"`
	TeacherId uuid.UUID `json:"teacher_id"`
	ClassName string    `json:"class_name"`
}

type UpdateStudentClassRequest struct {
	ClassId   uuid.UUID `json:"class_id" validate:"required,uuid4"`
	StudentId uuid.UUID `json:"student_id" validate:"required,uuid4"`
}

type Repository interface {
	CreateClass(ctx context.Context, class *Class) (*Class, error)
	GetAllClass(ctx context.Context) (*[]Class, error)
	UpdateStudentClass(ctx context.Context, classId *uuid.UUID, studentId *uuid.UUID) error
}

type Service interface {
	CreateClass(ctx context.Context, req *CreateClassRequest) (*CreateClassResponse, error)
	GetAllClass(ctx context.Context) (*[]GetAllClassResponse, error)
	UpdateStudentClass(ctx context.Context, req *UpdateStudentClassRequest) error
}
