package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type coord struct {
	x, y int
}

func day13a(input io.Reader, foldLimit int, ascii io.Writer) int {
	s := bufio.NewScanner(input)
	grid := make(map[coord]bool)
	var maxx, maxy int
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break // no more dots
		}
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		grid[coord{x, y}] = true
		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
	}

	// instructions
	var foldCount int
	for s.Scan() {
		foldCount++
		var axis string
		var along int
		fmt.Sscanf(s.Text(), "fold along %1s=%d", &axis, &along)

		if axis == "x" {
			for curx := maxx; curx > along; curx-- {
				for cury := 0; cury <= maxy; cury++ {
					key := coord{curx, cury}
					_, ok := grid[key]
					if !ok {
						continue
					}
					delete(grid, key)
					key.x = along - (key.x - along)
					grid[key] = true
				}
			}
			maxx = along - 1
		}
		if axis == "y" {
			for cury := maxy; cury > along; cury-- {
				for curx := 0; curx <= maxx; curx++ {
					key := coord{curx, cury}
					_, ok := grid[key]
					if !ok {
						continue
					}
					delete(grid, key)
					key.y = along - (key.y - along)
					grid[key] = true
				}
			}
			maxy = along - 1
		}
		if foldLimit > 0 && foldCount >= foldLimit {
			break
		}
	}
	if ascii != nil {
		for cury := 0; cury <= maxy; cury++ {
			for curx := 0; curx <= maxx; curx++ {
				if grid[coord{curx, cury}] {
					fmt.Fprint(ascii, "#")
				} else {
					fmt.Fprint(ascii, ".")
				}
			}
			fmt.Fprintln(ascii)
		}
	}
	return len(grid)
}

func Day13a() {
	f, err := os.Open("input/day13.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day13a(f, 1, nil)
	fmt.Printf("day 13a: %d\n", result)
}

func Day13b() {
	f, err := os.Open("input/day13.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var s strings.Builder
	day13a(f, -1, &s)
	fmt.Println("day 13b:")
	fmt.Println(s.String())
}
