package entities

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

func TestAssessmentEntities(t *testing.T) {
	t.Run("Success validate assessment", func(t *testing.T) {
		assessment := AssessmentResult{
			AssessmentValue: 90,
			AssessmentCode:  "chap-4",
		}

		err := assessment.Validate()
		require.Nil(t, err)
	})

	t.Run("Failed validate assessment, invalid assessment value", func(t *testing.T) {
		assessment := AssessmentResult{
			AssessmentValue: 101,
			AssessmentCode:  "chap-4",
		}

		err := assessment.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidAssessmentValue, err)
	})

	t.Run("Failed validate assessment, assessment code must required", func(t *testing.T) {
		assessment := AssessmentResult{
			AssessmentValue: 90,
			AssessmentCode:  "",
		}

		err := assessment.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrAssessmentCodeRequired, err)
	})
}
