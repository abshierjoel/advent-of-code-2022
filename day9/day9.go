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

	seen := map[Pos]struct{}{}
	head := Pos{x: 0, y: 0}
	tail := Pos{x: 0, y: 0}

	pos := Pos{x: tail.x, y: tail.y}
	seen[pos] = struct{}{}

	for _, line := range lines {
		move := strings.Fields(line)
		direction := move[0]
		distance, _ := strconv.Atoi(move[1])

		for i := 0; i < distance; i++ {
			head = doMove(head, direction)
			tail = pullTail(head, tail)

			pos := Pos{x: tail.x, y: tail.y}
			seen[pos] = struct{}{}
		}

	}

	fmt.Println(len(seen))
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
