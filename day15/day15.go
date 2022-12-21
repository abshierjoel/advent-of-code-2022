package main

import (
	"fmt"

	"github.com/abshierjoel/advent-of-code-2022/utils"
)

type Point struct {
	X int
	Y int
}

type Pair struct {
	Sensor Point
	Beacon Point
}

func main() {
	lines := utils.ReadLines("sensors.txt")
	row := 2000000

	pairs := parseSensorBeaconPairs(lines)

	// Part 1
	positionsWithoutBeacons := countPositionsWithouBeacon(pairs, row)
	fmt.Println(positionsWithoutBeacons)
}

func countPositionsWithouBeacon(pairs []Pair, row int) int {
	sensors := map[Point]int{}
	matches := map[int]struct{}{}
	for _, pair := range pairs {
		s := pair.Sensor
		b := pair.Beacon

		sensors[s] = abs(s.X-b.X) + abs(s.Y-b.Y)
		diameter := sensors[s] - abs(row-s.Y)

		for x := s.X - diameter; x <= s.X+diameter; x++ {
			if !(x == b.X && b.Y == row) {
				matches[x] = struct{}{}
			}
		}
	}

	return len(matches)
}

func parseSensorBeaconPairs(lines []string) []Pair {
	var pairs []Pair
	for _, line := range lines {
		var s, b Point
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.X, &s.Y, &b.X, &b.Y)

		pairs = append(pairs, Pair{
			Sensor: s,
			Beacon: b,
		})
	}
	return pairs
}

func abs(value int) int {
	if value >= 0 {
		return value
	}
	return value * (-1)
}

func sample() []string {
	return []string{
		"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
		"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
		"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
		"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
		"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
		"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
		"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
		"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
		"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
		"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
		"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
		"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
		"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
		"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
	}
}
