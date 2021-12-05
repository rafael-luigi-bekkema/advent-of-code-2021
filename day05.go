package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type point struct {
	x, y int
}

func day5a(input io.Reader) int {
	grid := make(map[point]int)
	s := bufio.NewScanner(input)
	var count int
	for s.Scan() {
		var fromx, fromy, tox, toy int
		line := s.Text()
		fmt.Sscanf(line, "%d,%d -> %d,%d", &fromx, &fromy, &tox, &toy)
		if fromx != tox && fromy != toy {
			continue
		}

		if fromx > tox {
			fromx, tox = tox, fromx
		}
		if fromy > toy {
			fromy, toy = toy, fromy
		}

		for x := fromx; x <= tox; x++ {
			for y := fromy; y <= toy; y++ {
				p := point{x, y}
				grid[p]++
				if grid[p] == 2 {
					count++
				}
			}
		}
	}
	return count
}

func Day5a() {
	f, err := os.Open("input/day05.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day5a(f)
	fmt.Printf("day 5a: %d\n", result)
}

func day5b(input io.Reader) int {
	grid := make(map[point]int)
	s := bufio.NewScanner(input)
	var count int
	for s.Scan() {
		var fromx, fromy, tox, toy int
		line := s.Text()
		fmt.Sscanf(line, "%d,%d -> %d,%d", &fromx, &fromy, &tox, &toy)

		var xdelta, ydelta int
		if fromx < tox {
			xdelta = 1
		} else if fromx > tox {
			xdelta = -1
		}
		if fromy < toy {
			ydelta = 1
		} else if fromy > toy {
			ydelta = -1
		}

		x, y := fromx, fromy
		for {
			p := point{x, y}
			grid[p]++
			if grid[p] == 2 {
				count++
			}

			if x == tox && y == toy {
				break
			}

			x += xdelta
			y += ydelta
		}
	}
	return count
}

func Day5b() {
	f, err := os.Open("input/day05.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day5b(f)
	fmt.Printf("day 5b: %d\n", result)
}
