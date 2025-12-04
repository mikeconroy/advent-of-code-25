package day3

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day3/input")
	return part1(input), part2(input)
}

func getMaxJoltage(bank string) int {
	firstJolt := 0
	firstJoltIndex := 0
	for i := 0; i < len(bank)-1; i++ {
		currJolt := int(bank[i]) - '0'
		if currJolt > firstJolt {
			firstJolt = currJolt
			firstJoltIndex = i
		}
	}

	secondJolt := 0
	for j := firstJoltIndex + 1; j < len(bank); j++ {
		currJolt := int(bank[j]) - '0'
		if currJolt > secondJolt {
			secondJolt = currJolt
		}
	}
	maxJolt := (firstJolt * 10) + secondJolt
	return maxJolt
}

func part1(input []string) string {
	totalJoltage := 0
	for _, bank := range input {
		totalJoltage += getMaxJoltage(bank)
	}
	return fmt.Sprint(totalJoltage)
}

func part2(input []string) string {
	return fmt.Sprint(0)
}
