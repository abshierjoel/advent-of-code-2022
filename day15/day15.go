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
	maxRow := 4000000

	pairs := parseSensorBeaconPairs(lines)

	// Part 1
	positionsWithoutBeacons, sensors := countPositionsWithouBeacon(pairs, row)
	fmt.Println(positionsWithoutBeacons)

	// Part 2
	tuningFreq := findTuningFreq(sensors, maxRow)
	fmt.Println(tuningFreq)
}

func findTuningFreq(sensors map[Point]int, maxRow int) int {
	for y := 0; y <= maxRow; y++ {
	loop:
		for x := 0; x <= maxRow; x++ {
			for s, diameter := range sensors {
				if abs(s.X-x)+abs(s.Y-y) <= diameter {
					x += diameter - abs(s.Y-y) + s.X - x
					continue loop
				}
			}

			fmt.Println(x)
			fmt.Println(y)
			return x*4000000 + y
		}
	}

	return 0
}

func countPositionsWithouBeacon(pairs []Pair, row int) (int, map[Point]int) {
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

	return len(matches), sensors
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
