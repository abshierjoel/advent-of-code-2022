package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abshierjoel/advent-of-code-2022/utils"
)

type CPU struct {
	Cycles int
	X      int
	Total  int
}

const filename = "instructions.txt"

func main() {
	cpu := newCpu()
	instructions := utils.ReadLines(filename)
	signalStrength := cpu.getSignalStrengthSum(instructions)

	fmt.Printf("\n\nSignal Strength: %d\n", signalStrength)
}

func (cpu *CPU) getSignalStrengthSum(instructions []string) int {
	for _, ins := range instructions {
		fields := strings.Fields(ins)

		switch fields[0] {
		case "noop":
			cpu.noop()
		case "addx":
			cpu.addx(fields[1])
		}
	}

	return cpu.Total
}

func (cpu *CPU) addx(addend_str string) {
	cpu.noop()

	addend, _ := strconv.Atoi(addend_str)
	cpu.X += addend
	cpu.Cycles++

	cpu.cycle(addend)
}

func (cpu *CPU) noop() {
	cpu.Cycles++
	cpu.cycle(0)
}

func (cpu *CPU) cycle(add int) {
	xPosition := (cpu.Cycles % 40)
	if xPosition >= cpu.X-1 && xPosition <= cpu.X+1 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}

	if (cpu.Cycles % 40) == 0 {
		fmt.Printf("\n")
	}

	if ((cpu.Cycles+1)%40)-20 == 0 {
		cpu.Total += (cpu.Cycles + 1) * cpu.X
	}
}

func newCpu() CPU { return CPU{X: 1, Cycles: 0} }

// RUAKHBEK
