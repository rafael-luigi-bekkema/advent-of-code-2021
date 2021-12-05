package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func day2a(input io.Reader) (int, error) {
	s := bufio.NewScanner(input)

	var horizontal, depth int
	for s.Scan() {
		var direction string
		var amount int
		_, err := fmt.Sscanf(s.Text(), "%s %d", &direction, &amount)
		if err != nil {
			return 0, fmt.Errorf("could not scan line: %w", err)
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
	return horizontal * depth, nil
}

func Day2a() error {
	f, err := os.Open("input/day02.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	num, err := day2a(f)
	if err != nil {
		return err
	}
	fmt.Printf("day 2a: %d\n", num)
	return nil
}

func day2b(input io.Reader) (int, error) {
	s := bufio.NewScanner(input)

	var horizontal, depth, aim int
	for s.Scan() {
		var direction string
		var amount int
		_, err := fmt.Sscanf(s.Text(), "%s %d", &direction, &amount)
		if err != nil {
			return 0, fmt.Errorf("could not scan line: %w", err)
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
	return horizontal * depth, nil
}

func Day2b() error {
	f, err := os.Open("input/day02.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	num, err := day2b(f)
	if err != nil {
		return err
	}
	fmt.Printf("day 2b: %d\n", num)
	return nil
}
