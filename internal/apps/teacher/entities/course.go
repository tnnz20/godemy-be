package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type Course struct {
	ID         uuid.UUID
	TeacherId  uuid.UUID
	CourseName string
	CourseCode string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (c Course) Validate() (err error) {
	if err := c.ValidateCourseName(); err != nil {
		return err
	}

	if err := c.ValidateCourseCode(); err != nil {
		return err
	}

	return
}

func (c Course) ValidateCourseName() (err error) {
	if len(c.CourseName) < 3 {
		return errs.ErrInvalidCourseNameLength
	}
	return
}

func (c Course) ValidateCourseCode() (err error) {
	if len(c.CourseCode) < 10 {
		return errs.ErrInvalidCourseCodeLength
	}
	return
}

func (c Course) IsCourseCodeExist() bool {
	return c.ID != uuid.Nil
}
