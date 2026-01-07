package day9

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day9/input")
	return part1(input), part2(input)
}

type Point struct {
	x, y int
}

func parseInput(input []string) []Point {
	points := make([]Point, 0)
	for _, line := range input {
		if line != "" {
			split := strings.Split(line, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			point := Point{x, y}
			points = append(points, point)
		}
	}
	return points
}

func (p Point) getArea(p2 Point) int {
	dx := math.Abs(float64(p.x-p2.x)) + 1
	dy := math.Abs(float64(p.y-p2.y)) + 1
	return int(dx * dy)
}

func part1(input []string) string {
	points := parseInput(input)
	greatestArea := 0
	for i := 0; i < len(points)-1; i++ {
		p1 := points[i]
		for j := i; j < len(points); j++ {
			p2 := points[j]
			area := p1.getArea(p2)
			if area > greatestArea {
				greatestArea = area
			}
		}
	}
	return fmt.Sprint(greatestArea)
}

func part2(input []string) string {
	return fmt.Sprint(0)
}
