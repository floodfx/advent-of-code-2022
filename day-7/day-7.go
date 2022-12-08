package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/floodfx/advent-of-code-2022/in"
	"golang.org/x/exp/slices"
)

const (
	Dir = iota
	File
)

type entry interface {
	Type() int
	Size() int
}

type file struct {
	size int
	name string
}

func (f file) Size() int {
	return f.size
}

func (f file) Type() int {
	return File
}

type dir struct {
	name    string
	parent  *dir
	entries map[string]entry
}

func (d dir) Size() int {
	// size of files in dir + size of dirs
	size := 0
	for _, v := range d.entries {
		size += v.Size()
	}
	return size
}

func (d dir) Type() int {
	return Dir
}

func part1(lines []string) *dir {
	pwd := &dir{
		entries: make(map[string]entry),
	}
	root := pwd
	for _, line := range lines {
		if line == "" {
			continue
		}
		switch line[0] {
		case '$':
			//processCmd
			cmd := strings.Split(line, " ")
			if cmd[1] == "cd" {
				name := cmd[2]
				if name == ".." {
					// pop out of
					pwd = pwd.parent
					continue
				}
				if name == "/" {
					pwd.name = name
					continue
				}
				// go into
				nd := pwd.entries[name].(dir)
				pwd = &nd
			}
			// ls - ignore for now
		case 'd':
			// push dir to entries
			name := strings.Split(line, " ")[1]
			newDir := dir{name: name, parent: pwd, entries: make(map[string]entry)}
			pwd.entries[name] = newDir
		default:
			// push file to entries
			ps := strings.Split(line, " ")
			name := ps[1]
			size, _ := strconv.Atoi(ps[0])
			f := file{name: name, size: size}
			pwd.entries[ps[1]] = f
		}
	}
	return root
}

type found struct {
	dirs []dir
}

func walkFS(e entry, maxSize int, f *found) {
	switch e.Type() {
	case Dir:
		d := e.(dir)
		if maxSize == -1 || d.Size() <= maxSize {
			f.dirs = append(f.dirs, d)
		}
		for _, e := range d.entries {
			walkFS(e, maxSize, f)
		}
	}
}

func printFS(e entry, depth int) {
	pad := strings.Repeat("\t", depth)
	switch e.Type() {
	case Dir:
		d := e.(dir)
		fmt.Printf("%s- %s (dir, size=%d)\n", pad, d.name, d.Size())
		for _, e := range d.entries {
			printFS(e, depth+1)
		}
	case File:
		f := e.(file)
		fmt.Printf("%s- %s (file, size=%d)\n", pad, f.name, f.Size())
	}
}

func main() {
	lines, err := in.Lines("./day-7/input.txt")
	if err != nil {
		panic(err)
	}

	root := part1(lines)
	// printFS(*root, 0)
	f := found{dirs: make([]dir, 0)}
	walkFS(*root, 100000, &f)

	sum := 0
	for _, e := range f.dirs {
		sum += e.Size()
	}
	fmt.Printf("\n\nPart 1 - Found sum: %d\n", sum)

	root = part1(lines)
	totalFS := 70000000
	need := 30000000
	walkFS(*root, -1, &f)
	// sort all the dirs by size
	slices.SortFunc(f.dirs, func(i, j dir) bool {
		return i.Size() < j.Size()
	})

	unused := totalFS - root.Size()
	missing := need - unused
	for _, e := range f.dirs {
		if e.Size() > missing {
			fmt.Printf("Part 2 - %s, (dir, size=%d)", e.name, e.Size())
			break
		}
		sum += e.Size()
	}

}
