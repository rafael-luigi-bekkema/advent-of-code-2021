package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func atoi(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}

type field struct {
	num    int
	marked bool
}
type board struct {
	fields [25]field
	won    bool
}

func day4prepare(input io.Reader) (boards []board, numbers []int) {
	s := bufio.NewScanner(input)
	s.Scan()
	for _, s := range strings.Split(s.Text(), ",") {
		numbers = append(numbers, atoi(s))
	}
	var i int
	for s.Scan() {
		line := s.Text()
		if line == "" {
			boards = append(boards, board{[25]field{}, false})
			i = 0
			continue
		}
		for _, s := range strings.Fields(line) {
			boards[len(boards)-1].fields[i] = field{atoi(s), false}
			i++
		}
	}
	return boards, numbers
}

func day4score(b board) int {
	var total int
	var cols, rows [5]int
	var win bool
	for i, f := range b.fields {
		col := i % 5
		row := i / 5
		if f.marked {
			cols[col]++
			rows[row]++
			if cols[col] == 5 || rows[row] == 5 {
				win = true
			}
			continue
		}
		total += f.num

	}
	if win {
		return total
	}
	return 0
}

func day4a(input io.Reader) int {
	boards, draws := day4prepare(input)

	for _, draw := range draws {
		for i, board := range boards {
			for j, field := range board.fields {
				if field.num == draw {
					boards[i].fields[j].marked = true
					if s := day4score(boards[i]); s > 0 {
						return s * draw
					}
				}
			}
		}
	}
	return 0
}

func Day4a() {
	f, err := os.Open("input/day04.txt")
	if err != nil {
		panic(err)
	}
	result := day4a(f)

	fmt.Printf("day 4a: %d\n", result)
}

func day4b(input io.Reader) int {
	boards, draws := day4prepare(input)

	var lastWin int
	for _, draw := range draws {
		for i, board := range boards {
			if board.won {
				continue
			}
			for j, field := range board.fields {
				if field.num == draw {
					boards[i].fields[j].marked = true
					if s := day4score(boards[i]); s > 0 {
						boards[i].won = true
						lastWin = s * draw
					}
				}
			}
		}
	}
	return lastWin
}

func Day4b() {
	f, err := os.Open("input/day04.txt")
	if err != nil {
		panic(err)
	}
	result := day4b(f)

	fmt.Printf("day 4b: %d\n", result)
}
