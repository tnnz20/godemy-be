package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type Profile struct {
	Id         uuid.UUID
	Name       string
	Date       time.Time
	Address    string
	Gender     string
	ProfileImg string
	UserID     uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (p Profile) Validate() (err error) {
	if err := p.ValidateName(); err != nil {
		return err
	}

	if err := p.ValidateUserID(); err != nil {
		return err
	}
	return
}

func (p Profile) ValidateName() (err error) {
	if p.Name == "" {
		return errs.ErrNameRequired
	}

	if len(p.Name) < 3 {
		return errs.ErrInvalidLengthName
	}

	return
}

func (p Profile) ValidateUserID() (err error) {
	if p.UserID == uuid.Nil {
		return errs.ErrUserIDRequired
	}

	return
}

func (p Profile) ValidateAddress() (err error) {
	if len(p.Address) < 5 {
		return errs.ErrInvalidLengthAddress
	}

	return
}

func (p Profile) ValidateGender() (err error) {
	validGender := map[string]bool{
		"male":   true,
		"female": true,
	}

	if _, ok := validGender[p.Gender]; !ok {
		return errs.ErrInvalidGender
	}

	return
}
