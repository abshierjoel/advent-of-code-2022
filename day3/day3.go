package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const filename = "contents.txt"
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Rucksack struct {
	Everything     string
	CompartmentOne string
	CompartmentTwo string
}

func main() {
	lines := readLines()
	rucksacks := linesToRucksacks(lines)
	score := getMatches(rucksacks)
	badgeScore := getBadges(rucksacks)

	fmt.Println(score)
	fmt.Println(badgeScore)
}

func getBadges(rucksacks []Rucksack) int {
	var badgeScore int

	for i := 0; i < len(rucksacks); i += 3 {
		for _, char := range strings.Split(rucksacks[i].Everything, "") {
			if strings.Contains(rucksacks[i+1].Everything, char) && strings.Contains(rucksacks[i+2].Everything, char) {
				badgeScore += charToScore(char)
				break
			}
		}
	}
	return badgeScore
}

func readLines() []string {
	file, _ := ioutil.ReadFile(filename)
	file_string := string(file)
	return strings.Split(file_string, "\n")
}

func linesToRucksacks(lines []string) []Rucksack {
	var rucksacks []Rucksack
	for _, line := range lines {
		length := len(line)
		midpoint := length / 2
		rucksack := Rucksack{
			Everything:     line,
			CompartmentOne: line[:midpoint],
			CompartmentTwo: line[midpoint:],
		}
		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks
}

func getMatches(rucksacks []Rucksack) int {
	var total int
	for _, rucksack := range rucksacks {
		total += getMatch(rucksack)
	}
	return total
}

func getMatch(rucksack Rucksack) int {
	var priority int
	for i := 0; i < len(rucksack.CompartmentOne); i++ {
		char := string(rucksack.CompartmentOne[i])

		if strings.Contains(rucksack.CompartmentTwo, char) {
			priority = charToScore(char)
		}
	}

	return priority
}

func charToScore(char string) int {
	return strings.Index(letters, char) + 1
}
