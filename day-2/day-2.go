package main

import (
	"fmt"
	"os"
	"strings"
)

// Input:
// Col 1 - A,B,C (Rock, Paper, Scissors)
// Col 2 - X,Y,Z (Rock, Paper, Scissors)

// Round Score = ShapeScore+OutcomeScore
// ShapeScore = 1 for Rock, 2 for Paper, 3 for Scissors
// OutcomeScore = 6 for Win, 3 for Draw, 0 for Loss
// Win Rules - Rock beats Scissors, Scissors beats Paper, Paper beats Rock

const (
	ROCK = iota
	PAPER
	SCISSORS
)

// index 0 is rock, 1 is paper, 2 is scissors
var evalMatrix = [][]int{
	{3, 6, 0},
	{0, 3, 6},
	{6, 0, 3},
}

func inputToShape(input string) int {
	switch input {
	case "A", "X":
		return ROCK
	case "B", "Y":
		return PAPER
	case "C", "Z":
		return SCISSORS
	default:
		panic("Invalid input")
	}
}

func roundToScore(their, my int) int {
	return evalMatrix[their][my] + my + 1
}

func part1(rounds []string) int {
	totalScore := 0
	for _, round := range rounds {
		if round == "" {
			continue
		}
		// split each line by space for each input
		inputs := strings.Split(round, " ")
		// convert inputs to shapes
		theirShape := inputToShape(inputs[0])
		myShape := inputToShape(inputs[1])
		// get score for this round
		round := roundToScore(theirShape, myShape)
		totalScore += round
		// fmt.Printf("total: %d, Round: %d\n", totalScore, round)
	}
	return totalScore
}

// Input Part 2:
// Col 1 - A,B,C (Rock, Paper, Scissors)
// Col 2 - X,Y,Z (Lose, Draw, Win)
func inputToOutcome(in string) int {
	switch in {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	default:
		panic("Invalid input")
	}
}

func outcomeToShape(their, outcome int) int {
	cols := evalMatrix[their]
	for i, col := range cols {
		if col == outcome {
			return i
		}
	}
	panic("Should have found the outcome")
}

func part2(rounds []string) int {
	total := 0
	for _, round := range rounds {
		if round == "" {
			continue
		}
		inputs := strings.Split(round, " ")
		their := inputToShape(inputs[0])
		my := outcomeToShape(their, inputToOutcome(inputs[1]))
		total += roundToScore(their, my)
	}
	return total
}

func main() {
	// read in input.txt to string
	b, err := os.ReadFile("./day-2/input.txt")
	if err != nil {
		panic(err)
	}
	d := string(b)

	// part 1
	rounds := strings.Split(d, "\n")
	totalScore := part1(rounds)
	fmt.Printf("\n\nPart 1 - Total Score is: %d", totalScore)

	// part 2
	totalScore = part2(rounds)
	fmt.Printf("\n\nPart 2 - Total Score is: %d", totalScore)
}
