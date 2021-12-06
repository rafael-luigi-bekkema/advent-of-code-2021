package main

import (
	"fmt"
	"os"
	"strings"
)

func day6a(fish []int, days int) int {
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

func Day6a() {
	data, err := os.ReadFile("input/day06.txt")
	if err != nil {
		panic(err)
	}
	var fish []int
	for _, snum := range strings.Split(strings.TrimRight(string(data), "\n"), ",") {
		fish = append(fish, atoi(snum))
	}
	days := 80
	result := day6a(fish, days)
	fmt.Printf("day 6a: %d\n", result)
}
