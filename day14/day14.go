package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abshierjoel/advent-of-code-2022/utils"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
)

const filename = "cave.txt"

type Point struct {
	X int
	Y int
}

type Path struct {
	StartPoint Point
	EndPoint Point
}

func main() {
	lines := utils.ReadLines(filename)
	paths := makePaths(lines)

	grid := makeGrid(paths)
	var rocks []Point
	for rock := range grid.Iterator().C {
		rocks = append(rocks, rock)
	}

	lowestPoint := lo.MaxBy(rocks, func(point Point, max Point) bool {
		return (point.Y > max.Y)
	})

	// Part 1
	fallCount := fillWithSand(grid, lowestPoint.Y + 1, 0, false)
	fmt.Println(fallCount)

	// Part 2
	floorfallCount := fillWithSand(grid, lowestPoint.Y + 1, 0, true)
	fmt.Println(floorfallCount)
}

func fillWithSand(grid mapset.Set[Point], lowestPoint int, count int, floor bool) int {
	fillPoint := Point{X: 500, Y: 0}

	if grid.Contains(fillPoint) {
		return count
	}

	grid, success := addSand(grid, lowestPoint, fillPoint, floor)

	if success {
		return fillWithSand(grid, lowestPoint, count + 1, floor)
	}

	return count
}

func addSand(grid mapset.Set[Point], lowestPoint int, point Point, floor bool) (mapset.Set[Point], bool) {
	pos := dropSand(grid, lowestPoint, &point, floor)

	if pos == nil {
		return grid, false
	}
	grid.Add(*pos)

	return grid, true
}

func dropSand(grid mapset.Set[Point], lowestPoint int, point *Point, floor bool) *Point {
	nextPosition := findNextPosition(grid, lowestPoint, point, floor)

	if nextPosition == nil { 
		return nil
	}
	if nextPosition.X == point.X && nextPosition.Y == point.Y {
		return point
	}

	return dropSand(grid, lowestPoint, nextPosition, floor)
}

func findNextPosition(grid mapset.Set[Point], lowestPoint int, point *Point, floor bool) *Point {
	if point.Y == lowestPoint {
		if floor {
			return point
		} else {
			return nil
		}
	}
	if(!grid.Contains(Point{X: point.X, Y: point.Y + 1})) {
		return &Point{X: point.X, Y: point.Y + 1}
	}
	if(!grid.Contains(Point{X: point.X - 1, Y: point.Y + 1})) {
		return &Point{X: point.X - 1, Y: point.Y + 1}
	}
	if(!grid.Contains(Point{X: point.X + 1, Y: point.Y + 1})) {
		return &Point{X: point.X + 1, Y: point.Y + 1}
	}
	return point
}


func makeGrid(paths []Path) mapset.Set[Point] {
	set := mapset.NewSet[Point]()
	points := lo.FlatMap[Path, Point](paths, func(path Path, _ int) []Point {
		var points []Point
		maxX := lo.Max([]int{path.StartPoint.X, path.EndPoint.X})
		maxY := lo.Max([]int{path.StartPoint.Y, path.EndPoint.Y})
		minX := lo.Min([]int{path.StartPoint.X, path.EndPoint.X})
		minY := lo.Min([]int{path.StartPoint.Y, path.EndPoint.Y})

		for i := minX; i <= maxX; i++ {
			points = append(points, Point{X: i, Y: minY})
		}

		for i := minY; i <= maxY; i++ {
			points = append(points, Point{X: minX, Y: i})
		}

		return points
	})

	lo.ForEach(points, func(point Point, _ int) { set.Add(point) })

	return set
}

func makePaths(lines []string) []Path {
	return lo.FlatMap[string, Path](lines, func(line string, _ int) []Path {
		parts := strings.Split(line, " -> ")
		var paths []Path

		for i := 0; i < len(parts) - 1; i++ {
			start := strings.Split(parts[i], ",")
			end := strings.Split(parts[i+1], ",")
			startX,_ := strconv.Atoi(start[0])
			startY,_ := strconv.Atoi(start[1])
			endX,_ := strconv.Atoi(end[0])
			endY,_ := strconv.Atoi(end[1])

			paths = append(paths, Path{
				StartPoint: Point{X: startX, Y: startY}, 
				EndPoint: Point{X: endX, Y: endY},
			})
		}

		return paths
	})
}

func sample() []string {
	return []string{
		"498,4 -> 498,6 -> 496,6",
		"503,4 -> 502,4 -> 502,9 -> 494,9",
	}
}
