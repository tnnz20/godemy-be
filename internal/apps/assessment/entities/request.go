package entities

import "github.com/google/uuid"

type CreateAssessmentRequest struct {
	UsersId         uuid.UUID `json:"users_id"`
	AssessmentValue float32   `json:"assessment_value"`
	AssessmentCode  string    `json:"assessment_code"`
}

type GetAssessmentRequest struct {
	UsersId uuid.UUID `json:"users_id"`
}

type GetAssessmentByAssessmentCodeRequest struct {
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `json:"assessment_code"`
}

type CreateUsersAssessmentRequest struct {
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `json:"assessment_code"`
	RandomArrayId  []uint8   `json:"random_array_id"`
}

type GetUsersAssessmentRequest struct {
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `json:"assessment_code"`
}

type UpdateUsersAssessmentStatusRequest struct {
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `json:"assessment_code"`
	Status         uint8     `json:"status"`
}
