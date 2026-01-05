package day8

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func TestDay8Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expected := "40"
	result := part1(input, 10)
	if result != expected {
		t.Fatal("Day 8 - Part 1 output should be", expected, "but got", result)
	}
}

func TestDay8Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expected := "25272"
	result := part2(input)
	if result != expected {
		t.Fatal("Day 8 - Part 2 output should be", expected, "but got", result)
	}
}
