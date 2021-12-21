package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type player struct {
	score    int
	position int
}

func day21a(startPos1, startPos2 int) int {
	var last, throwCount int
	throw := func() int {
		throwCount += 3
		last++
		if last > 100 {
			last = 1
		}
		v1 := last
		last++
		if last > 100 {
			last = 1
		}
		v2 := last
		last++
		if last > 100 {
			last = 1
		}
		return v1 + v2 + last
	}
	players := [2]player{
		{score: 0, position: startPos1 - 1},
		{score: 0, position: startPos2 - 1},
	}
	for {
		for i := range players {
			advance := throw()
			players[i].position = (players[i].position + advance) % 10
			players[i].score += players[i].position + 1
			// fmt.Printf("Player %d rolls %d and moves to %d for a score of %d\n", i+1, advance, players[i].position+1, players[i].score)
			if players[i].score >= 1000 {
				return players[(i+1)%2].score * throwCount
			}
		}
	}
}

func Day21a() {
	data, err := os.ReadFile("input/day21.txt")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(data), "\n")

	var player, pos1, pos2 int
	fmt.Sscanf(parts[0], "Player %d starting position: %d", &player, &pos1)
	fmt.Sscanf(parts[1], "Player %d starting position: %d", &player, &pos2)
	result := day21a(pos1, pos2)
	fmt.Println("day 21a:", result)
}

type universe struct {
	players [2]player
	last    int
}

func day21b(startPos1, startPos2 int) int {
	var last, throwCount int
	universeCount := 1
	throw := func() int {
		throwCount += 3
		rand.Seed(time.Now().UnixNano())
		last = rand.Intn(4)
		universeCount *= 27
		return last
	}
outer:
	for i := 0; i < 100; i++ {
		players := [2]player{
			{score: 0, position: startPos1 - 1},
			{score: 0, position: startPos2 - 1},
		}
		throwCount = 0
		universeCount = 1
		for {
			for i := range players {
				advance := throw()
				players[i].position = (players[i].position + advance) % 10
				players[i].score += players[i].position + 1
				if players[i].score >= 21 {
					fmt.Printf("Player %d wins with %d throws, score: %d\n", i+1, throwCount, players[i].score)
					fmt.Println(universeCount)
					continue outer

					// return players[(i+1)%2].score * throwCount
				}
			}
		}
	}
	return 0
}
