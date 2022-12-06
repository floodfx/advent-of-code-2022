package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/floodfx/advent-of-code-2022/in"
)

var d = regexp.MustCompile(`\d+`)

type Move struct {
	Count int
	From  int
	To    int
}

type Stacks struct {
	s []string
}

func sampleStack() []string {
	return []string{
		"", // index offset
		"ZN",
		"MCD",
		"P",
	}
}

func inputStacks() []string {
	return []string{
		"", // index offset
		"GDVZJSB",
		"ZSMGVP",
		"CLBSWTQF",
		"HJGWMRVQ",
		"CLSNFMD",
		"RGCD",
		"HGTRJDSQ",
		"PFV",
		"DRSTJ",
	}
}

func apply9000Move(m *Move, stacks *Stacks) {
	dest := stacks.s[m.To]
	from := stacks.s[m.From]
	for i := 0; i < m.Count; i++ {
		fromEnd := len(from) - 1
		crate := from[fromEnd]
		from = from[:fromEnd]
		dest = dest + string(crate)
	}
	stacks.s[m.To] = dest
	stacks.s[m.From] = from
}

func apply9001Move(m *Move, stacks *Stacks) {
	fl := len(stacks.s[m.From])
	crates := ""
	if m.Count > 1 {
		fromStart := fl - m.Count
		crates = stacks.s[m.From][fromStart:]
		stacks.s[m.From] = stacks.s[m.From][:fromStart]
	} else {
		fromEnd := fl - 1
		crates = string(stacks.s[m.From][fromEnd])
		stacks.s[m.From] = stacks.s[m.From][:fromEnd]
	}
	stacks.s[m.To] += crates
}

func extractMove(line string) *Move {
	if line == "" {
		return nil
	}
	ds := d.FindAllString(line, -1)
	c, _ := strconv.Atoi(ds[0])
	f, _ := strconv.Atoi(ds[1])
	t, _ := strconv.Atoi(ds[2])

	m := &Move{
		Count: c,
		From:  f,
		To:    t,
	}
	return m
}

func part1(lines []string, stacks *Stacks) string {
	for _, line := range lines {
		m := extractMove(line)
		if m == nil {
			continue
		}
		apply9000Move(m, stacks)
	}
	msg := ""
	for i := 1; i < len(stacks.s); i++ {
		s := stacks.s[i]
		msg += string(s[len(s)-1])
	}
	return msg
}

func part2(lines []string, stacks *Stacks) string {

	for _, line := range lines {
		m := extractMove(line)
		if m == nil {
			continue
		}
		apply9001Move(m, stacks)
	}
	msg := ""
	for i := 1; i < len(stacks.s); i++ {
		s := stacks.s[i]
		msg += string(s[len(s)-1])
	}
	return msg
}

func main() {
	lines, err := in.Lines("./day-5/input.txt")
	if err != nil {
		panic(err)
	}

	stks := &Stacks{
		s: inputStacks(),
	}
	msg := part1(lines, stks)
	fmt.Printf("\n\nPart 1 - Msg is: %s, stacks:%v\n", msg, stks)

	stks2 := &Stacks{
		s: inputStacks(),
	}
	msg = part2(lines, stks2)
	fmt.Printf("\n\nPart 2 - Msg is: %s, stacks:%v\n", msg, stks2)
}
