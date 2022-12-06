package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/abshierjoel/advent-of-code-2022/utils"
)

const filename = "stacks.txt"

type Instruction struct {
	Count int
	From  int
	To    int
}

func main() {
	file := readFile()
	raw_stacks, raw_instructions := parseParts(file)

	stacks := buildStacks(raw_stacks)
	instructions := buildInstructions(raw_instructions)

	// Part 1
	part1_stacks := utils.CopyMap(stacks)
	reordered_stacks := executeCrateMover9000(instructions, part1_stacks)
	stack_tops := getTopOfStacks(reordered_stacks)
	fmt.Println("\nPart 1 Stack Tops")
	fmt.Println(stack_tops)

	// Part 2
	part2_stacks := utils.CopyMap(stacks)
	reordered_stacks_two := executeCrateMover9001(instructions, part2_stacks)
	stack_tops_two := getTopOfStacks(reordered_stacks_two)
	fmt.Println("\nPart 2 Stack Tops")
	fmt.Println(stack_tops_two)
}

func getTopOfStacks(stacks map[int][]string) string {
	var tops string
	for i := 1; i <= len(stacks); i++ {
		top := stacks[i][len(stacks[i])-1]
		tops += top
	}
	return tops
}

func executeCrateMover9001(instructions []Instruction, stacks map[int][]string) map[int][]string {
	for _, inst := range instructions {
		count := inst.Count
		from := inst.From
		to := inst.To

		val := stacks[from][len(stacks[from])-count : len(stacks[from])]
		stacks[to] = append(stacks[to], val...)
		stacks[from] = pop_n(stacks[from], count)
	}

	return stacks
}

func executeCrateMover9000(instructions []Instruction, stacks map[int][]string) map[int][]string {
	for _, inst := range instructions {
		from := inst.From
		to := inst.To

		for i := 0; i < inst.Count; i++ {
			var val string
			stacks[from], val = pop(stacks[from])
			stacks[to] = push(stacks[to], val)
		}
	}

	return stacks
}

func buildInstructions(raw_instructions string) []Instruction {
	instruction_rows := strings.Split(raw_instructions, "\n")

	var instructions []Instruction
	for _, inst := range instruction_rows {
		fields := strings.Fields(inst)
		count, _ := strconv.Atoi(fields[1])
		from, _ := strconv.Atoi(fields[3])
		to, _ := strconv.Atoi(fields[5])

		instructions = append(instructions, Instruction{
			Count: count,
			From:  from,
			To:    to,
		})
	}

	return instructions
}

func buildStacks(raw_stacks string) map[int][]string {
	stack_rows := strings.Split(raw_stacks, "\n")
	stacks := map[int][]string{}
	for _, row := range stack_rows[:len(stack_rows)-1] {
		// var stack []string
		for j := 0; j < len(row); j += 4 {
			crateNum := j/4 + 1
			if row[j+1] != ' ' {
				val := string(row[j+1])
				stacks[crateNum] = append([]string{val}, stacks[crateNum]...)
			}
		}

	}
	return stacks
}

func pop(stack []string) ([]string, string) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}

func pop_n(stack []string, count int) []string {
	return stack[:len(stack)-count]
}

func push(stack []string, crate string) []string {
	stack = append(stack, crate)
	return stack
}

func parseParts(file string) (string, string) {
	parts := strings.Split(file, "\n\n")
	return parts[0], parts[1]
}

func readSample() string {
	file, _ := ioutil.ReadFile("sample_stacks.txt")
	return string(file)
}

func readFile() string {
	file, _ := ioutil.ReadFile(filename)
	return string(file)
}
