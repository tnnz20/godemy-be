package errs

import "errors"

var (
	ErrInvalidEmail  = errors.New("invalid email")
	ErrEmailRequired = errors.New("email required")

	ErrInvalidLengthName = errors.New("name must be at least 3 characters")
	ErrNameRequired      = errors.New("name required")

	ErrPasswordRequired      = errors.New("password required")
	ErrInvalidLengthPassword = errors.New("password must be at least 8 characters")

	ErrInvalidRole  = errors.New("invalid role")
	ErrRoleRequired = errors.New("role required")
)
