package main

import (
	"fmt"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/input"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/usecase"
	"os"
)

func main() {
	landSpecsInput, err := input.GetInput(os.Stdin)
	if err != nil {
		input.Exit()
	}

	landSpecs, err := input.GetWidthLengthCount(landSpecsInput[0])
	if err != nil || len(landSpecs) != 3 {
		input.Exit()
	}

	width := landSpecs[0]
	length := landSpecs[1]
	treesCount := landSpecs[2]

	plantation := usecase.NewPlantation(length, width)
	specs, err := input.GetTreeSpecs(treesCount, landSpecsInput)
	if err != nil {
		input.Exit()
	}

	err = plantation.PopulateLand(specs)
	if err != nil {
		input.Exit()
	}

	res := plantation.CalculateTravel()
	fmt.Printf("%d\n", res)
}
