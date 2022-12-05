package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const filename = "stacks.txt"

type Instruction struct {
	Count int
	From  int
	To    int
}

func main() {
	// file := readFile()
	file := readSample()
	raw_stacks, raw_instructions := parseParts(file)

	fmt.Printf("-- Stacks -- \n%s\n\n", raw_stacks)
	fmt.Printf("-- Instructions -- \n%s\n\n", raw_instructions)

	stacks := buildStacks(raw_stacks)
	instructions := buildInstructions(raw_instructions)
	reordered_stacks := executeInstructions(instructions, stacks)

	// Part 1
	stack_tops := getTopOfStacks(reordered_stacks)

	fmt.Println(stack_tops)
}

func getTopOfStacks(stacks map[int][]string) string {
	var tops string
	for i := 1; i <= len(stacks); i++ {
		top := stacks[i][len(stacks[i])-1]
		tops += top
	}
	return tops
}

func executeInstructions(instructions []Instruction, stacks map[int][]string) map[int][]string {
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
