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

func isValidId2(id int) bool {
	idStr := strconv.Itoa(id)
	idLen := len(idStr)

	//123,123,123
	//Take 1 and check if all digits are 1
	//Take 12 and check if remaining digits are 12...
	//Take 123 and check if remaining digits are 123...
	for count := 1; count <= idLen/2; count++ {
		// fmt.Println(idStr, count, idStr[0:count])
		seq := idStr[0:count]
		if idLen%count == 0 {
			newId := strings.Repeat(seq, idLen/count)
			if idStr == newId {
				return true
			}
		}
	}
	return false
}

func part2(input []string) string {
	ranges := getRanges(input)
	var sumInvalids uint64
	sumInvalids = 0
	for _, idRange := range ranges {
		for i := idRange.floor; i <= idRange.ceiling; i++ {
			if isValidId2(i) {
				sumInvalids += uint64(i)
			}
		}
	}
	return fmt.Sprint(sumInvalids)
}
