package main

import (
	"bytes"
	"fmt"
	"os"
)

func day11a(input []byte, steps int) int {
	width := bytes.Index(input, []byte{'\n'})
	grid := make([]int, 0, len(input))
	for _, v := range input {
		if v == '\n' {
			continue
		}
		grid = append(grid, int(v-'0'))
	}
	height := len(grid) / width

	var total int
	var flash func(i int)
	flash = func(i int) {
		total++
		row := i / width
		col := i % width
		for j := 0; j < 9; j++ {
			if j == 4 {
				continue
			}
			y := row + j/3 - 1
			x := col + j%3 - 1
			if x < 0 || x > width-1 || y < 0 || y > height-1 {
				continue
			}
			newi := y*width + x
			if grid[newi] < 10 && grid[newi] != 0 {
				grid[newi]++
			}
		}
	}
	for step := 0; step < steps; step++ {
		for i := range grid {
			grid[i]++
		}
		for {
			var pops int
			for i, energy := range grid {
				if energy == 10 {
					flash(i)
					grid[i] = 0
					pops++
					continue
				}
			}
			if pops == 0 {
				break
			}
		}
	}
	return total
}

func Day11a() {
	data, err := os.ReadFile("input/day11.txt")
	if err != nil {
		panic(err)
	}
	result := day11a(data, 100)
	fmt.Printf("day 11a: %d\n", result)
}

func day11b(input []byte) int {
	width := bytes.Index(input, []byte{'\n'})
	grid := make([]int, 0, len(input))
	for _, v := range input {
		if v == '\n' {
			continue
		}
		grid = append(grid, int(v-'0'))
	}
	height := len(grid) / width

	var total int
	var flash func(i int)
	flash = func(i int) {
		total++
		row := i / width
		col := i % width
		for j := 0; j < 9; j++ {
			if j == 4 {
				continue
			}
			y := row + j/3 - 1
			x := col + j%3 - 1
			if x < 0 || x > width-1 || y < 0 || y > height-1 {
				continue
			}
			newi := y*width + x
			if grid[newi] < 10 && grid[newi] != 0 {
				grid[newi]++
			}
		}
	}
	var step int
	for {
		allflash := true
		for i, v := range grid {
			if v != 0 {
				allflash = false
			}
			grid[i]++
		}
		if step > 10_000 {
			break
		}
		if allflash {
			return step
		}
		for {
			var pops int
			for i, energy := range grid {
				if energy == 10 {
					flash(i)
					grid[i] = 0
					pops++
					continue
				}
			}
			if pops == 0 {
				break
			}
		}

		step++
	}
	panic("solution not found")
}

func Day11b() {
	data, err := os.ReadFile("input/day11.txt")
	if err != nil {
		panic(err)
	}
	result := day11b(data)
	fmt.Printf("day 11b: %d\n", result)
}
