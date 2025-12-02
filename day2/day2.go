package day2

import (
	"fmt"
	"github.com/mikeconroy/advent-of-code-25/utils"
	"strconv"
	"strings"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day2/input")
	input = strings.Split(input[0], ",")
	return part1(input), part2(input)
}

type Range struct {
	floor   int
	ceiling int
}

func getRanges(input []string) []Range {
	ranges := []Range{}
	for _, ids := range input {
		idRange := strings.Split(ids, "-")
		floor, _ := strconv.Atoi(idRange[0])
		ceil, _ := strconv.Atoi(idRange[1])
		ranges = append(ranges, Range{floor, ceil})
	}
	return ranges
}

func isValidId(id int) bool {
	idStr := strconv.Itoa(id)
	idLen := len(idStr)
	// All odd length digits are valid
	if idLen%2 != 0 {
		return true
	}
	lhs := idStr[0 : idLen/2]
	rhs := idStr[idLen/2 : idLen]
	if lhs == rhs {
		return false
	} else {
		return true
	}
}

func part1(input []string) string {
	ranges := getRanges(input)
	sumInvalids := 0
	for _, idRange := range ranges {
		for i := idRange.floor; i <= idRange.ceiling; i++ {
			if !isValidId(i) {
				sumInvalids += i
			}
		}
	}
	return fmt.Sprint(sumInvalids)
}

func part2(input []string) string {
	return fmt.Sprint(0)
}
