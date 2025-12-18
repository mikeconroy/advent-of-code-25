package day7

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func TestDay7Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expected := "21"
	result := part1(input)
	if result != expected {
		t.Fatal("Day 7 - Part 1 output should be", expected, "but got", result)
	}
}

func TestDay7Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expected := "40"
	result := part2(input)
	if result != expected {
		t.Fatal("Day 7 - Part 2 output should be", expected, "but got", result)
	}
}
