package main

import (
	"fmt"
	"os"
	"strings"
)

func day25file() string {
	data, err := os.ReadFile("input/day25.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func day25a(input string) int {
	width := strings.Index(input, "\n")
	grid := []byte(strings.ReplaceAll(input, "\n", ""))
	height := len(grid) / width
	step := 0
	render := func() {
		for i := range grid {
			if i%width == 0 {
				fmt.Println()
				if i == 0 {
					fmt.Println("step", step)
				}
			}
			fmt.Printf("%c", grid[i])
		}
		fmt.Println()
	}
	_ = render
	for {
		step++
		newgrid := make([]byte, len(grid))
		copy(newgrid, grid)
		var moves int
		for i, v := range grid {
			if v != '>' {
				continue
			}
			row := i / width
			col := i%width + 1
			if col >= width {
				col = 0
			}
			if idx := row*width + col; grid[idx] == '.' {
				newgrid[i] = '.'
				newgrid[idx] = '>'
				moves++
			}
		}

		copy(grid, newgrid)
		for i, v := range grid {
			if v != 'v' {
				continue
			}
			row := i/width + 1
			if row >= height {
				row = 0
			}
			col := i % width
			if idx := row*width + col; grid[idx] == '.' {
				newgrid[i] = '.'
				newgrid[idx] = 'v'
				moves++
			}
		}
		grid = newgrid
		// render()
		if moves == 0 {
			break
		}
	}
	return step
}

func Day25a() {
	result := day25a(day25file())
	fmt.Println("day 25a:", result)
}


func Day25() {
	result := day25a(day25file())
	fmt.Println("day 25:", result)
}
