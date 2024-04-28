package errs

import "errors"

var (
	ErrInvalidLengthUUID = errors.New("uuid must be 16 characters")
	ErrUserIDRequired    = errors.New("user id required")

	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")

	ErrInvalidEmail  = errors.New("invalid email")
	ErrEmailRequired = errors.New("email required")
	ErrEmailNotFound = errors.New("email not found")

	ErrInvalidLengthName = errors.New("name must be at least 3 characters")
	ErrNameRequired      = errors.New("name required")

	ErrPasswordRequired      = errors.New("password required")
	ErrInvalidLengthPassword = errors.New("password must be at least 8 characters")
	ErrWrongPassword         = errors.New("wrong password")

	ErrInvalidRole  = errors.New("invalid role")
	ErrRoleRequired = errors.New("role required")
)

var (
	ErrorMapping = map[error]uint32{
		ErrInvalidLengthUUID: 400,
		ErrUserIDRequired:    400,

		ErrUserNotFound:       404,
		ErrEmailAlreadyExists: 409,

		ErrInvalidEmail:  400,
		ErrEmailRequired: 400,
		ErrEmailNotFound: 404,

		ErrInvalidLengthName: 400,
		ErrNameRequired:      400,

		ErrPasswordRequired:      400,
		ErrInvalidLengthPassword: 400,
		ErrWrongPassword:         400,

		ErrInvalidRole:  400,
		ErrRoleRequired: 400,
	}
)
