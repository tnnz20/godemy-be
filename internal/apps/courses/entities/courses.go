package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type Courses struct {
	ID         uuid.UUID
	UsersId    uuid.UUID
	CourseName string
	CourseCode string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CoursesPagination struct {
	Limit  int
	Offset int
}

func NewCourses(usersId uuid.UUID, courseName, courseCode string) Courses {
	return Courses{
		ID:         uuid.New(),
		UsersId:    usersId,
		CourseName: courseName,
		CourseCode: courseCode,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func NewCoursesPagination(limit, offset int) CoursesPagination {
	return CoursesPagination{
		Limit:  limit,
		Offset: offset,
	}
}

func (c Courses) Validate() (err error) {
	if err := c.ValidateCourseName(); err != nil {
		return err
	}

	if err := c.ValidateCourseCode(); err != nil {
		return err
	}

	return
}

func (c Courses) ValidateCourseName() (err error) {
	if c.CourseName == "" {
		return errs.ErrCourseNameRequired
	}

	if len(c.CourseName) < 3 {
		return errs.ErrInvalidCourseNameLength
	}
	return
}

func (c Courses) ValidateCourseCode() (err error) {
	if c.CourseCode == "" {
		return errs.ErrCourseCodeRequired
	}

	if len(c.CourseCode) != 10 {
		return errs.ErrInvalidCourseCodeLength
	}
	return
}

func (c Courses) IsCourseCodeExist() bool {
	return c.ID != uuid.Nil
}
