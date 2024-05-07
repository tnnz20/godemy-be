package entities

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

const (
	validEmail    = "jhon@gmail.com"
	ValidPassword = "jhonpassword"
	ValidName     = "Jhon"
	ValidRole     = "student"
)

func TestUsersEntities(t *testing.T) {
	t.Run("Success validate users", func(t *testing.T) {
		User := Users{
			Email:    validEmail,
			Password: ValidPassword,
			Name:     ValidName,
		}

		err := User.ValidateAuth()
		require.Nil(t, err)

	})

	t.Run("Failed validate users, email must required", func(t *testing.T) {
		Users := Users{
			Email:    "",
			Password: ValidPassword,
			Name:     ValidName,
		}

		err := Users.ValidateAuth()

		require.NotNil(t, err)
		require.Equal(t, errs.ErrEmailRequired, err)
	})

	t.Run("Failed validate users, invalid email", func(t *testing.T) {
		Users := Users{
			Email:    "Jhon",
			Password: ValidPassword,
			Name:     ValidName,
		}

		err := Users.ValidateAuth()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidEmail, err)
	})

	t.Run("Failed validate users, password must required", func(t *testing.T) {
		Users := Users{
			Email:    validEmail,
			Password: "",
			Name:     ValidName,
		}

		err := Users.ValidateAuth()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrPasswordRequired, err)
	})

	t.Run("Failed validate users, password length must be at least 8 characters", func(t *testing.T) {
		Users := Users{
			Email:    validEmail,
			Password: "jhon",
			Name:     ValidName,
		}

		err := Users.ValidateAuth()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidLengthPassword, err)
	})

	t.Run("Failed validate users, name must required", func(t *testing.T) {
		Users := Users{
			Email:    validEmail,
			Password: ValidPassword,
			Name:     "",
		}

		err := Users.ValidateAuth()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrNameRequired, err)
	})

	t.Run("Failed validate users, name length must be at least 3 characters", func(t *testing.T) {
		Users := Users{
			Email:    validEmail,
			Password: ValidPassword,
			Name:     "Jh",
		}

		err := Users.ValidateAuth()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidLengthName, err)

	})
}

func TestRolesEntities(t *testing.T) {
	t.Run("Success validate roles", func(t *testing.T) {
		Role := Roles{
			Role: ValidRole,
		}

		err := Role.ValidateRole()
		require.Nil(t, err)
	})

	t.Run("Failed validate roles, role must required", func(t *testing.T) {
		Role := Roles{
			Role: "",
		}

		err := Role.ValidateRole()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrRoleRequired, err)
	})

	t.Run("Failed validate roles, role must be student or teacher", func(t *testing.T) {
		Role := Roles{
			Role: "admin",
		}

		err := Role.ValidateRole()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidRole, err)
	})
}
