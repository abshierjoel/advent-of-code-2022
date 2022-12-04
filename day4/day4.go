package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abshierjoel/advent-of-code-2022/utils"
)

const filename = "pears.txt"

type Elf struct {
	Min int
	Max int
}

type Pair struct {
	ElfA Elf
	ElfB Elf
}

func main() {
	lines := utils.ReadLines(filename)
	pairs := getPairs(lines)

	// Part 1
	pairCount := countPairs(pairs)
	fmt.Printf("%d full pairs were found!\n", pairCount)

	// Part 2
	overlappingCount := countOverlapping(pairs)
	fmt.Printf("%d overlapping pairs were found!\n", overlappingCount)
}

func countOverlapping(pairs []Pair) int {
	var overlappingCount int
	for _, pair := range pairs {
		if rangesOverlap(pair.ElfA.Min, pair.ElfA.Max, pair.ElfB.Min, pair.ElfB.Max) {
			overlappingCount += 1
		}
	}
	return overlappingCount
}

func countPairs(pairs []Pair) int {
	var pairCount int
	for _, pair := range pairs {
		if rangeFullyContains(pair.ElfA.Min, pair.ElfA.Max, pair.ElfB.Min, pair.ElfB.Max) {
			pairCount += 1
		}
	}
	return pairCount
}

func rangesOverlap(minA int, maxA int, minB int, maxB int) bool {
	if (maxA >= minB) && (minA <= maxB) {
		return true
	}

	return false
}

func rangeFullyContains(minA int, maxA int, minB int, maxB int) bool {
	if (minA >= minB && maxA <= maxB) || (minB >= minA && maxB <= maxA) {
		return true
	}

	return false
}

func getPairs(lines []string) []Pair {
	var pairs []Pair
	for _, line := range lines {
		if line != "" {
			pairs = append(pairs, getPair(line))
		}
	}
	return pairs
}

func getPair(line string) Pair {
	elves := strings.Split(line, ",")

	elfA := strings.Split(elves[0], "-")
	elfB := strings.Split(elves[1], "-")

	elfAMin, _ := strconv.Atoi(elfA[0])
	elfAMax, _ := strconv.Atoi(elfA[1])
	elfBMin, _ := strconv.Atoi(elfB[0])
	elfBMax, _ := strconv.Atoi(elfB[1])

	return Pair{
		ElfA: Elf{
			Min: elfAMin,
			Max: elfAMax,
		},
		ElfB: Elf{
			Min: elfBMin,
			Max: elfBMax,
		},
	}
}
