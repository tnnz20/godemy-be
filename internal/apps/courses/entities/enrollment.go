package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type Enrollment struct {
	ID        uuid.UUID
	UsersId   uuid.UUID
	CoursesId uuid.UUID
	Progress  uint8
	CreatedAt int64
	UpdatedAt int64
}

func NewEnrollment(usersId, coursesId uuid.UUID) Enrollment {
	return Enrollment{
		ID:        uuid.New(),
		UsersId:   usersId,
		CoursesId: coursesId,
		Progress:  0,
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: time.Now().UnixMilli(),
	}
}

func (e *Enrollment) UpdateProgress(progress uint8) (err error) {
	if e.Progress < progress {
		e.Progress = progress
		e.UpdatedAt = time.Now().UnixMilli()

		return
	}

	return errs.ErrInvalidProgress

}

func (e Enrollment) IsEnrolled() bool {
	return e.ID != uuid.Nil
}
