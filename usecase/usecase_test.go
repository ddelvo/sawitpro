package usecase_test

import (
	"github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlantation(t *testing.T) {

	t.Run("when plantation is created is should return plantation", func(t *testing.T) {
		length := 5
		width := 1

		pln := usecase.NewPlantation(length, width)

		assert.NotNil(t, pln)
	})

	t.Run("when plantation is populated, it should not error", func(t *testing.T) {
		length := 5
		width := 1

		var trees [][]int
		trees = append(trees, []int{1, 1, 3})

		pln := usecase.NewPlantation(length, width)
		err := pln.PopulateLand(trees)
		assert.Nil(t, err)
	})

	t.Run("when plantation is populated with wrong input, it should return error", func(t *testing.T) {
		length := 5
		width := 1

		var trees [][]int
		trees = append(trees, []int{2, 2, 3})

		pln := usecase.NewPlantation(length, width)
		err := pln.PopulateLand(trees)
		assert.ErrorIs(t, err, apperror.ErrBadInput)
	})

	t.Run("when drone travel distance is calculated, it should return the right amount", func(t *testing.T) {
		length := 2
		width := 6

		var trees [][]int
		trees = append(trees, []int{3, 1, 5})
		trees = append(trees, []int{3, 2, 10})

		pln := usecase.NewPlantation(length, width)
		_ = pln.PopulateLand(trees)
		res := pln.CalculateTravel()
		assert.Equal(t, 132, res)
	})
}
