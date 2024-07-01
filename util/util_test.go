package util_test

import (
	"github.com/SawitProRecruitment/JuniorBackendEngineering/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtil(t *testing.T) {
	t.Run("abs is called, it should return the absolute value of negative", func(t *testing.T) {
		x := -10

		xAbs := util.Abs(x)

		assert.Equal(t, 10, xAbs)
	})

	t.Run("abs is called, it should return the absolute value of positive", func(t *testing.T) {
		x := 10

		xAbs := util.Abs(x)

		assert.Equal(t, 10, xAbs)
	})
}
