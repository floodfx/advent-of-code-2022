package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay2(t *testing.T) {
	d := `A Y
B X
C Z`

	// split by newline for each line
	rounds := strings.Split(d, "\n")

	total := part1(rounds)
	if total != 15 {
		panic(fmt.Sprintf("Should be 15 was:%d", total))
	}
}
