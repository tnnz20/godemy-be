package entities

import (
	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type Roles struct {
	UsersId uuid.UUID
	Role    string
}

func (r Roles) ValidateRole() (err error) {
	if r.Role == "" {
		return errs.ErrRoleRequired
	}

	validRole := map[string]bool{
		"student": true,
		"teacher": true,
	}

	if _, ok := validRole[r.Role]; !ok {
		return errs.ErrInvalidRole
	}

	return
}
