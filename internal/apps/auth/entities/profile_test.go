package entities

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

func TestProfileEntity(t *testing.T) {
	t.Run("Success validate profile", func(t *testing.T) {
		profile := Profile{
			Name:   "Jhon",
			UserID: uuid.New(),
		}

		err := profile.Validate()
		require.Nil(t, err)
	})

	t.Run("Failed validate name must required", func(t *testing.T) {
		profile := Profile{
			Name:   "",
			UserID: uuid.New(),
		}

		err := profile.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrNameRequired, err)
	})

	t.Run("Failed validate name must have 3 characters", func(t *testing.T) {
		profile := Profile{
			Name:   "Jh",
			UserID: uuid.New(),
		}

		err := profile.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidLengthName, err)
	})

	t.Run("Failed validate user id must required", func(t *testing.T) {
		profile := Profile{
			Name:   "Jhon",
			UserID: uuid.Nil,
		}

		err := profile.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrUserIDRequired, err)
	})
}
