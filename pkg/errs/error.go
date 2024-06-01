package errs

import "errors"

var (
	// General
	ErrUserIDRequired = errors.New("user id required")
	ErrUnauthorized   = errors.New("unauthorized")

	// Users
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")

	ErrInvalidEmail  = errors.New("invalid email")
	ErrEmailRequired = errors.New("email required")
	ErrEmailNotFound = errors.New("email not found")

	ErrPasswordRequired      = errors.New("password required")
	ErrInvalidLengthPassword = errors.New("password must be at least 8 characters")
	ErrWrongPassword         = errors.New("wrong password")

	ErrInvalidLengthName = errors.New("name must be at least 3 characters")
	ErrNameRequired      = errors.New("name required")

	ErrInvalidLengthAddress = errors.New("address must be at least 5 characters")
	ErrInvalidGender        = errors.New("gender must be male or female")

	ErrInvalidRole  = errors.New("invalid role")
	ErrRoleRequired = errors.New("role required")

	// Courses
	ErrCourseNameRequired      = errors.New("course name required")
	ErrInvalidCourseNameLength = errors.New("course name must be at least 3 characters")
	ErrCourseCodeRequired      = errors.New("course code required")
	ErrInvalidCourseCodeLength = errors.New("course code must be 10 characters")
	ErrCourseNotFound          = errors.New("course not found")
	ErrCourseCodeAlreadyExist  = errors.New("course code already exist")
	ErrCourseEmpty             = errors.New("course still empty")

	// Enrollment
	ErrInvalidProgress          = errors.New("invalid progress must be greater than before")
	ErrCourseEnrollmentNotFound = errors.New("course enrollment not found")
	ErrUserAlreadyEnrolled      = errors.New("user already enrolled course")
	ErrEmptyEnrollment          = errors.New("enrollment still empty")

	// Assessment
	ErrInvalidAssessmentValue = errors.New("invalid assessment value")
	ErrAssessmentCodeRequired = errors.New("assessment code required")
	ErrAssessmentNotFound     = errors.New("assessment not found")

	ErrInvalidAssessmentStatus        = errors.New("invalid assessment status")
	ErrAssessmentStatusAlreadyCreated = errors.New("assessment status already created")
)

var (
	ErrorMapping = map[error]uint32{
		ErrUserIDRequired: 400,

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

		ErrInvalidLengthAddress: 400,
		ErrInvalidGender:        400,

		ErrInvalidCourseNameLength: 400,
		ErrInvalidCourseCodeLength: 400,
		ErrCourseCodeRequired:      400,
		ErrCourseNameRequired:      400,

		ErrCourseEmpty:            404,
		ErrCourseNotFound:         404,
		ErrCourseCodeAlreadyExist: 409,

		ErrInvalidProgress:          400,
		ErrCourseEnrollmentNotFound: 404,
		ErrUserAlreadyEnrolled:      409,
		ErrEmptyEnrollment:          404,

		ErrInvalidAssessmentValue: 400,
		ErrAssessmentCodeRequired: 400,
		ErrAssessmentNotFound:     404,

		ErrInvalidAssessmentStatus:        400,
		ErrAssessmentStatusAlreadyCreated: 409,
	}
)
