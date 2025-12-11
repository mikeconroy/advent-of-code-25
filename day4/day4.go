package day4

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day4/input")
	return part1(input), part2(input)
}

type Grid struct {
	grid      map[key]rune
	colsCount int
	rowsCount int
}

type key struct {
	x, y int
}

func loadGrid(input []string) Grid {
	y := 0
	cols := 0
	grid := make(map[key]rune)
	for _, line := range input {
		x := 0
		for _, char := range line {
			currKey := key{x, y}
			grid[currKey] = char
			x++
			cols = x
		}
		y++
	}
	return Grid{grid, y, cols}
}

func (g Grid) print() {
	for y := 0; y < g.rowsCount; y++ {
		for x := 0; x < g.colsCount; x++ {
			fmt.Print(string(g.grid[key{x, y}]))
		}
		fmt.Println()
	}
}

func (g Grid) countSurroundingsAt(loc key) int {
	count := 0
	for yOffset := -1; yOffset < 2; yOffset++ {
		for xOffset := -1; xOffset < 2; xOffset++ {
			if yOffset == 0 && xOffset == 0 {
				continue
			}
			newX := loc.x + xOffset
			newY := loc.y + yOffset
			if newX < 0 || newY < 0 {
				continue
			}
			if newX >= g.colsCount || newY >= g.rowsCount {
				continue
			}
			if g.grid[key{newX, newY}] == '@' {

				count++
			}
		}
	}
	return count
}

func part1(input []string) string {
	grid := loadGrid(input)
	accessCount := 0
	for y := 0; y < grid.rowsCount; y++ {
		for x := 0; x < grid.colsCount; x++ {
			if grid.grid[key{x, y}] == '@' {
				if grid.countSurroundingsAt(key{x, y}) < 4 {
					accessCount++
				}
			}
		}
	}
	return fmt.Sprint(accessCount)
}

func part2(input []string) string {
	return fmt.Sprint(0)
}
