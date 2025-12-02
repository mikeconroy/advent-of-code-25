package day1

import (
	"fmt"
	"github.com/mikeconroy/advent-of-code-25/utils"
	"strconv"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day1/input")
	return part1(input), part2(input)
}

func processInstruction(currPos int, instruction string) (int, int) {
	dir := instruction[0]
	distance, _ := strconv.Atoi(instruction[1:len(instruction)])
	newPos := currPos
	zeroHitCount := distance / 100
	distance %= 100
	if dir == 'L' {
		newPos -= distance
		if newPos == 0 {
			zeroHitCount += 1
		} else if newPos < 0 {
			if currPos != 0 {
				zeroHitCount += 1
			}
			newPos += 100
		}
	} else if dir == 'R' {
		newPos += distance
		if newPos > 99 {
			zeroHitCount += 1
			newPos -= 100
		}
	}

	// fmt.Println("Curr:", currPos, "Instruction:", instruction, "Dist:", distance, "NewPos:", newPos, "Hits:", zeroHitCount)
	return newPos, zeroHitCount
}

func part1(input []string) string {
	currPos := 50
	countAtZero := 0
	for _, instruction := range input {
		if instruction != "" {
			currPos, _ = processInstruction(currPos, instruction)
			if currPos == 0 {
				countAtZero += 1
			}
		}
	}
	return fmt.Sprint(countAtZero)
}

func part2(input []string) string {
	currPos := 50
	zeroCount := 0
	for _, instruction := range input {
		if instruction != "" {
			var zeroHitCount int
			currPos, zeroHitCount = processInstruction(currPos, instruction)
			zeroCount += zeroHitCount
		}
	}
	return fmt.Sprint(zeroCount)
}
