package auth

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type Auth struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a Auth) Validate() (err error) {
	if err := a.ValidateEmail(); err != nil {
		return err
	}

	if err := a.ValidatePassword(); err != nil {
		return err
	}

	if err := a.ValidateRole(); err != nil {
		return err
	}

	return
}

func (a Auth) ValidateEmail() (err error) {
	if a.Email == "" {
		return errs.ErrEmailRequired
	}

	splitEmail := strings.Split(a.Email, "@")
	if len(splitEmail) != 2 {
		return errs.ErrInvalidEmail
	}

	return
}

func (a Auth) ValidatePassword() (err error) {
	if a.Password == "" {
		return errs.ErrPasswordRequired
	}

	if len(a.Password) < 8 {
		return errs.ErrInvalidLengthPassword
	}

	return
}

func (a Auth) ValidateRole() (err error) {
	if a.Role == "" {
		return errs.ErrRoleRequired
	}

	if a.Role != "student" && a.Role != "teacher" {
		return errs.ErrInvalidRole
	}

	return
}
