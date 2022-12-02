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
	turns := linesToTurns(lines)

	// Part 1
	partOneScore := partOne(turns)
	fmt.Printf("You scored %d in Part 1\n", partOneScore)

	// Part 2
	partTwoScore := partTwo(turns)
	fmt.Printf("You scored %d in Part 2\n", partTwoScore)
}

func partTwo(turns []Turn) int {
	newTurns := updateTurns(turns)

	var totalScore int
	for i := 0; i < len(newTurns); i++ {
		totalScore += scoreYourHand(newTurns[i].YourHand)
		totalScore += getPartTwoResult(newTurns[i])
	}
	return totalScore
}

func updateTurns(turns []Turn) []Turn {
	var newTurns []Turn
	for i := 0; i < len(turns); i++ {
		newTurns = append(newTurns, updateTurn(turns[i]))
	}

	return newTurns
}

func updateTurn(turn Turn) Turn {
	newTurn := turn

	if turn.YourHand == "rock" {
		newTurn.YourHand = losingMove(turn.TheirHand)
	} else if turn.YourHand == "paper" {
		newTurn.YourHand = turn.TheirHand
	} else if turn.YourHand == "scissors" {
		newTurn.YourHand = winningMove(turn.TheirHand)
	}

	return newTurn
}

func losingMove(hand string) string {
	switch hand {
	case "rock":
		return "scissors"
	case "paper":
		return "rock"
	case "scissors":
		return "paper"
	default:
		return ""
	}
}

func winningMove(hand string) string {
	switch hand {
	case "rock":
		return "paper"
	case "paper":
		return "scissors"
	case "scissors":
		return "rock"
	default:
		return ""
	}
}

func getPartTwoResult(turn Turn) int {
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

func partOne(turns []Turn) int {
	var totalScore int
	for i := 0; i < len(turns); i++ {
		totalScore += scoreYourHand(turns[i].YourHand)
		totalScore += getPartOneResult(turns[i])
	}
	return totalScore
}

func getPartOneResult(turn Turn) int {
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

func linesToTurns(lines []string) []Turn {
	var turns []Turn
	for i := 0; i < len(lines); i++ {
		currentTurn := toTurn(lines[i])
		turns = append(turns, currentTurn)
	}
	return turns
}

func toTurn(move string) Turn {
	hands := strings.Split(move, " ")

	return Turn{
		YourHand:  toHand(hands[1]),
		TheirHand: toHand(hands[0]),
	}
}

func toHand(hand string) string {
	switch hand {
	case "X":
		return "rock"
	case "Y":
		return "paper"
	case "Z":
		return "scissors"
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
