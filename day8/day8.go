package day8

import (
	"cmp"
	"fmt"
	"github.com/spakin/disjoint"
	"slices"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day8/input")
	return part1(input, 1000), part2(input)
}

type Point struct {
	x, y, z int
	circuit *disjoint.Element
}

func (p *Point) String() string {
	str := fmt.Sprint("(", p.x, ",", p.y, ",", p.z, ")")
	return str
}
func (p1 *Point) distanceTo(p2 *Point) float64 {
	dx := float64(p1.x - p2.x)
	dy := float64(p1.y - p2.y)
	dz := float64(p1.z - p2.z)
	// No need to SQRT as we don't care about the actual distance. Just relative distance.
	return (dx * dx) + (dy * dy) + (dz * dz)
}

func loadPoints(input []string) []*Point {
	points := make([]*Point, len(input)-1)

	for idx, line := range input {
		if line != "" {
			split := strings.Split(line, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			z, _ := strconv.Atoi(split[2])
			point := Point{x: x, y: y, z: z, circuit: disjoint.NewElement()}
			points[idx] = &point
		}
	}
	return points
}

type Edge struct {
	p1   *Point
	p2   *Point
	dist float64
}

func (e Edge) String() string {
	return fmt.Sprint(e.p1, e.p2, e.dist)
}

func compEdges(e1, e2 Edge) int {
	return cmp.Compare(e1.dist, e2.dist)
}

// Kruskals Algorithm
func part1(input []string, targetConnections int) string {
	points := loadPoints(input)
	edges := make([]Edge, 0)
	// Calculate the length of all possible edges.
	for idx, p1 := range points[:len(points)-1] {
		for _, p2 := range points[idx+1:] {
			edge := Edge{p1, p2, p1.distanceTo(p2)}
			edges = append(edges, edge)
		}
	}
	slices.SortFunc(edges, compEdges)
	for cxns := 0; cxns < targetConnections; cxns++ {
		disjoint.Union(edges[cxns].p1.circuit, edges[cxns].p2.circuit)
	}

	circuits := make(map[*disjoint.Element]int)
	for _, point := range points {
		circuit := point.circuit.Find()
		count, ok := circuits[circuit]
		if !ok {
			circuits[circuit] = 1
		} else {
			circuits[circuit] = count + 1
		}
	}
	countSlice := make([]int, len(circuits))
	for _, count := range circuits {
		countSlice = append(countSlice, count)
	}
	slices.Sort(countSlice)
	slices.Reverse(countSlice)

	return fmt.Sprint(countSlice[0] * countSlice[1] * countSlice[2])
}

func printPoints(points []*Point) {
	for _, point := range points {
		fmt.Println(point)
	}
}

func part2(input []string) string {
	return fmt.Sprint(0)
}
