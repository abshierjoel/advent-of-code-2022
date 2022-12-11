package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var filename = "strategy.txt"

type Turn struct {
	YourHand  string
	TheirHand string
}

func main() {
	lines := readLines()

	// Part 1
	partOneTurns := linesToTurns(lines, partOneHand)
	partOneScore := scoreGame(partOneTurns)
	fmt.Printf("You scored %d in Part 1\n", partOneScore)

	// Part 2
	partTwoTurns := linesToTurns(lines, partTwoHand)
	partTwoScore := scoreGame(partTwoTurns)
	fmt.Printf("You scored %d in Part 2\n", partTwoScore)

	fmt.Println("Expected Output:")
	fmt.Println("You scored 14827 in Part 1")
	fmt.Println("You scored 13889 in Part 2")
}

func scoreGame(turns []Turn) int {
	var totalScore int
	for i := 0; i < len(turns); i++ {
		totalScore += scoreYourHand(turns[i].YourHand)
		totalScore += getResults(turns[i])
	}
	return totalScore
}

func getResults(turn Turn) int {
	if turn.YourHand == turn.TheirHand {
		return 3
	} else if turn.YourHand == "rock" && turn.TheirHand == "scissors" {
		return 6
	} else if turn.YourHand == "paper" && turn.TheirHand == "rock" {
		return 6
	} else if turn.YourHand == "scissors" && turn.TheirHand == "paper" {
		return 6
	} else {
		return 0
	}
}

func linesToTurns(lines []string, handFunc func(string, string) string) []Turn {
	var turns []Turn
	for i := 0; i < len(lines); i++ {
		currentTurn := toTurn(lines[i], handFunc)
		turns = append(turns, currentTurn)
	}
	return turns
}

func toTurn(move string, handFunc func(string, string) string) Turn {
	hands := strings.Split(move, " ")

	return Turn{
		YourHand:  handFunc(hands[0], hands[1]),
		TheirHand: toTheirHand(hands[0]),
	}
}

func partOneHand(their_hand string, your_hand string) string {
	switch your_hand {
	case "X":
		return "rock"
	case "Y":
		return "paper"
	case "Z":
		return "scissors"
	default:
		return ""
	}
}

func partTwoHand(their_hand string, your_hand string) string {
	switch your_hand {
	case "X":
		return losingMove(their_hand)
	case "Y":
		return toTheirHand(their_hand)
	case "Z":
		return winningMove(their_hand)
	default:
		return ""
	}
}

func losingMove(hand string) string {
	switch hand {
	case "A":
		return "scissors"
	case "B":
		return "rock"
	case "C":
		return "paper"
	default:
		return ""
	}
}

func winningMove(hand string) string {
	switch hand {
	case "A":
		return "paper"
	case "B":
		return "scissors"
	case "C":
		return "rock"
	default:
		return ""
	}
}

func toTheirHand(hand string) string {
	switch hand {
	case "A":
		return "rock"
	case "B":
		return "paper"
	case "C":
		return "scissors"
	default:
		return ""
	}
}

func scoreYourHand(hand string) int {
	switch hand {
	case "rock":
		return 1
	case "paper":
		return 2
	case "scissors":
		return 3
	default:
		return 0
	}
}

func readLines() []string {
	file, _ := ioutil.ReadFile(filename)
	file_string := string(file)
	return strings.Split(file_string, "\n")
}
