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

func Day14a() {
	f, err := os.Open("input/day14.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day14b(day14parseInput(f), 10)
	fmt.Printf("day 14a: %d\n", result)
}

func day14b(input day14input, steps int) int {
	type pair [2]byte

	pairs := map[pair]int{}
	for i := 0; i < len(input.template)-1; i++ {
		p := pair{input.template[i], input.template[i+1]}
		pairs[p]++
	}
	ccount := map[byte]int{}
	for _, c := range input.template {
		ccount[c]++
	}

	// We don't have to build the complete string.
	// Just keep track of the pair counts and count new chars we make.
	for step := 1; step <= steps; step++ {
		newpairs := map[pair]int{}
		for p, count := range pairs {
			b := input.rules[p]
			p1 := pair{p[0], b}
			p2 := pair{b, p[1]}

			ccount[b] += count

			newpairs[p1] += count
			newpairs[p2] += count
		}
		pairs = newpairs
	}

	var imin, imax int
	for _, count := range ccount {
		if imin == 0 || count < imin {
			imin = count
		}
		if count > imax {
			imax = count
		}
	}
	return imax - imin
}

func Day14b() {
	f, err := os.Open("input/day14.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day14b(day14parseInput(f), 40)
	fmt.Printf("day 14b: %d\n", result)
}
