package day10

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func TestDay10Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expected := "7"
	result := part1(input)
	if result != expected {
		t.Fatal("Day 10 - Part 1 output should be", expected, "but got", result)
	}
}

func TestDay10Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expected := "0"
	result := part2(input)
	if result != expected {
		t.Fatal("Day 10 - Part 2 output should be", expected, "but got", result)
	}
}
