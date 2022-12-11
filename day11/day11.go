package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/abshierjoel/advent-of-code-2022/utils"
)

const filename = "monkeys.txt"

type Monkey struct {
	Name           string
	Inventory      []int
	DoWorry        func(old int) int
	DivisibleTest  int
	TestPassTarget int
	TestFailTarget int
	ItemsInspected int ``
}

func (monkey *Monkey) addItem(item int) {
	monkey.Inventory = append(monkey.Inventory, item)
}

func (monkey *Monkey) popItem() int {
	worryLvl := monkey.Inventory[0]
	monkey.Inventory = monkey.Inventory[1:]
	return worryLvl
}

func main() {
	lines := utils.ReadLines(filename)
	monkeys := parseMonkeys(lines)

	// Part 1
	monkeyBusiness := findMonkeyBusiness(monkeys, 20)
	fmt.Printf("\nMonkey Business: %d\n\n", monkeyBusiness)
}

func findMonkeyBusiness(monkeys []Monkey, rounds int) int {
	for round := 1; round <= rounds; round++ {
		monkeys = doRound(monkeys)
	}
	sortByItemsInspected(monkeys)

	return monkeys[0].ItemsInspected * monkeys[1].ItemsInspected
}

func sortByItemsInspected(monkeys []Monkey) {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].ItemsInspected > monkeys[j].ItemsInspected
	})
}

func doRound(monkeys []Monkey) []Monkey {
	for i := 0; i < len(monkeys); i++ {
		for len(monkeys[i].Inventory) > 0 {
			item := monkeys[i].popItem()
			worryLvl := monkeys[i].DoWorry(item) / 3
			target := monkeys[i].findTarget(worryLvl)
			monkeys[target].addItem(worryLvl)

			monkeys[i].ItemsInspected++
		}
	}

	return monkeys
}

func (monkey *Monkey) findTarget(worryLvl int) int {
	var target int
	if (worryLvl % monkey.DivisibleTest) == 0 {
		target = monkey.TestPassTarget
	} else {
		target = monkey.TestFailTarget
	}
	return target
}

func parseMonkeys(lines []string) []Monkey {
	var monkeys []Monkey
	var monkey Monkey

	for _, line := range lines {
		lineType := getLineType(line)
		lineParts := strings.Fields(line)

		switch lineType {
		case "new":
			monkey = Monkey{
				Name: "Monkey " + fmt.Sprint(lineParts[1]),
			}
		case "inventory":
			var inv []int
			for _, item_str := range lineParts[2:] {
				trimmed := strings.Trim(item_str, ",")
				item, _ := strconv.Atoi(trimmed)
				inv = append(inv, item)
			}
			monkey.Inventory = inv
		case "operation":
			operator := lineParts[4]
			argument := lineParts[5]
			intArg, _ := strconv.Atoi(argument)

			switch operator {
			case "*":
				if argument == "old" {
					monkey.DoWorry = square
					break
				}
				monkey.DoWorry = multiply(intArg)
			case "+":
				monkey.DoWorry = add(intArg)
			case "-":
				monkey.DoWorry = subtract(intArg)
			case "/":
				monkey.DoWorry = divide(intArg)
			}
		case "test":
			test, _ := strconv.Atoi(lineParts[3])
			monkey.DivisibleTest = test
		case "pass":
			pass, _ := strconv.Atoi(lineParts[5])
			monkey.TestPassTarget = pass
		case "fail":
			fail, _ := strconv.Atoi(lineParts[5])
			monkey.TestFailTarget = fail
		case "insert":
			monkeys = append(monkeys, monkey)
		default:
			fmt.Println("There was an error")
			fmt.Printf("`%s`", line)
		}
	}

	return monkeys
}

func getLineType(line string) string {
	if strings.Contains(line, "Monkey") {
		return "new"
	}
	if strings.Contains(line, "Starting") {
		return "inventory"
	}
	if strings.Contains(line, "Operation") {
		return "operation"
	}
	if strings.Contains(line, "Test") {
		return "test"
	}
	if strings.Contains(line, "If true") {
		return "pass"
	}
	if strings.Contains(line, "If false") {
		return "fail"
	}
	if line == "" {
		return "insert"
	}
	return "error"
}

func add(y int) func(int) int {
	return func(x int) int {
		return x + y
	}
}

func subtract(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func multiply(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func divide(x int) func(int) int {
	return func(y int) int {
		return y / x
	}
}

func square(x int) int {
	return x * x
}
