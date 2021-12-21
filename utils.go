package main

import (
	"bufio"
	"io"
	"strconv"
)

func atoi(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}

func sum(ints []int) (total int) {
	for _, i := range ints {
		total += i
	}
	return total
}

func minmax(ints []int) (min, max int) {
	for i, num := range ints {
		if i == 0 || num < min {
			min = num
		}
		if i == 0 || num > max {
			max = num
		}
	}
	return
}

func scanLines(input io.Reader) []string {
	var lines []string
	s := bufio.NewScanner(input)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
