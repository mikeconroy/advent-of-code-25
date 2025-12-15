package day6

import (
	"fmt"
	"github.com/mikeconroy/advent-of-code-25/utils"
	"strconv"
	"strings"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day6/input")
	input = input[:len(input)-1]
	return part1(input), part2(input)
}

type Operation int

const (
	MULTIPLY = iota
	DIVIDE
	ADD
	SUBTRACT
)

var operationMap = map[string]Operation{
	"*": MULTIPLY,
	"/": DIVIDE,
	"+": ADD,
	"-": SUBTRACT,
}

type Problem struct {
	nums []int
	op   Operation
}

func parseInput(input []string) []Problem {
	probs := make([]Problem, 0)
	for lineIdx, line := range input {
		splits := strings.Fields(line)
		for valIdx := 0; valIdx < len(splits); valIdx++ {
			var prob Problem
			op, ok := operationMap[splits[valIdx]]
			if ok {
				probs[valIdx].op = op
				continue
			}
			val, _ := strconv.Atoi(splits[valIdx])
			if lineIdx == 0 {
				prob = Problem{nums: make([]int, 1)}
				prob.nums[0] = val
				probs = append(probs, prob)
			} else {
				prob = probs[valIdx]
				prob.nums = append(prob.nums, val)
				probs[valIdx] = prob
			}
		}
	}
	return probs
}

func solveProblems(probs []Problem) int {
	total := 0
	for _, problem := range probs {
		res := problem.nums[0]
		for _, val := range problem.nums[1:] {
			if problem.op == MULTIPLY {
				res *= val
			} else if problem.op == ADD {
				res += val
			}
		}
		total += res
	}
	return total
}

func part1(input []string) string {
	problems := parseInput(input)
	res := solveProblems(problems)
	return fmt.Sprint(res)
}

func part2(input []string) string {
	opSplit := strings.Fields(input[len(input)-1])
	probs := make([]Problem, len(opSplit))
	for idx, opStr := range opSplit {
		probs[idx].op, _ = operationMap[opStr]
	}

	cols := make([]string, len(input[0]))

	for _, line := range input[:len(input)-1] {
		for idx := 0; idx < len(line); idx++ {
			cols[idx] += string(line[idx])
		}
	}

	currProbIdx := 0
	for _, col := range cols {
		col = strings.Trim(col, " ")
		if col == "" {
			currProbIdx++
		} else {
			val, _ := strconv.Atoi(col)
			probs[currProbIdx].nums = append(probs[currProbIdx].nums, val)
		}
	}
	res := solveProblems(probs)
	return fmt.Sprint(res)
}
