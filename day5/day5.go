package day5

import (
	"fmt"
	"github.com/mikeconroy/advent-of-code-25/utils"
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

func part2(input []string) string {
	return fmt.Sprint(0)
}
