package main

import (
	"fmt"
	"os"
	"strings"
)

func day6(fish []int, days int) int {
	var timers [9]int

	for _, timer := range fish {
		timers[timer]++
	}

	for day := 1; day <= days; day++ {
		first := timers[0]
		copy(timers[:], timers[1:])
		timers[6] += first
		timers[8] = first
	}
	return sum(timers[:])
}

func day6wFile(days int) int {
	data, err := os.ReadFile("input/day06.txt")
	if err != nil {
		panic(err)
	}
	var fish []int
	for _, snum := range strings.Split(strings.TrimRight(string(data), "\n"), ",") {
		fish = append(fish, atoi(snum))
	}
	result := day6(fish, days)
	return result
}

func Day6a() {
	result := day6wFile(80)
	fmt.Printf("day 6a: %d\n", result)
}

func Day6b() {
	result := day6wFile(256)
	fmt.Printf("day 6b: %d\n", result)
}
