package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type Assessment struct {
	ID              uuid.UUID
	UsersId         uuid.UUID
	CoursesId       uuid.UUID
	AssessmentValue float32
	AssessmentCode  string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewAssessment(userId, coursesId uuid.UUID, assessmentValue float32, assessmentCode string) Assessment {
	return Assessment{
		ID:              uuid.New(),
		UsersId:         userId,
		CoursesId:       coursesId,
		AssessmentValue: assessmentValue,
		AssessmentCode:  assessmentCode,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (a Assessment) Validate() (err error) {
	if err = a.ValidateAssessmentValue(); err != nil {
		return
	}

	if err = a.ValidateAssessmentCode(); err != nil {
		return
	}

	return
}

func (a Assessment) ValidateAssessmentValue() (err error) {
	if a.AssessmentValue < 0 || a.AssessmentValue > 100 {
		return errs.ErrInvalidAssessmentValue
	}
	return
}

func (a Assessment) ValidateAssessmentCode() (err error) {
	if a.AssessmentCode == "" {
		return errs.ErrAssessmentCodeRequired
	}

	return
}
