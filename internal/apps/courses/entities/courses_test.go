package entities

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tnnz20/godemy-be/pkg/errs"
	"github.com/tnnz20/godemy-be/pkg/helpers"
)

var (
	CourseNameTest = "golang-fundamental-test"
)

func TestCoursesEntities(t *testing.T) {
	randomString := helpers.GenerateRandomString(7)
	CourseCodeTest := fmt.Sprintf("go-%s", randomString)

	t.Run("Success validate courses", func(t *testing.T) {
		Course := Courses{
			CourseName: CourseNameTest,
			CourseCode: CourseCodeTest,
		}

		err := Course.Validate()
		require.Nil(t, err)
	})

	t.Run("Failed validate courses, course name must required", func(t *testing.T) {
		Course := Courses{
			CourseName: "",
			CourseCode: CourseCodeTest,
		}

		err := Course.Validate()

		require.NotNil(t, err)
		require.Equal(t, errs.ErrCourseNameRequired, err)
	})

	t.Run("Failed validate courses, invalid course name length", func(t *testing.T) {
		Course := Courses{
			CourseName: "go",
			CourseCode: CourseCodeTest,
		}

		err := Course.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidCourseNameLength, err)
	})

	t.Run("Failed validate courses, course code must required", func(t *testing.T) {
		Course := Courses{
			CourseName: CourseNameTest,
			CourseCode: "",
		}

		err := Course.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrCourseCodeRequired, err)
	})

	t.Run("Failed validate courses, invalid course code length", func(t *testing.T) {
		Course := Courses{
			CourseName: CourseNameTest,
			CourseCode: "go",
		}

		err := Course.Validate()
		require.NotNil(t, err)
		require.Equal(t, errs.ErrInvalidCourseCodeLength, err)
	})
}
