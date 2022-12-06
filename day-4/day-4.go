package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/floodfx/advent-of-code-2022/in"
)

var d = regexp.MustCompile(`\d+`)

func checkFullContains(short, long [2]int) bool {
	return long[0] <= short[0] && long[1] >= short[1]
}

func checkAnyOverlap(a, b [2]int) bool {
	if a[0] <= b[0] && b[0] <= a[1] {
		return true
	}
	if a[0] <= b[1] && b[1] <= a[1] {
		return true
	}
	return false
}

func part1(lines []string) int {
	overlapping := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		ds := d.FindAllString(line, -1)
		// count range sizes for each
		d0, _ := strconv.Atoi(ds[0])
		d1, _ := strconv.Atoi(ds[1])
		d2, _ := strconv.Atoi(ds[2])
		d3, _ := strconv.Atoi(ds[3])
		r1len := d1 - d0
		r2len := d3 - d2

		fullContains := false
		if r1len > r2len {
			fullContains = checkFullContains([2]int{d2, d3}, [2]int{d0, d1})
		} else {
			fullContains = checkFullContains([2]int{d0, d1}, [2]int{d2, d3})
		}
		if fullContains {
			overlapping++
		}
	}
	return overlapping
}

func part2(lines []string) int {
	overlapping := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		ds := d.FindAllString(line, -1)
		d0, _ := strconv.Atoi(ds[0])
		d1, _ := strconv.Atoi(ds[1])
		d2, _ := strconv.Atoi(ds[2])
		d3, _ := strconv.Atoi(ds[3])

		fullContains := checkAnyOverlap([2]int{d2, d3}, [2]int{d0, d1}) || checkAnyOverlap([2]int{d0, d1}, [2]int{d2, d3})
		if fullContains {
			overlapping++
		}
	}
	return overlapping
}

func main() {
	lines, err := in.Lines("./day-4/input.txt")
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
