package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func day8a(input io.Reader) int {
	s := bufio.NewScanner(input)
	var count int
	for s.Scan() {
		line := s.Text()
		outputs := strings.Fields(strings.Split(line, "|")[1])
		for _, output := range outputs {
			l := len(output)
			if l == 2 || l == 4 || l == 3 || l == 7 {
				count++
			}
		}
	}
	return count
}

func Day8a() {
	f, err := os.Open("input/day08.txt")
	if err != nil {
		panic(err)
	}
	result := day8a(f)
	fmt.Printf("day 8a: %d\n", result)
}

func day8b(input io.Reader) int {
	type Display [7]byte
	type Segments [7][]byte
	filter := func(v1, v2 []byte) []byte {
		var result []byte
		for _, val1 := range v1 {
			for _, val2 := range v2 {
				if val1 == val2 {
					result = append(result, val1)
					break
				}
			}
		}
		return result
	}
	var clean func(segs Segments, hit int) Segments
	clean = func(segs Segments, hit int) Segments {
	outer:
		for {
			for j, seg := range segs {
				if len(segs[j]) == 1 {
					continue
				}
				if j == hit {
					continue
				}
				for k, val := range seg {
					if val == segs[hit][0] {
						segs[j] = append(segs[j][:k], segs[j][k+1:]...)
						continue
					}
				}
				if len(segs[j]) == 1 {
					segs = clean(segs, j)
					continue outer
				}
			}
			break
		}
		return segs
	}

	displays := map[Display]int{
		{1, 1, 1, 0, 1, 1, 1}: 0,
		{0, 0, 1, 0, 0, 1, 0}: 1,
		{1, 0, 1, 1, 1, 0, 1}: 2,
		{1, 0, 1, 1, 0, 1, 1}: 3,
		{0, 1, 1, 1, 0, 1, 0}: 4,
		{1, 1, 0, 1, 0, 1, 1}: 5,
		{1, 1, 0, 1, 1, 1, 1}: 6,
		{1, 0, 1, 0, 0, 1, 0}: 7,
		{1, 1, 1, 1, 1, 1, 1}: 8,
		{1, 1, 1, 1, 0, 1, 1}: 9,
	}
	/*
	    aaaa    0000
	   b    c  1    2
	   b    c  1    2
	    dddd    3333
	   e    f  4    5
	   e    f  4    5
	    gggg    6666
	*/
	s := bufio.NewScanner(input)
	var count int
	for s.Scan() {
		line := s.Text()
		var segs Segments
		for i := range segs {
			segs[i] = []byte("abcdefg")
		}

		parts := strings.Split(line, " | ")
		for _, digit := range strings.Fields(parts[0]) {
			var numbs []int
			switch len(digit) {
			case 2: // 1
				numbs = []int{2, 5}
			case 3: // 7
				numbs = []int{0, 2, 5}
			case 4: // 4
				numbs = []int{1, 2, 3, 5}
			case 5: // 2 or 3 or 5
				// common: a d g
				numbs = []int{0, 3, 6}
			case 6: // 0 or 6 or 9
				// common: a b f g
				numbs = []int{0, 1, 5, 6}
			case 7: // 8
				continue
			}

			for _, i := range numbs {
				segs[i] = filter(segs[i], []byte(digit))
				if len(segs[i]) == 1 {
					segs = clean(segs, i)
				}
			}
		}
		digMap := make(map[byte]int)
		for i, d := range segs {
			digMap[d[0]] = i
		}
		var number int
		output := strings.Fields(parts[1])
		for i, digit := range output {
			var display Display
			for _, odig := range []byte(digit) {
				display[digMap[odig]] = 1
			}
			number += displays[display] * int(math.Pow(10, float64(len(output)-i-1)))
		}
		// fmt.Println(number)
		count += number
	}
	return count
}

func Day8b() {
	f, err := os.Open("input/day08.txt")
	if err != nil {
		panic(err)
	}
	result := day8b(f)
	fmt.Printf("day 8b: %d\n", result)
}
