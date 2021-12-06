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

func scanLines(input io.Reader) []string {
	var lines []string
	s := bufio.NewScanner(input)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}
