package day9

import (
	"cmp"
	"fmt"
	"math"
	"slices"
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

func createRgTiles(points []Point) map[Point]bool {
	tiles := make(map[Point]bool)
	prevPoint := points[len(points)-1]
	for _, point := range points {
		newPoint := Point{prevPoint.x, prevPoint.y}
		tiles[point] = true
		for newPoint.x != point.x || newPoint.y != point.y {
			// if tiles[Point{newPoint.x, newPoint.y}] != 'R' {
			tiles[Point{newPoint.x, newPoint.y}] = true
			// }
			if prevPoint.y > point.y {
				newPoint.y--
			} else if prevPoint.y < point.y {
				newPoint.y++
			}

			if prevPoint.x > point.x {
				newPoint.x--
			} else if prevPoint.x < point.x {
				newPoint.x++
			}
		}
		prevPoint = point
	}
	return tiles
}

func printTiles(tiles map[Point]bool) {
	maxX := 0
	maxY := 0
	for point, _ := range tiles {
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
	}
	for y := 0; y < maxX+1; y++ {
		for x := 0; x < maxY+1; x++ {
			if _, ok := tiles[Point{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// Ray Casting Algorithm
// Casting a ray on the diagonal \ as horizontal/vertical
// Rays won't work when hitting an edge.
var isInsideRgCache map[Point]bool

func (p Point) isInsideRg(rgTiles map[Point]bool) bool {
	if res, ok := isInsideRgCache[p]; ok {
		return res
	}
	x := p.x - 1
	y := p.y - 1

	crossedEdgeCount := 0
	if rgTiles[p] {
		return true
	}
	for x >= 0 && y >= 0 {
		currPoint := Point{x, y}
		if rgTiles[currPoint] {
			crossedEdgeCount++
		}
		x--
		y--
	}
	res := !(crossedEdgeCount%2 == 0)
	isInsideRgCache[p] = res
	return res
}

func rectInsideRg(p1, p2 Point, rgTiles map[Point]bool) bool {
	rectEdge := make([]Point, 0)
	p3 := Point{p1.x, p2.y}
	p4 := Point{p2.x, p1.y}
	rectEdge = append(rectEdge, getEdge(p1, p3)...)
	rectEdge = append(rectEdge, getEdge(p1, p4)...)
	rectEdge = append(rectEdge, getEdge(p2, p3)...)
	rectEdge = append(rectEdge, getEdge(p2, p4)...)

	for _, p := range rectEdge {
		if !rgTiles[p] {
			return false
		}
	}
	return true
}

func getEdge(p1, p2 Point) []Point {
	edgePoints := make([]Point, 0)
	newPoint := Point{p1.x, p1.y}
	for newPoint != p2 {
		if newPoint.y < p2.y {
			newPoint.y++
		} else if newPoint.y > p2.y {
			newPoint.y--
		}
		if newPoint.x < p2.x {
			newPoint.x++
		} else if newPoint.x > p2.x {
			newPoint.x--
		}
		edgePoints = append(edgePoints, Point{newPoint.x, newPoint.y})
	}
	return edgePoints
}

type Rect struct {
	p1, p2 Point
	area   int
}

func sortRects(r1, r2 Rect) int {
	return cmp.Compare(r1.area, r2.area)
}

func floodFill(rgTiles map[Point]bool, currPoint Point) {
	stack := []Point{currPoint}
	for len(stack) > 0 {
		currPoint = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rgTiles[currPoint] = true
		north := Point{currPoint.x, currPoint.y - 1}
		south := Point{currPoint.x, currPoint.y + 1}
		east := Point{currPoint.x + 1, currPoint.y}
		west := Point{currPoint.x - 1, currPoint.y}
		if !rgTiles[north] {
			stack = append(stack, north)
		}
		if !rgTiles[south] {
			stack = append(stack, south)
		}
		if !rgTiles[east] {
			stack = append(stack, east)
		}
		if !rgTiles[west] {
			stack = append(stack, west)
		}
	}
}

func findPointInsideRg(rgTiles map[Point]bool) Point {
	maxX := 0
	maxY := 0
	for tile, _ := range rgTiles {
		if tile.x > maxX {
			maxX = tile.x
		}
		if tile.y > maxY {
			maxY = tile.y
		}
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			newPoint := Point{x, y}
			if !rgTiles[newPoint] {
				if newPoint.isInsideRg(rgTiles) {
					return newPoint
				}
			}
		}
	}
	return Point{0, 0}
}

func getCompressedTiles(tiles []Point) ([]Point, map[Point]Point) {
	xSet := make(map[int]bool, 0)
	ySet := make(map[int]bool, 0)
	for _, point := range tiles {
		xSet[point.x] = true
		ySet[point.y] = true
	}

	xs := make([]int, 0)
	for x, _ := range xSet {
		xs = append(xs, x)
	}

	ys := make([]int, 0)
	for y, _ := range ySet {
		ys = append(ys, y)
	}

	slices.Sort(xs)
	slices.Sort(ys)

	compressedTiles := make([]Point, 0)
	tileMapping := make(map[Point]Point)
	for _, point := range tiles {
		newPoint := Point{slices.Index(xs, point.x), slices.Index(ys, point.y)}
		compressedTiles = append(compressedTiles, newPoint)
		tileMapping[newPoint] = point
	}
	return compressedTiles, tileMapping

}

// To solve part 2 we:
// Compress the input tiles to shrink the range e.g. if we had points {10, 10}, {10, 12}, {12,12}, {13, 15}
// They would become {0,0}, {0, 1}, {1, 1}, {2, 2}
// Compressing the tiles allows us to flood fill the area (not possible without the compression due to such a large space being used).
// To start the floodfill we iterate over all the points in the grid and check if it's inside the given shape by using the Ray Casting Algorithm.
// We then create a list of Rectanges - 1 for each pair of tiles given
// Calculate the area of the rectangles based on the original points (uncompressed)
// Sort the Rectangles by Area in Desc order
// Iterate through the rectangles and return the first one that is valid (all edge points are part of the original shape).
func part2(input []string) string {
	isInsideRgCache = make(map[Point]bool)
	compressedTiles, tileMapping := getCompressedTiles(parseInput(input))
	rgTiles := createRgTiles(compressedTiles)

	pointInside := findPointInsideRg(rgTiles)
	floodFill(rgTiles, pointInside)
	rects := make([]Rect, 0)
	for i := 0; i < len(compressedTiles)-1; i++ {
		p1 := compressedTiles[i]
		for j := i; j < len(compressedTiles); j++ {
			p2 := compressedTiles[j]
			uncompArea := tileMapping[p1].getArea(tileMapping[p2])
			newRect := Rect{p1, p2, uncompArea}
			rects = append(rects, newRect)
		}
	}
	slices.SortFunc(rects, sortRects)
	slices.Reverse(rects)
	for _, rect := range rects {
		if rectInsideRg(rect.p1, rect.p2, rgTiles) {
			return fmt.Sprint(rect.area)
		}
	}
	return fmt.Sprint(0)
}
