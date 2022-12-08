package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/floodfx/advent-of-code-2022/in"
)

type tree struct {
	id     int
	row    int
	col    int
	height int
}

func (t *tree) scenicScore(f forrest) int {
	// edges always get 0
	if t.row == 0 || t.row == f.rows-1 {
		return 0
	}
	if t.col == 0 || t.col == f.cols-1 {
		return 0
	}
	// interior tree
	row := f.row(*t)
	col := f.col(*t)
	leftScore := 0
	// left viz
	for i := t.col - 1; i >= 0; i-- {
		leftScore += 1
		lt := row[i]
		if lt.height >= t.height {
			break
		}
	}
	// right viz
	rightViz := 0
	for i := t.col + 1; i < f.cols; i++ {
		rightViz += 1
		rt := row[i]
		if rt.height >= t.height {
			break
		}
	}
	// up viz
	upViz := 0
	for i := t.row - 1; i >= 0; i-- {
		upViz += 1
		ut := col[i]
		if ut.height >= t.height {
			break
		}
	}
	// down viz
	downViz := 0
	for i := t.row + 1; i < f.rows; i++ {
		downViz += 1
		dt := col[i]
		if dt.height >= t.height {
			break
		}
	}
	return leftScore * rightViz * upViz * downViz
}

func (t *tree) visible(f forrest) bool {
	// if first or last row then visible
	if t.row == 0 || t.row == f.rows-1 {
		return true
	}
	// if first or last col then visible
	if t.col == 0 || t.col == f.cols-1 {
		return true
	}
	// interior tree
	row := f.row(*t)
	col := f.col(*t)
	leftViz := true
	// left viz
	for i := t.col - 1; i >= 0; i-- {
		lt := row[i]
		if lt.height >= t.height {
			leftViz = false
			break
		}
	}
	if leftViz {
		return true
	}
	// right viz
	rightViz := true
	for i := t.col + 1; i < f.cols; i++ {
		rt := row[i]
		if rt.height >= t.height {
			rightViz = false
			break
		}
	}
	if rightViz {
		return true
	}
	// up viz
	upViz := true
	for i := t.row - 1; i >= 0; i-- {
		ut := col[i]
		if ut.height >= t.height {
			upViz = false
			break
		}
	}
	if upViz {
		return true
	}
	// down viz
	downViz := true
	for i := t.row + 1; i < f.rows; i++ {
		dt := col[i]
		if dt.height >= t.height {
			downViz = false
			break
		}
	}
	return downViz
}

type forrest struct {
	trees []tree
	cols  int
	rows  int
}

func (f *forrest) row(tr tree) []tree {
	ind := tr.row * f.cols
	return f.trees[ind : ind+f.cols]
}
func (f *forrest) col(tr tree) []tree {
	col := make([]tree, 0)
	for i := 0; i < f.rows; i++ {
		ind := i*f.cols + tr.col
		col = append(col, f.trees[ind])
	}
	return col
}
func (f *forrest) visible() int {
	vizCount := 0
	for _, tr := range f.trees {
		if tr.visible(*f) {
			vizCount++
		}
	}
	return vizCount
}

func buildForrest(lines []string) forrest {
	f := forrest{
		trees: make([]tree, 0),
		rows:  len(lines) - 1,
	}
	for n, line := range lines {
		if line == "" {
			continue
		}
		ts := strings.Split(line, "")
		cols := len(ts)
		f.cols = cols
		for i, t := range ts {
			h, _ := strconv.Atoi(t)
			tr := &tree{
				id:     i + (n * cols),
				height: h,
				row:    n,
				col:    i,
			}
			f.trees = append(f.trees, *tr)
		}
	}
	return f
}

func main() {
	lines, err := in.Lines("./day-8/input.txt")
	if err != nil {
		panic(err)
	}

	f := buildForrest(lines)
	fmt.Printf("\nPart 1 - Visible Trees: %d", f.visible())

	maxScenic := 0
	for _, tr := range f.trees {
		ss := tr.scenicScore(f)
		if ss > maxScenic {
			maxScenic = ss
		}
	}
	fmt.Printf("\nPart 2 - Max Scenic Score: %d", maxScenic)
}
