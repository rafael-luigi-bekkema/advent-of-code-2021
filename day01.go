package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func day1a(input io.Reader) int {
	s := bufio.NewScanner(input)
	var prev, count int
	for s.Scan() {
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
		if prev != 0 && num > prev {
			count++
		}
		prev = num
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return count
}

func Day1a() {
	f, err := os.Open("input/day01.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	count := day1a(f)
	fmt.Printf("day 1a: %d\n", count)
}

func day1b(input io.Reader) int {
	s := bufio.NewScanner(input)
	var count int
	var win []int
	for s.Scan() {
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
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
		panic(err)
	}
	return count
}

func Day1b() {
	f, err := os.Open("input/day01.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	count := day1b(f)
	fmt.Printf("day 1b: %d\n", count)
}
