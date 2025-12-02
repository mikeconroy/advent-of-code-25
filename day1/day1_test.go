package day1

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func TestDay1Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "3" {
		t.Fatal("Day 1 - Part 1 output should be 3")
	}
}

func TestDay1Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "6" {
		t.Fatal("Day 1 - Part 2 output should be 6")
	}
}
