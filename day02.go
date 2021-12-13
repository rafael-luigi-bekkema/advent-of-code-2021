package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func day2a(input io.Reader) int {
	s := bufio.NewScanner(input)

	var horizontal, depth int
	for s.Scan() {
		var direction string
		var amount int
		_, err := fmt.Sscanf(s.Text(), "%s %d", &direction, &amount)
		if err != nil {
			panic(err)
		}
		switch direction {
		case "down":
			depth += amount
		case "up":
			depth -= amount
		case "forward":
			horizontal += amount
		}
	}
	return horizontal * depth
}

func Day2a() {
	f, err := os.Open("input/day02.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	num := day2a(f)
	fmt.Printf("day 2a: %d\n", num)
}

func day2b(input io.Reader) int {
	s := bufio.NewScanner(input)

	var horizontal, depth, aim int
	for s.Scan() {
		var direction string
		var amount int
		_, err := fmt.Sscanf(s.Text(), "%s %d", &direction, &amount)
		if err != nil {
			panic(err)
		}
		switch direction {
		case "down":
			aim += amount
		case "up":
			aim -= amount
		case "forward":
			horizontal += amount
			depth += aim * amount
		}
	}
	return horizontal * depth
}

func Day2b() {
	f, err := os.Open("input/day02.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	num := day2b(f)
	fmt.Printf("day 2b: %d\n", num)
}
