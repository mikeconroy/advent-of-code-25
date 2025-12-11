package day4

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func TestDay4Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expected := "13"
	result := part1(input)
	if result != expected {
		t.Fatal("Day 4 - Part 1 output should be", expected, "but got", result)
	}
}

func TestDay4Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expected := "43"
	result := part2(input)
	if result != expected {
		t.Fatal("Day 4 - Part 2 output should be", expected, "but got", result)
	}
}
