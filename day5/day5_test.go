package day5

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func TestDay5Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expected := "3"
	result := part1(input)
	if result != expected {
		t.Fatal("Day 5 - Part 1 output should be", expected, "but got", result)
	}
}

func TestDay5Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expected := "0"
	result := part2(input)
	if result != expected {
		t.Fatal("Day 5 - Part 2 output should be", expected, "but got", result)
	}
}
