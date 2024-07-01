package input_test

import (
	"fmt"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/input"
	"github.com/stretchr/testify/assert"
	"io"
	"strconv"
	"strings"
	"testing"
)

type errorReader struct{}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("read error")
}

func TestInput(t *testing.T) {

	t.Run("when GetInput is called, it should be able to receive input", func(t *testing.T) {
		tests := []struct {
			input    string
			expected []string
		}{
			{
				input:    "line1\nline2\nline3\n",
				expected: []string{"line1", "line2", "line3"},
			},
		}

		r := strings.NewReader(tests[0].input)
		result, err := input.GetInput(r)

		assert.Nil(t, err)
		assert.Equal(t, tests[0].expected, result)
	})

	t.Run("when GetInput is called and there is an error, it should return error", func(t *testing.T) {
		tests := []struct {
			input    io.Reader
			expected []string
		}{
			{
				input:    &errorReader{},
				expected: []string{},
			},
		}

		_, err := input.GetInput(tests[0].input)
		assert.ErrorIs(t, err, apperror.ErrInputScan)
	})

	t.Run("when GetWidthLengthCount is called, it should return width, length and count", func(t *testing.T) {
		landSpecs := "5 1 3"

		specs, err := input.GetWidthLengthCount(landSpecs)

		assert.Nil(t, err)
		assert.Equal(t, 5, specs[0])
		assert.Equal(t, 1, specs[1])
		assert.Equal(t, 3, specs[2])
	})

	t.Run("when GetWidthLengthCount is called but there is non-numeric string, it should return error", func(t *testing.T) {
		landSpecs := "5 1 a"
		_, err := input.GetWidthLengthCount(landSpecs)

		assert.ErrorIs(t, err, strconv.ErrSyntax)
	})

	t.Run("when GetTreeSpecs is called, it should return the trees spec", func(t *testing.T) {
		number := 3
		landSpecs := []string{"5 1 3", "2 1 5", "3 1 3", "4 1 4"}

		res, err := input.GetTreeSpecs(number, landSpecs)

		assert.Nil(t, err)
		assert.Equal(t, res[0], []int{2, 1, 5})
		assert.Equal(t, res[1], []int{3, 1, 3})
		assert.Equal(t, res[2], []int{4, 1, 4})
	})

	t.Run("when GetTreeSpecs is called but there is non-numeric string, it should return error", func(t *testing.T) {
		number := 3
		landSpecs := []string{"5 1 3", "2 1 5", "3 1 a", "4 1 4"}

		_, err := input.GetTreeSpecs(number, landSpecs)

		assert.ErrorIs(t, err, apperror.ErrBadInput)
	})

}
