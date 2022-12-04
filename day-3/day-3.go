package main

import (
	"fmt"

	"github.com/floodfx/advent-of-code-2022/in"
)

// Input:
// Each line has two sets of items which are half of the line each
// Each character is a an item
// Each line pair shares one item
// Priority of item char: a-z = 1-26, A-Z = 27-52
// Goal -
// - find the shared item in each line pair
// - figure out the priority of the shared item
// - sum the priority of all shared items

var charPriority map[rune]int

// build charPriority
func init() {
	charPriority = make(map[rune]int)
	p := 0
	for r := 'a'; r <= 'z'; r++ {
		charPriority[r] = p
		p++
	}
	for r := 'A'; r <= 'Z'; r++ {
		charPriority[r] = p
		p++
	}
}

func part1(lines []string) int {
	totalScore := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		half := len(line) / 2
		items1 := line[:half]
		items2 := line[half:]
		item1Runes := make(map[rune]bool, len(items1))
		// collect item1 runes first
		for i := 0; i < half; i++ {
			r := rune(items1[i])
			item1Runes[r] = true
		}
		// loop through item2 runes to find match
		var sharedRune rune
		for i := 0; i < half; i++ {
			r := rune(items2[i])
			if item1Runes[r] {
				sharedRune = r
				break
			}
		}
		p := charPriority[sharedRune] + 1
		// fmt.Printf("\nShared rune: %q, priority: %d", sharedRune, p)
		// add priority of sharedRune to score
		totalScore += p
	}
	return totalScore
}

// Group lines by 3
// Find common rune between each line
// Find priority of common rune
func part2(lines []string) int {
	totalScore := 0
	for i := 0; i < len(lines); i += 3 {
		if lines[i] == "" {
			continue
		}
		items1 := lines[i]
		items2 := lines[i+1]
		items3 := lines[i+2]
		// add all runes from first items
		itemRunes := make(map[rune]bool, len(items1))
		// collect item1 runes first
		for i := 0; i < len(items1); i++ {
			r := rune(items1[i])
			itemRunes[r] = true
		}
		// loop through items2 runes to find match
		sharedRunes := make(map[rune]bool)
		for i := 0; i < len(items2); i++ {
			r := rune(items2[i])
			sharedRunes[r] = itemRunes[r]
		}
		// loop through items3 runes to fine match from shared
		var sharedRune rune
		for i := 0; i < len(items3); i++ {
			r := rune(items3[i])
			if sharedRunes[r] {
				sharedRune = r
				break
			}
		}
		p := charPriority[sharedRune] + 1
		// fmt.Printf("\nShared rune: %q, priority: %d", sharedRune, p)
		// add priority of sharedRune to score
		totalScore += p
	}
	return totalScore
}

func main() {
	lines, err := in.Lines("./day-3/input.txt")
	if err != nil {
		panic(err)
	}

	// part 1
	totalScore := part1(lines)
	fmt.Printf("\n\nPart 1 - Total Score is: %d", totalScore)

	// part 2
	totalScore = part2(lines)
	fmt.Printf("\n\nPart 2 - Total Score is: %d", totalScore)
}
