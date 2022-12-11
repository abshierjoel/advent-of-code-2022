package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abshierjoel/advent-of-code-2022/utils"
)

const filename = "moves.txt"

type Pos struct {
	x int
	y int
}

func main() {
	lines := utils.ReadLines(filename)

	// Part 1
	tailPositionsA := knottyBoy(lines, 2)
	fmt.Printf("The tail visited %d positions with a length of 2.\n", tailPositionsA)

	// Part 2
	tailPositionsB := knottyBoy(lines, 10)
	fmt.Printf("The tail visited %d positions with a length of 10.\n", tailPositionsB)
}

func knottyBoy(lines []string, ropeLength int) int {
	seen := map[Pos]struct{}{}
	rope := make([]Pos, ropeLength)

	pos := Pos{x: rope[0].x, y: rope[0].y}
	seen[pos] = struct{}{}

	for _, line := range lines {
		move := strings.Fields(line)
		direction := move[0]
		distance, _ := strconv.Atoi(move[1])

		for i := 0; i < distance; i++ {
			rope[0] = doMove(rope[0], direction)

			for j := 1; j < len(rope); j++ {
				rope[j] = pullTail(rope[j-1], rope[j])
			}

			pos := Pos{x: rope[len(rope)-1].x, y: rope[len(rope)-1].y}
			seen[pos] = struct{}{}
		}
	}

	return len(seen)
}

func pullTail(head Pos, tail Pos) Pos {
	if max(abs(head.x-tail.x), abs(head.y-tail.y)) > 1 {
		if head.x > tail.x {
			tail.x += 1
		} else if head.x < tail.x {
			tail.x -= 1
		}

		if head.y > tail.y {
			tail.y += 1
		} else if head.y < tail.y {
			tail.y -= 1
		}
	}

	return tail

}

func doMove(pos Pos, direction string) Pos {
	switch direction {
	case "L":
		pos.x -= 1
	case "U":
		pos.y -= 1
	case "R":
		pos.x += 1
	case "D":
		pos.y += 1
	}
	return pos
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
