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

type GetAssessmentResultByAssessmentCodeRequest struct {
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `query:"assessment_code"`
	ModelPaginationPayload
}

type CreateUsersAssessmentRequest struct {
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `json:"assessment_code"`
	RandomArrayId  []uint8   `json:"random_array_id"`
}

type GetUsersAssessmentRequest struct {
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `query:"assessment_code"`
}

type UpdateUsersAssessmentStatusRequest struct {
	UsersId        uuid.UUID `json:"users_id"`
	AssessmentCode string    `json:"assessment_code"`
	Status         uint8     `json:"status"`
}

type ModelPaginationPayload struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

func (p *ModelPaginationPayload) GenerateDefaultValue() ModelPaginationPayload {
	// limit rows
	if p.Limit <= 0 {
		p.Limit = 10
	}

	// skip rows
	if p.Offset <= 0 {
		p.Offset = 0
	}

	return *p
}
