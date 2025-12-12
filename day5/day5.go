package day5

import (
	"cmp"
	"fmt"
	"github.com/mikeconroy/advent-of-code-25/utils"
	"slices"
	"strconv"
	"strings"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day5/input")
	return part1(input), part2(input)
}

type Range struct {
	min, max int
}

type Database struct {
	ranges      []Range
	ingredients []int
}

func loadInput(input []string) Database {
	ranges := make([]Range, 0)
	ingredients := make([]int, 0)
	for _, line := range input {
		if strings.Contains(line, "-") {
			split := strings.Split(line, "-")
			min, _ := strconv.Atoi(split[0])
			max, _ := strconv.Atoi(split[1])
			ranges = append(ranges, Range{min, max})
		} else if line != "" {
			id, _ := strconv.Atoi(line)
			ingredients = append(ingredients, id)
		}
	}
	return Database{ranges, ingredients}
}

func (db Database) isFresh(ing int) bool {
	for _, freshRange := range db.ranges {
		if ing >= freshRange.min && ing <= freshRange.max {
			return true
		}
	}
	return false
}

func part1(input []string) string {
	db := loadInput(input)
	freshCount := 0
	for _, ing := range db.ingredients {
		if db.isFresh(ing) {
			freshCount++
		}
	}
	return fmt.Sprint(freshCount)
}

func cmpRange(a, b Range) int {
	return cmp.Compare(a.min, b.min)
}

// 312407320826103 - too low
// 403435807163373 - too high
// 348548952146327 - Not correct
// 345761047266506 - Not Correct
func part2(input []string) string {
	db := loadInput(input)
	slices.SortFunc(db.ranges, cmpRange)
	currRange := db.ranges[0]
	result := 0
	for i := 1; i < len(db.ranges); i++ {
		newRange := db.ranges[i]
		if newRange.min <= currRange.max {
			if newRange.max > currRange.max {
				currRange.max = newRange.max
			}
		} else {
			result += (currRange.max - currRange.min + 1)
			currRange = newRange
		}
	}
	result += (currRange.max - currRange.min + 1)
	return fmt.Sprint(result)
}
