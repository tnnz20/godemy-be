package entities

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

const (
	validEmail       = "jhon@gmail.com"
	ValidatePassword = "jhonpassword"
	ValidateRole     = "student"
)

func TestUserEntity(t *testing.T) {
	t.Run("Success validate", func(t *testing.T) {
		user := User{
			Email:    validEmail,
			Password: ValidatePassword,
			Role:     ValidateRole,
		}

		err := user.Validate()
		require.Nil(t, err)
	})

	t.Run("Failed validate email", func(t *testing.T) {
		user := User{
			Email:    "Jhon",
			Password: ValidatePassword,
			Role:     ValidateRole,
		}

		err := user.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidEmail, err)
	})

	t.Run("Failed validate email must required", func(t *testing.T) {
		user := User{
			Email:    "",
			Password: ValidatePassword,
			Role:     ValidateRole,
		}

		err := user.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrEmailRequired, err)
	})

	t.Run("Failed validate password", func(t *testing.T) {
		user := User{
			Email:    validEmail,
			Password: "jon",
			Role:     ValidateRole,
		}

		err := user.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidLengthPassword, err)
	})

	t.Run("Failed validate password must required", func(t *testing.T) {
		user := User{
			Email:    validEmail,
			Password: "",
			Role:     ValidateRole,
		}

		err := user.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrPasswordRequired, err)
	})

	t.Run("Failed validate password must have 8 characters", func(t *testing.T) {
		user := User{
			Email:    validEmail,
			Password: "jhonpas",
			Role:     ValidateRole,
		}

		err := user.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidLengthPassword, err)
	})

	t.Run("Failed validate role", func(t *testing.T) {
		user := User{
			Email:    validEmail,
			Password: "jhonpassword",
			Role:     "admin",
		}

		err := user.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidRole, err)
	})

	t.Run("Failed validate role must required", func(t *testing.T) {
		user := User{
			Email:    validEmail,
			Password: "jhonpassword",
			Role:     "",
		}

		err := user.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrRoleRequired, err)
	})

}
