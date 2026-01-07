package day9

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func TestDay9Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expected := "50"
	result := part1(input)
	if result != expected {
		t.Fatal("Day 9 - Part 1 output should be", expected, "but got", result)
	}
}

func TestDay9Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expected := "0"
	result := part2(input)
	if result != expected {
		t.Fatal("Day 9 - Part 2 output should be", expected, "but got", result)
	}
}
