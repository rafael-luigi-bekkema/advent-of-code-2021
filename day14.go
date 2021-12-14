package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type day14input struct {
	template []byte
	rules    map[[2]byte]byte
}

func day14parseInput(input io.Reader) day14input {
	s := bufio.NewScanner(input)
	var pinp day14input
	pinp.rules = make(map[[2]byte]byte)
	for s.Scan() {
		line := s.Text()
		if len(pinp.template) == 0 {
			pinp.template = []byte(line)
			continue
		}
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")

		pair := [2]byte{parts[0][0], parts[0][1]}
		insert := parts[1][0]

		if _, ok := pinp.rules[pair]; ok {
			panic("double rule: " + line)
		}
		pinp.rules[pair] = insert
	}
	return pinp
}

func day14a(input day14input, steps int) int {
	template := make([]byte, len(input.template), 524_288_000)
	copy(template, input.template)

	for step := 1; step <= steps; step++ {
		for i := 0; i < len(template)-1; i++ {
			var pair [2]byte
			copy(pair[:], template[i:i+2])
			if insert, ok := input.rules[pair]; ok {
				template = append(template, 0)
				copy(template[i+2:], template[i+1:])
				template[i+1] = insert
				i++
			}
		}
	}

	counts := map[byte]int{}
	for _, b := range template {
		counts[b]++
	}
	var minc, maxc int
	for _, c := range counts {
		if maxc == 0 || c > maxc {
			maxc = c
		}
		if minc == 0 || c < minc {
			minc = c
		}
	}
	return maxc - minc
}

func Day14a() {
	f, err := os.Open("input/day14.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day14a(day14parseInput(f), 10)
	fmt.Printf("day 14a: %d\n", result)
}
