package auth

import (
	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type Profile struct {
	ID   uuid.UUID
	Name string
}

func (p Profile) Validate() (err error) {
	if err := p.ValidateName(); err != nil {
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
