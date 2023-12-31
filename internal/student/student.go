package student

import (
	"context"

	"github.com/google/uuid"
)

type Student struct {
	ID        uuid.UUID `json:"id"`
	UsersId   uuid.UUID `json:"users_id"`
	ClassId   uuid.UUID `json:"class_id"`
	Threshold int       `json:"threshold"`
}

type Assessment struct {
	ID              uuid.UUID `json:"id"`
	StudentId       uuid.UUID `json:"student_id"`
	AssessmentValue int       `json:"assessment_value"`
	CodeAssessment  string    `json:"code_assessment"`
}

type GetStudentByUserIdRequest struct {
	UsersId uuid.UUID `json:"users_id" validate:"required,uuid4"`
}

type GetStudentByUserIdResponse struct {
	ID        uuid.UUID `json:"id"`
	UsersId   uuid.UUID `json:"users_id"`
	ClassId   uuid.UUID `json:"class_id"`
	Threshold int       `json:"threshold"`
}

type Repository interface {
	GetStudentByUserId(ctx context.Context, userId *uuid.UUID) (*Student, error)
}

type Service interface {
	GetStudentByUserId(ctx context.Context, req *GetStudentByUserIdRequest) (*GetStudentByUserIdResponse, error)
}
