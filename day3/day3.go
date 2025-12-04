package day3

import (
	"fmt"
	"strconv"

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

func getMaxJolts(bank string, joltsRequired int, level int, cache map[string]int) int {
	key := bank + "-" + strconv.Itoa(joltsRequired)
	if cache[key] != 0 {
		return cache[key]
	}
	// fmt.Println(strings.Repeat(" ", level), "Bank:", bank, "JoltsRequired:", joltsRequired)
	// Base Case - if no more need to be activated we return the current bank
	// If the amount to be activated is the length of the remaining batteries we return that value.
	if len(bank) <= joltsRequired {
		result, _ := strconv.Atoi(bank)
		return result
	} else {
		// Compare Max Jolts if we take the first value vs. continuing
		var firstDigitResult int
		if joltsRequired-1 == 0 {
			firstDigitResult = int(bank[0]) - '0'
		} else {
			firstDigitResult, _ = strconv.Atoi(string(bank[0]) + strconv.Itoa(getMaxJolts(bank[1:len(bank)], joltsRequired-1, level+1, cache)))
		}
		withoutFirstDigit := getMaxJolts(bank[1:len(bank)], joltsRequired, level+1, cache)
		// fmt.Println(strings.Repeat(" ", level), "Bank:", bank, "JoltsReqs:", joltsRequired, "First:", firstDigitResult, "Without:", withoutFirstDigit, bank[1:len(bank)])
		if firstDigitResult > withoutFirstDigit {
			cache[key] = firstDigitResult
			return firstDigitResult
		} else {
			cache[key] = withoutFirstDigit
			return withoutFirstDigit
		}
	}
}

func part1(input []string) string {
	totalJoltage := 0
	for _, bank := range input {
		result := getMaxJolts(bank, 2, 0, make(map[string]int))
		totalJoltage += result
	}
	return fmt.Sprint(totalJoltage)
}

func part2(input []string) string {
	memoisedJolts := make(map[string]int)
	totalJoltage := 0
	for _, bank := range input {
		totalJoltage += getMaxJolts(bank, 12, 0, memoisedJolts)
	}
	return fmt.Sprint(totalJoltage)
}
