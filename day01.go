package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func day1a(input io.Reader) (int, error) {
	s := bufio.NewScanner(input)
	var prev, count int
	for s.Scan() {
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			return 0, err
		}
		if prev != 0 && num > prev {
			count++
		}
		prev = num
	}
	if err := s.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

func Day1a() error {
	f, err := os.Open("input/day01.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	count, err := day1a(f)
	if err != nil {
		return err
	}
	fmt.Printf("day 1a: %d", count)
	return nil
}

func sum(ints []int) int {
	var total int
	for _, i := range ints {
		total += i
	}
	return total
}

func day1b(input io.Reader) (int, error) {
	s := bufio.NewScanner(input)
	var count int
	var win []int
	for s.Scan() {
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			return 0, err
		}
		win = append(win, num)
		if len(win) < 4 {
			continue
		}
		if sum(win[1:4]) > sum(win[0:3]) {
			count++
		}
		win = win[1:4]
	}
	if err := s.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

func Day1b() error {
	f, err := os.Open("input/day01.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	count, err := day1b(f)
	if err != nil {
		return err
	}
	fmt.Printf("day 1b: %d", count)
	return nil
}
