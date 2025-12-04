package day2

import (
	"strings"
	"testing"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func TestDay2Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	input = strings.Split(input[0], ",")
	if part1(input) != "1227775554" {
		t.Fatal("Day 2 - Part 1 output should be 1227775554")
	}
}

func TestDay2Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	input = strings.Split(input[0], ",")
	result := part2(input)
	if result != "4174379265" {
		t.Fatal("Day 2 - Part 2 output should be 4174379265 but got", result)
	}
}
