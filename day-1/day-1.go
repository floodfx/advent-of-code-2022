package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// read in input.txt to string
	b, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	d := string(b)

	// split string by two newlines
	elfParts := strings.Split(d, "\n\n")

	// track totals for all elves
	totals := make([]int, len(elfParts))
	maxCals := 0
	mostCalElf := 0
	for i, elfPart := range elfParts {
		// split each part into lines
		lines := strings.Split(elfPart, "\n")
		elfCals := 0
		for _, line := range lines {
			// skip empty lines (which should just be the last line
			// my IDE always adds a newline to the end of the file
			if line == "" {
				continue
			}
			// parse each line into an int
			cals, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			elfCals += cals
		}
		// add total for this elf
		totals[i] = elfCals

		// save max elf and their cals
		if elfCals > maxCals {
			maxCals = elfCals
			mostCalElf = i
		}

	}
	fmt.Printf("\n\nElf \"%d\" has the most calories with %d", mostCalElf, maxCals)

	// sort totals
	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})
	// sum top 3
	sum := 0
	for i := 0; i < 3; i++ {
		sum += totals[i]
	}
	fmt.Printf("\n\nThe sum of the top 3 elves is %d", sum)
}
