package main

import (
	"fmt"

	"github.com/floodfx/advent-of-code-2022/in"
)

func checkUniqueN(in string, n int) bool {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if in[i] == in[j] {
				return false
			}
		}
	}
	return true
}

func checkUnique(in string) bool {
	if in[0] != in[1] && in[0] != in[2] && in[0] != in[3] &&
		in[1] != in[2] && in[1] != in[3] &&
		in[2] != in[3] {
		return true
	}
	return false
}

func part1(lines []string) int {
	last14 := ""
	count := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		i := 0
		if last14 == "" {
			last14 = line[0:14]
			count += 14
			i += 14
		}
		if checkUnique(last14) {
			return count
		}
		for j := i; j < len(line); j++ {
			last14 += string(line[j])
			last14 = last14[1:]
			count++
			if checkUniqueN(last14, 14) {
				return count
			}
		}
	}
	return count
}

func main() {
	lines, err := in.Lines("./day-6/input.txt")
	if err != nil {
		panic(err)
	}

	count := part1(lines)
	fmt.Printf("\n\nPart 1 - Count is: %d\n", count)

	// o2 := &Output{}
	// msg = part2(lines, o2)
	// fmt.Printf("\n\nPart 2 - Msg is: %s, o2:%v\n", msg, o2)
}
