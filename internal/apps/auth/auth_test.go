package auth

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

func TestAuth(t *testing.T) {
	t.Run("Success validate", func(t *testing.T) {
		auth := Auth{
			Email:    validEmail,
			Password: ValidatePassword,
			Role:     ValidateRole,
		}

		err := auth.Validate()
		require.Nil(t, err)
	})

	t.Run("Failed validate email", func(t *testing.T) {
		auth := Auth{
			Email:    "Jhon",
			Password: ValidatePassword,
			Role:     ValidateRole,
		}

		err := auth.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidEmail, err)
	})

	t.Run("Failed validate email must required", func(t *testing.T) {
		auth := Auth{
			Email:    "",
			Password: ValidatePassword,
			Role:     ValidateRole,
		}

		err := auth.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrEmailRequired, err)
	})

	t.Run("Failed validate password", func(t *testing.T) {
		auth := Auth{
			Email:    validEmail,
			Password: "jon",
			Role:     ValidateRole,
		}

		err := auth.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidLengthPassword, err)
	})

	t.Run("Failed validate password must required", func(t *testing.T) {
		auth := Auth{
			Email:    validEmail,
			Password: "",
			Role:     ValidateRole,
		}

		err := auth.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrPasswordRequired, err)
	})

	t.Run("Failed validate role", func(t *testing.T) {
		auth := Auth{
			Email:    validEmail,
			Password: "jhonpassword",
			Role:     "admin",
		}

		err := auth.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidRole, err)
	})

	t.Run("Failed validate role must required", func(t *testing.T) {
		auth := Auth{
			Email:    validEmail,
			Password: "jhonpassword",
			Role:     "",
		}

		err := auth.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrRoleRequired, err)
	})

}
