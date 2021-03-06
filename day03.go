package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func day3a(input io.Reader) int {
	s := bufio.NewScanner(input)

	var commons []int
	for s.Scan() {
		line := s.Text()
		if commons == nil {
			commons = make([]int, len(line))
		}
		for i, c := range line {
			switch c {
			case '1':
				commons[i]++
			case '0':
				commons[i]--
			default:
				panic(fmt.Sprintf("unexpected char: %v", c))
			}
		}
	}
	bgamma := make([]rune, len(commons))
	bepsilon := make([]rune, len(commons))
	for i, c := range commons {
		if c > 0 {
			bgamma[i] = '1'
			bepsilon[i] = '0'
		} else {
			bgamma[i] = '0'
			bepsilon[i] = '1'
		}
	}
	gamma, err := strconv.ParseUint(string(bgamma), 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseUint(string(bepsilon), 2, 64)
	if err != nil {
		panic(err)
	}

	return int(gamma * epsilon)
}

func Day3a() {
	f, err := os.Open("input/day03.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day3a(f)
	fmt.Printf("day 3a: %d\n", result)
}

func day3commons(input []string, idx int) int {
	var common int
	for _, line := range input {
		switch c := line[idx]; c {
		case '1':
			common++
		case '0':
			common--
		default:
			panic(fmt.Sprintf("unexpected char: %v", c))
		}
	}
	return common
}

func day3b(input io.Reader) int {
	o2 := scanLines(input)
	co2 := make([]string, len(o2))
	copy(co2, o2)

	for i := 0; len(o2) > 1; i++ {
		common := day3commons(o2, i)
		var newo2 []string
		for _, line := range o2 {
			comp := byte('0')
			if common >= 0 {
				comp = '1'
			}
			if line[i] == comp {
				newo2 = append(newo2, line)
			}
		}
		o2 = newo2
	}

	for i := 0; len(co2) > 1; i++ {
		common := day3commons(co2, i)
		var newco2 []string
		for _, line := range co2 {
			comp := byte('1')
			if common >= 0 {
				comp = '0'
			}
			if line[i] == comp {
				newco2 = append(newco2, line)
			}
		}
		co2 = newco2
	}

	o2gen, err := strconv.ParseUint(o2[0], 2, 64)
	if err != nil {
		panic(err)
	}
	co2gen, err := strconv.ParseUint(co2[0], 2, 64)
	if err != nil {
		panic(err)
	}

	return int(o2gen * co2gen)
}

func Day3b() {
	f, err := os.Open("input/day03.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day3b(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("day 3b: %d\n", result)
}
