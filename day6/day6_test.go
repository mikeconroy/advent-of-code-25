package day6

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func TestDay6Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expected := "4277556"
	result := part1(input)
	if result != expected {
		t.Fatal("Day 6 - Part 1 output should be", expected, "but got", result)
	}
}

func TestDay6Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	input = input[:len(input)-1]
	expected := "3263827"
	result := part2(input)
	if result != expected {
		t.Fatal("Day 6 - Part 2 output should be", expected, "but got", result)
	}
}
