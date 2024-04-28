package entities

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

func TestProfileEntity(t *testing.T) {
	t.Run("Success validate profile", func(t *testing.T) {
		profile := Profile{
			Name: "Jhon",
		}

		err := profile.Validate()
		require.Nil(t, err)
	})

	t.Run("Failed validate name must required", func(t *testing.T) {
		profile := Profile{
			Name: "",
		}

		err := profile.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrNameRequired, err)
	})

	t.Run("Failed validate name must have 3 characters", func(t *testing.T) {
		profile := Profile{
			Name: "Jh",
		}

		err := profile.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidLengthName, err)
	})
}
