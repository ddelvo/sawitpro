package input

import (
	"bufio"
	"fmt"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"
	"io"
	"os"
	"strconv"
	"strings"
)

func GetInput(stdin io.Reader) []string {
	scanner := bufio.NewScanner(stdin)
	var landSpecsInput []string

	for scanner.Scan() {
		line := scanner.Text()
		landSpecsInput = append(landSpecsInput, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return landSpecsInput
}

func GetWidthLengthCount(spec string) ([]int, error) {

	substrings := strings.Fields(spec)
	res := make([]int, len(substrings))

	for i, s := range substrings {
		num, err := strconv.Atoi(s)
		if err != nil {
			return res, err
		}
		res[i] = num
	}

	return res, nil
}

func getXYHeight(spec string) ([]int, error) {

	substrings := strings.Fields(spec)
	res := make([]int, len(substrings))

	for i, s := range substrings {
		num, err := strconv.Atoi(s)
		if err != nil {
			return res, err
		}
		res[i] = num
	}

	return res, nil
}

func GetTreeSpecs(number int, landSpecs []string) ([][]int, error) {
	var res [][]int
	for i := 1; i <= number; i++ {
		treeSpecsInput := landSpecs[i]

		specs, err := getXYHeight(treeSpecsInput)
		if err != nil || len(specs) != 3 {
			return res, apperror.ErrBadInput
		}
		res = append(res, specs)
	}
	return res, nil
}

func Exit() {
	fmt.Fprintln(os.Stderr, "FAIL")
	os.Exit(1)
}
