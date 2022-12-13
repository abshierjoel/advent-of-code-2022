package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/abshierjoel/advent-of-code-2022/utils"
)

const filename = "packets.txt"

type Packet struct {
	Value    int
	Elements []*Packet
	Parent   *Packet
}

type Pair struct {
	PacketA *Packet
	PacketB *Packet
}

func main() {
	lines := utils.ReadLines(filename)
	pairs := parsePackets(lines)

	// Part 1
	printOrderedIndexSum(pairs)

	// Part 2
	var allPackets []*Packet
	for _, pair := range pairs {
		allPackets = append(allPackets, pair.PacketA)
		allPackets = append(allPackets, pair.PacketB)
	}

	allPackets = append(allPackets, parsePacket("[[2]]"))
	allPackets = append(allPackets, parsePacket("[[6]]"))

	sort.Slice(allPackets, func(i, j int) bool {
		return isInOrder(allPackets[i], allPackets[j]) == 1
	})

	key := 1
	for i, packet := range allPackets {
		fmt.Println(packet)
		if isInOrder(packet, parsePacket("[[2]]")) == 0 || isInOrder(packet, parsePacket("[[6]]")) == 0 {
			fmt.Println(i)
			key *= i + 1
		}
	}

	fmt.Println(key)
}

func printOrderedIndexSum(pairs []Pair) {
	var results []int
	for i, pair := range pairs {
		if isInOrder(pair.PacketA, pair.PacketB) >= 0 {
			results = append(results, i+1)
		}
	}

	var sum int
	for _, v := range results {
		sum = sum + v
	}

	fmt.Println("Sum of Ordered Packet Indices")
	fmt.Println(sum)
}

func isInOrder(a *Packet, b *Packet) int {
	switch {
	case len(a.Elements) == 0 && len(b.Elements) == 0:
		if a.Value < b.Value {
			return 1
		} else if a.Value == b.Value {
			return 0
		} else {
			return -1
		}
	case a.Value >= 0:
		return isInOrder(&Packet{-1, []*Packet{a}, nil}, b)
	case b.Value >= 0:
		return isInOrder(a, &Packet{-1, []*Packet{b}, nil})
	default:
		var i int
		for i = 0; i < len(a.Elements) && i < len(b.Elements); i++ {
			order := isInOrder(a.Elements[i], b.Elements[i])
			if order != 0 {
				return order
			}
		}
		if i < len(a.Elements) {
			return -1
		} else if i < len(b.Elements) {
			return 1
		}
	}

	return 0
}

func parsePackets(lines []string) []Pair {
	var pairs []Pair
	var one *Packet
	var two *Packet

	for _, line := range lines {
		if line == "" {
			pairs = append(pairs, Pair{one, two})
			one, two = nil, nil
		} else if one == nil {
			one = parsePacket(line)
		} else if two == nil {
			two = parsePacket(line)
		}
	}

	return pairs
}

func parsePacket(line string) *Packet {
	packet := Packet{Value: -1, Elements: []*Packet{}, Parent: nil}
	currentParent := &packet

	var numStr string
	for _, c := range line {
		switch c {
		case '[':
			node := Packet{-1, []*Packet{}, currentParent}
			currentParent.Elements = append(currentParent.Elements, &node)
			currentParent = &node
		case ']':
			if len(numStr) > 0 {
				number, _ := strconv.Atoi(numStr)
				currentParent.Value = number
				numStr = ""
			}
			currentParent = currentParent.Parent
		case ',':
			if len(numStr) > 0 {
				number, _ := strconv.Atoi(numStr)
				currentParent.Value = number
				numStr = ""
			}
			currentParent = currentParent.Parent
			node := Packet{-1, []*Packet{}, currentParent}
			currentParent.Elements = append(currentParent.Elements, &node)
			currentParent = &node
		default:
			numStr += string(c)
		}
	}

	return &packet
}
