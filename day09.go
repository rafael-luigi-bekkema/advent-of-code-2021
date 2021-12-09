package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

func day9a(input []byte) int {
	width := bytes.Index(input, []byte{'\n'})
	rows := len(input) / width
	input = bytes.ReplaceAll(input, []byte{'\n'}, nil)

	var total int
	for i, b := range input {
		row := i / width
		col := i % width
		if row > 0 && input[(row-1)*width+col] <= b {
			continue
		}
		if row < rows-1 && input[(row+1)*width+col] <= b {
			continue
		}
		if col > 0 && input[row*width+col-1] <= b {
			continue
		}
		if col < width-1 && input[row*width+col+1] <= b {
			continue
		}
		total += 1 + int(b-'0')
	}
	return total
}

func Day9a() {
	data, err := os.ReadFile("input/day09.txt")
	if err != nil {
		panic(err)
	}
	data = bytes.TrimRight(data, "\n")
	result := day9a(data)
	fmt.Printf("day 9a: %d\n", result)
}

func day9b(input []byte) int {
	width := bytes.Index(input, []byte{'\n'})
	rows := len(input) / width
	input = bytes.ReplaceAll(input, []byte{'\n'}, nil)

	var size func(i int, basin map[int]bool)
	size = func(i int, basin map[int]bool) {
		if input[i] == '9' || basin[i] {
			return
		}
		basin[i] = true
		row := i / width
		col := i % width

		if row > 0 { // Up
			size((row-1)*width+col, basin)
		}
		if row < rows-1 { // Down
			size((row+1)*width+col, basin)
		}
		if col > 0 { // Left
			size(row*width+col-1, basin)
		}
		if col < width-1 { // Right
			size(row*width+col+1, basin)
		}
	}

	// Find low points
	var sizes []int
	for i, b := range input {
		row := i / width
		col := i % width
		if row > 0 && input[(row-1)*width+col] <= b {
			continue
		}
		if row < rows-1 && input[(row+1)*width+col] <= b {
			continue
		}
		if col > 0 && input[row*width+col-1] <= b {
			continue
		}
		if col < width-1 && input[row*width+col+1] <= b {
			continue
		}

		basin := make(map[int]bool)
		size(i, basin)
		sizes = append(sizes, len(basin))
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})
	var total int
	for _, s := range sizes[:3] {
		if total == 0 {
			total = s
			continue
		}
		total *= s
	}
	return total
}

func Day9b() {
	data, err := os.ReadFile("input/day09.txt")
	if err != nil {
		panic(err)
	}
	data = bytes.TrimRight(data, "\n")
	result := day9b(data)
	fmt.Printf("day 9b: %d\n", result)
}
