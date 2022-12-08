package main

import (
	"fmt"
	"strconv"

	"github.com/abshierjoel/advent-of-code-2022/utils"
)

const filename = "patch.txt"

func main() {
	lines := utils.ReadLines(filename)
	patch := parseTreePatch(lines)

	// Part 1
	visibleTrees := getVisibleTrees(patch)
	fmt.Printf("Visible Trees %d\n", visibleTrees)
}

func getVisibleTrees(patch [][]int) int {
	visibleCount := 0
	rows := len(patch)
	cols := len(patch[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			height := patch[i][j]
			fromLeft, fromRite, fromTop, fromBot := getSurroundings(patch, i, j)

			equalFunc := func(a int) bool { return a < height }

			if allMatch(fromLeft, equalFunc) || allMatch(fromRite, equalFunc) || allMatch(fromTop, equalFunc) || allMatch(fromBot, equalFunc) {
				visibleCount++
			}
		}
	}

	return visibleCount
}

func getColumn(patch [][]int, col int) []int {
	column := make([]int, 0)
	for _, row := range patch {
		column = append(column, row[col])
	}
	return column
}

func getSurroundings(patch [][]int, i int, j int) ([]int, []int, []int, []int) {
	fromLeft := patch[i][:j]
	fromRite := patch[i][(j + 1):]
	fromTop := getColumn(patch, j)[:i]
	fromBot := getColumn(patch, j)[(i + 1):]

	return fromLeft, fromRite, fromTop, fromBot
}

func allMatch(trees []int, matchFunc func(int) bool) bool {
	for _, tree := range trees {
		if !matchFunc(tree) {
			return false
		}
	}
	return true
}

func parseTreePatch(rows []string) [][]int {
	patch := make([][]int, len(rows))

	for i, row := range rows {
		patch[i] = make([]int, len(row))

		for j := 0; j < len(row); j++ {
			patch[i][j] = byteToInt(row[j])
		}

	}

	return patch
}

func byteToInt(aByte byte) int {
	aInt, _ := strconv.Atoi(string(aByte))
	return aInt
}
