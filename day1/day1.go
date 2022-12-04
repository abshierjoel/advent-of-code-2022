package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	Calories  []int
	TotalCals int
}

var filename = "calories.txt"

func main() {
	elves := getElves()
	sortElves(elves)

	// Part 1
	totalCalories := getElfCaloriesCarryingMost(elves)
	fmt.Printf("Part #1: %d\n", totalCalories)

	// Part 2
	topThreeCaloriesSummed := sumElvesCarrying(elves, 3)
	fmt.Printf("Part #2: %d\n", topThreeCaloriesSummed)
}

func getElfCaloriesCarryingMost(elves []Elf) int {
	return elves[0].TotalCals
}

func sumElvesCarrying(elves []Elf, number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		sum += elves[i].TotalCals
	}
	return sum
}

func sortElves(elves []Elf) {
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].TotalCals > elves[j].TotalCals
	})
}

func getElves() []Elf {
	var elves []Elf
	file, _ := ioutil.ReadFile(filename)
	file_string := string(file)
	counts := strings.Split(file_string, "\n\n")

	for _, elfRow := range counts {
		elf := parseElf(elfRow)
		elf.TotalCals = sumCalories(elf)
		elves = append(elves, elf)
	}
	return elves
}

func parseElf(elfData string) Elf {
	var elf Elf

	for _, calorie := range strings.Split(strings.TrimSpace(elfData), "\n") {
		calInt, _ := strconv.Atoi(calorie)
		elf.Calories = append(elf.Calories, calInt)
	}

	return elf
}

func sumCalories(elf Elf) int {
	sum := 0
	for _, cal := range elf.Calories {
		sum += cal
	}

	return sum
}
