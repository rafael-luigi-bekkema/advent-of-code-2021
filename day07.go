package main

import (
	"fmt"
	"os"
	"strings"
)

func day7a(input []int) int {
	min, max := minmax(input)
	var mincost int
	for pos := min; pos <= max; pos++ {
		var cost int
		for _, crab := range input {
			cost += abs(pos - crab)
		}
		if pos == min || cost < mincost {
			mincost = cost
		}
	}
	return mincost
}

func day7input() []int {
	data, err := os.ReadFile("input/day07.txt")
	if err != nil {
		panic(err)
	}
	var crabs []int
	for _, snum := range strings.Split(strings.TrimRight(string(data), "\n"), ",") {
		crabs = append(crabs, atoi(snum))
	}
	return crabs
}

func Day7a() {
	result := day7a(day7input())
	fmt.Printf("day 7a: %d\n", result)
}

func day7b(input []int) int {
	calcFuel := func(moves int) int {
		var fuel int
		for i := 1; i <= moves; i++ {
			fuel += i
		}
		return fuel
	}
	min, max := minmax(input)
	var mincost int
	for pos := min; pos <= max; pos++ {
		var cost int
		for _, crab := range input {
			cost += calcFuel(abs(pos - crab))
		}
		if pos == min || cost < mincost {
			mincost = cost
		}
	}
	return mincost
}

func Day7b() {
	result := day7b(day7input())
	fmt.Printf("day 7b: %d\n", result)
}
