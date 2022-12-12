package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonkeyBusiness(t *testing.T) {
	lines := strings.Split(sample(), "\n")
	monkeys := parseMonkeys(lines)
	result := findMonkeyBusiness(monkeys, 20, partAWorry)

	assert.Equal(t, 10605, result)
}

func TestBigMonkeyBusiness(t *testing.T) {
	lines := strings.Split(sample(), "\n")
	monkeys := parseMonkeys(lines)
	lcm := lcm(monkeys)

	result := findMonkeyBusiness(monkeys, 10000, partBWorry(lcm))
	assert.Equal(t, 2713310158, result)
}

func TestLCM(t *testing.T) {
	lines := strings.Split(sample(), "\n")
	monkeys := parseMonkeys(lines)
	lcm := lcm(monkeys)

	assert.Equal(t, 96577, lcm)
}

func TestParseMonkeys(t *testing.T) {
	lines := strings.Split(sample(), "\n")
	result := parseMonkeys(lines)

	expected := []Monkey{
		{
			Name:           "Monkey 0:",
			Inventory:      []int{79, 98},
			DoWorry:        (func(int) int)(nil),
			DivisibleTest:  23,
			TestPassTarget: 2,
			TestFailTarget: 3,
			ItemsInspected: 0,
		},
		{
			Name:           "Monkey 1:",
			Inventory:      []int{54, 65, 75, 74},
			DoWorry:        (func(int) int)(nil),
			DivisibleTest:  19,
			TestPassTarget: 2,
			TestFailTarget: 0,
			ItemsInspected: 0,
		},
		{
			Name:           "Monkey 2:",
			Inventory:      []int{79, 60, 97},
			DoWorry:        (func(int) int)(nil),
			DivisibleTest:  13,
			TestPassTarget: 1,
			TestFailTarget: 3,
			ItemsInspected: 0,
		},
		{
			Name:           "Monkey 3:",
			Inventory:      []int{74},
			DoWorry:        (func(int) int)(nil),
			DivisibleTest:  17,
			TestPassTarget: 0,
			TestFailTarget: 1,
			ItemsInspected: 0,
		},
	}

	assert.Equal(t, len(expected), len(result))
	// assert.Equal(t, expected, result)
}

func sample() string {
	return `		Monkey 0:
			Starting items: 79, 98
			Operation: new = old * 19
			Test: divisible by 23
				If true: throw to monkey 2
				If false: throw to monkey 3

		Monkey 1:
			Starting items: 54, 65, 75, 74
			Operation: new = old + 6
			Test: divisible by 19
				If true: throw to monkey 2
				If false: throw to monkey 0

		Monkey 2:
			Starting items: 79, 60, 97
			Operation: new = old * old
			Test: divisible by 13
				If true: throw to monkey 1
				If false: throw to monkey 3

		Monkey 3:
			Starting items: 74
			Operation: new = old + 3
			Test: divisible by 17
				If true: throw to monkey 0
				If false: throw to monkey 1
`
}
