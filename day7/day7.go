package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/abshierjoel/advent-of-code-2022/utils"
)

const filename = "terminal.txt"
const totalSpace = 70000000
const neededSpace = 30000000
const directoryMaxSize = 100000

func main() {
	lines := utils.ReadLines(filename)

	files := buildFileList(lines)
	sizes := getDirectorySizes(files)
	rootSize := sizes["/"]

	// Part 1
	sumSmallDirectories := sumSmallDirectories(sizes)
	fmt.Println(sumSmallDirectories)

	// Part 2
	smallestDirSize := findSmallestDirectorySizeToClean(sizes, rootSize)
	fmt.Println(smallestDirSize)
}

func findSmallestDirectorySizeToClean(sizes map[string]int, rootSize int) int {
	var current = rootSize
	for _, size := range sizes {
		spaceUsed := totalSpace - (rootSize - size)

		if (current > size) && (spaceUsed > neededSpace) {
			current = size
		}
	}

	return current
}

func sumSmallDirectories(sizes map[string]int) int {
	var total int
	for _, size := range sizes {
		if size <= directoryMaxSize {
			total += size
		}
	}
	return total
}

func parseCommand(cmd string, dir string, files map[string]int) (map[string]int, string) {
	if strings.Contains(cmd, "$") {
		parts := strings.Split(cmd, " ")
		if strings.Contains(parts[1], "cd") {
			dir = path.Join(dir, parts[2])
		}
		return files, dir
	} else {
		var size int
		var name string
		fmt.Sscanf(cmd, "%d %s", &size, &name)
		files[path.Join(dir, name)] = size

		return files, dir
	}
}

func buildFileList(lines []string) map[string]int {
	dir := ""
	files := map[string]int{}

	for _, line := range lines {
		files, dir = parseCommand(line, dir, files)
	}

	return files
}

func getDirectorySizes(files map[string]int) map[string]int {
	sizes := map[string]int{}

	for f, s := range files {
		for d := path.Dir(f); d != "/"; d = path.Dir(d) {
			sizes[d] += s
		}
		sizes["/"] += s
	}

	return sizes
}
