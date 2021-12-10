package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func day10a(input io.Reader) int {
	s := bufio.NewScanner(input)
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	values := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	var score int
	var i int
	for s.Scan() {
		i++
		var stack []rune
		line := s.Text()
		for _, c := range line {
			if c == '{' || c == '(' || c == '[' || c == '<' {
				stack = append(stack, c)
				continue
			}
			other, ok := pairs[c]
			if !ok {
				panic(fmt.Sprintf("unexpected char: %c", c))
			}
			if stack[len(stack)-1] == other {
				// matching pair, pop the stack
				stack = stack[:len(stack)-1]
				continue
			}
			score += values[c]
			break
		}
	}
	return score
}

func Day10a() {
	f, err := os.Open("input/day10.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day10a(f)
	fmt.Printf("day 10a: %d\n", result)
}

func day10b(input io.Reader) int {
	s := bufio.NewScanner(input)
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	values := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	var scores []int
	var i int
outer:
	for s.Scan() {
		i++
		var stack []rune
		line := s.Text()
		for _, c := range line {
			if c == '{' || c == '(' || c == '[' || c == '<' {
				stack = append(stack, c)
				continue
			}
			other, ok := pairs[c]
			if !ok {
				panic(fmt.Sprintf("unexpected char: %c", c))
			}
			if stack[len(stack)-1] == other {
				// matching pair, pop the stack
				stack = stack[:len(stack)-1]
				continue
			}
			// corrupt line
			continue outer
		}
		if len(stack) > 0 { // incomplete line
			var score int
			for j := len(stack) - 1; j >= 0; j-- {
				points, ok := values[stack[j]]
				if !ok {
					panic(fmt.Sprintf("unexpected char: %c", stack[j]))
				}
				score = score*5 + points
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func Day10b() {
	f, err := os.Open("input/day10.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day10b(f)
	fmt.Printf("day 10b: %d\n", result)
}
