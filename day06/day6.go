package main

import (
	"fmt"

	"github.com/abshierjoel/advent-of-code-2022/utils"
)

const filename = "datastream.txt"

type Queue struct {
	Queue  []string
	MaxLen int
}

func main() {
	lines := utils.ReadLines(filename)

	for _, buffer := range lines {
		// Part 1
		packetStartPos := findPacketStart(buffer, 4)
		fmt.Printf("Found start of packet at %d\n", packetStartPos)

		// Part 2
		messageStartPos := findPacketStart(buffer, 14)
		fmt.Printf("Found start of message at %d\n", messageStartPos)
	}
}

func findPacketStart(buffer string, distinctChars int) int {
	queue := Queue{Queue: []string{}, MaxLen: distinctChars}

	for ind, char := range buffer {
		if len(queue.Queue) == distinctChars && isUniqueList(queue.Queue) {
			return ind
		}

		queue = pushIn(queue, string(char))
	}

	return -1
}

func isUniqueList(list []string) bool {
	distinct := make(map[string]struct{}, len(list))

	for _, elem := range list {
		distinct[elem] = struct{}{}
	}

	return len(distinct) == len(list)
}

func pushIn(queue Queue, value string) Queue {
	if len(queue.Queue) == queue.MaxLen {
		queue.Queue = queue.Queue[1:]
	}
	queue.Queue = append(queue.Queue, value)
	return queue
}
