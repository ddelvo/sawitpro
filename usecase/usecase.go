package usecase

import (
	"github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/util"
)

const (
	distanceBetweenTrees = 10
	droneExtraHeight     = 2
)

type TravelInfo struct {
	totalHorizontal int
	totalVertical   int
	currentHeight   int
}

type Plantation struct {
	land [][]int
}

func NewPlantation(length, width int) *Plantation {
	land := make([][]int, length)
	for i := range land {
		land[i] = make([]int, width)
	}
	return &Plantation{land: land}
}

func (p *Plantation) PopulateLand(trees [][]int) error {
	for _, tree := range trees {
		x, y, height := tree[0], tree[1], tree[2]
		if x > len(p.land[0]) || y > len(p.land) {
			return apperror.ErrBadInput
		}
		p.land[y-1][x-1] = height
	}
	return nil
}

func (p *Plantation) CalculateTravel() int {
	info := &TravelInfo{
		totalHorizontal: -distanceBetweenTrees,
		totalVertical:   droneExtraHeight,
		currentHeight:   0,
	}

	for i, row := range p.land {
		if (i+1)%2 == 0 {
			p.traverseRowReverse(row, info)
		} else {
			p.traverseRow(row, info)
		}
	}

	info.totalVertical += util.Abs(info.currentHeight)
	return info.totalVertical + info.totalHorizontal
}

func (p *Plantation) traverseRow(row []int, info *TravelInfo) {
	for _, height := range row {
		p.processTree(height, info)
	}
}

func (p *Plantation) traverseRowReverse(row []int, info *TravelInfo) {
	for j := len(row) - 1; j >= 0; j-- {
		p.processTree(row[j], info)
	}
}

func (p *Plantation) processTree(height int, info *TravelInfo) {
	info.totalHorizontal += distanceBetweenTrees
	if height != 0 {
		heightChanges := height - info.currentHeight
		info.totalVertical += util.Abs(heightChanges)
		info.currentHeight = height
	}
}
