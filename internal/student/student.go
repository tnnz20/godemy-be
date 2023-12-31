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

type IncrementThresholdRequest struct {
	UsersId   uuid.UUID `json:"users_id" validate:"required,uuid4"`
	Threshold int       `json:"threshold" validate:"required"`
}

type InsertAssessmentRequest struct {
	UsersId         uuid.UUID `json:"users_id" validate:"required,uuid4"`
	AssessmentValue int       `json:"assessment_value" validate:"required,min=0,max=100,number"`
	CodeAssessment  string    `json:"code_assessment" validate:"required,max=20"`
}

type InsertAssessmentResponse struct {
	StudentId       uuid.UUID `json:"student_id"`
	AssessmentValue int       `json:"assessment_value"`
	CodeAssessment  string    `json:"code_assessment"`
}

type Repository interface {
	GetStudentByUserId(ctx context.Context, userId *uuid.UUID) (*Student, error)
	IncrementThreshold(ctx context.Context, studentId *uuid.UUID, threshold *int) error
	InsertAssessment(ctx context.Context, studentId *uuid.UUID, assessment *Assessment) (*Student, *Assessment, error)
}

type Service interface {
	GetStudentByUserId(ctx context.Context, req *GetStudentByUserIdRequest) (*GetStudentByUserIdResponse, error)
	IncrementThreshold(ctx context.Context, req *IncrementThresholdRequest) error
	InsertAssessment(ctx context.Context, req *InsertAssessmentRequest) (*InsertAssessmentResponse, error)
}
