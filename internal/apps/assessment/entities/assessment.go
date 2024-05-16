package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type AssessmentResult struct {
	ID              uuid.UUID
	UsersId         uuid.UUID
	CoursesId       uuid.UUID
	AssessmentValue float32
	AssessmentCode  string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewAssessmentResult(userId, coursesId uuid.UUID, assessmentValue float32, assessmentCode string) AssessmentResult {
	return AssessmentResult{
		ID:              uuid.New(),
		UsersId:         userId,
		CoursesId:       coursesId,
		AssessmentValue: assessmentValue,
		AssessmentCode:  assessmentCode,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

const (
	AssessmentStatusCreated = "CREATED"
	AssessmentStatusOnGoing = "ON_GOING"
	AssessmentStatusDone    = "DONE"
)

var (
	AssessmentStatusMapping = map[uint8]string{
		1:  AssessmentStatusCreated,
		5:  AssessmentStatusOnGoing,
		10: AssessmentStatusDone,
	}
)

type AssessmentUser struct {
	ID             uuid.UUID
	UsersId        uuid.UUID
	AssessmentCode string
	RandomArrayId  []int
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewAssessmentUser(userId uuid.UUID, assessmentCode string, randomArrayId []int) AssessmentUser {
	return AssessmentUser{
		ID:             uuid.New(),
		UsersId:        userId,
		AssessmentCode: assessmentCode,
		RandomArrayId:  randomArrayId,
		Status:         AssessmentStatusMapping[1],
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (a AssessmentResult) Validate() (err error) {
	if err = a.ValidateAssessmentValue(); err != nil {
		return
	}

	if err = a.ValidateAssessmentCode(); err != nil {
		return
	}

	return
}

func (a AssessmentResult) ValidateAssessmentValue() (err error) {
	if a.AssessmentValue < 0 || a.AssessmentValue > 100 {
		return errs.ErrInvalidAssessmentValue
	}
	return
}

func (a AssessmentResult) ValidateAssessmentCode() (err error) {
	if a.AssessmentCode == "" {
		return errs.ErrAssessmentCodeRequired
	}
	return
}

func (au AssessmentUser) ValidateAssessmentUserCode() (err error) {
	if au.AssessmentCode == "" {
		return errs.ErrAssessmentCodeRequired
	}
	return
}

func (au *AssessmentUser) UpdateStatus(status uint8) (err error) {

	if status != 1 && status != 5 && status != 10 {
		return errs.ErrInvalidAssessmentStatus
	}
	au.Status = AssessmentStatusMapping[status]
	return
}
