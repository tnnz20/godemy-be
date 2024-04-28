package auth

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type Users struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u Users) Validate() (err error) {
	if err := u.ValidateEmail(); err != nil {
		return err
	}

	if err := u.ValidatePassword(); err != nil {
		return err
	}

	if err := u.ValidateRole(); err != nil {
		return err
	}

	return
}

func (u Users) ValidateEmail() (err error) {
	if u.Email == "" {
		return errs.ErrEmailRequired
	}

	splitEmail := strings.Split(u.Email, "@")
	if len(splitEmail) != 2 {
		return errs.ErrInvalidEmail
	}

	return
}

func (u Users) ValidatePassword() (err error) {
	if u.Password == "" {
		return errs.ErrPasswordRequired
	}

	if len(u.Password) < 8 {
		return errs.ErrInvalidLengthPassword
	}

	return
}

func (u Users) ValidateRole() (err error) {
	if u.Role == "" {
		return errs.ErrRoleRequired
	}

	if u.Role != "student" && u.Role != "teacher" {
		return errs.ErrInvalidRole
	}

	return
}
