package main

import (
	"fmt"

	"github.com/abshierjoel/advent-of-code-2022/utils"
)

const filename = "datastream.txt"

func main() {
	lines := utils.ReadLines(filename)
	// lines := sampleBuffer()

	for _, buffer := range lines {
		startPos := findPacketStart(buffer)
		fmt.Printf("Found start of packet at %d\n", startPos)
	}
}

func findPacketStart(buffer string) int {
	var queue []string
	for ind, char := range buffer {
		if len(queue) == 4 && isUniqueList(queue) {
			return ind
		}

		queue = pushIn(queue, string(char))
	}

	return -1
}

func isUniqueList(list []string) bool {
	distinct := make(map[string]bool)

	for _, elem := range list {
		distinct[elem] = true
	}

	return len(distinct) == len(list)
}

func pushIn(queue []string, value string) []string {
	if len(queue) == 4 {
		queue = queue[1:]
	}
	return append(queue, value)
}

func sampleBuffer() []string {
	return []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	}
}
