package day7

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day7/input")
	return part1(input), part2(input)
}

func loadGrid(input []string) [][]rune {
	var grid [][]rune
	for _, line := range input {
		if line != "" {
			row := make([]rune, len(line))
			for idx, char := range line {
				row[idx] = char
			}
			grid = append(grid, row)
		}
	}
	return grid
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
}

func part1(input []string) string {
	grid := loadGrid(input)
	splitCount := 0
	for y, row := range grid {
		for x, char := range row {
			if y != 0 {
				if grid[y-1][x] == '|' || grid[y-1][x] == 'S' {
					if char == '.' {
						grid[y][x] = '|'
					} else if char == '^' {
						splitCount += 1
						grid[y][x-1] = '|'
						grid[y][x+1] = '|'
					}
				}
			}
		}
	}
	return fmt.Sprint(splitCount)
}

type Key struct {
	x, y int
}

func traverse(x, y int, grid [][]rune, cache map[Key]int) int {
	key := Key{x, y}
	if res, ok := cache[key]; ok {
		return res
	}

	if y == len(grid) {
		return 1
	}

	if grid[y][x] == '^' {
		res := traverse(x-1, y+1, grid, cache) + traverse(x+1, y+1, grid, cache)
		cache[key] = res
		return res
	} else {
		res := traverse(x, y+1, grid, cache)
		cache[key] = res
		return res
	}
}

func part2(input []string) string {
	grid := loadGrid(input)
	startX, startY := 0, 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'S' {
				startX = x
				startY = y
				break
			}
		}
	}
	res := traverse(startX, startY, grid, make(map[Key]int))
	return fmt.Sprint(res)
}
