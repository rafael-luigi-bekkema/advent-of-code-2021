package main

import (
	"fmt"
	"os"
	"strings"
)

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
			players[i].move(advance)
			// fmt.Printf("Player %d rolls %d and moves to %d for a score of %d\n", i+1, advance, players[i].position+1, players[i].score)
			if players[i].score >= 1000 {
				return players[(i+1)%2].score * throwCount
			}
		}
	}
}

func day21file() (int, int) {
	data, err := os.ReadFile("input/day21.txt")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(data), "\n")

	var player, pos1, pos2 int
	fmt.Sscanf(parts[0], "Player %d starting position: %d", &player, &pos1)
	fmt.Sscanf(parts[1], "Player %d starting position: %d", &player, &pos2)
	return pos1, pos2
}

func Day21a() {
	result := day21a(day21file())
	fmt.Println("day 21a:", result)
}

type player struct {
	score    int
	position int
}

func (p *player) move(n int) {
	p.position = (p.position + n) % 10
	p.score += p.position + 1
}

func (p player) String() string {
	return fmt.Sprintf("pos %d score %d", p.position, p.score)
}

type universe struct {
	players [2]player
	sum     int
	activeP int
}

type wuniverse struct {
	universe
	weight int
}

func (u universe) String() string {
	return fmt.Sprintf("%v, sum: %d, activeP: %d", u.players, u.sum, u.activeP)
}

func (u *universe) roll(n, throws int) bool {
	u.sum += n
	if throws%3 == 0 {
		u.players[u.activeP].move(u.sum)
		if u.players[u.activeP].score >= 21 {
			return true
		}
		// fmt.Printf("player %d roll %d score %d\n", u.activeP+1, u.throws, u.players[u.activeP].score)
		u.sum = 0
		u.activeP = (u.activeP + 1) % 2
	}
	return false
}

func day21b(startPos1, startPos2 int) int {
	universes := []wuniverse{
		{
			universe: universe{
				players: [2]player{{score: 0, position: startPos1 - 1}, {score: 0, position: startPos2 - 1}},
				sum:     0,
				activeP: 0,
			},
			weight: 1,
		},
	}
	var wins [2]int
	throws := 0
	throw := func(i int) {
		for _, n := range []int{2, 3} {
			newu := universes[i]
			if won := newu.roll(n, throws); won {
				wins[newu.activeP] += newu.weight
				continue
			}
			universes = append(universes, newu)
			if len(universes)%10_000_000 == 0 {
				fmt.Println("wins 1", wins[0], "wins 2", wins[1])
				fmt.Println("universes", len(universes))
			}
		}
		u := &universes[i]
		if won := u.roll(1, throws); won {
			wins[u.activeP] += u.weight
			universes[len(universes)-1], universes[i] = universes[i], universes[len(universes)-1]
			universes = universes[:len(universes)-1]
		}
	}

outer:
	for r := 0; r < 50; r++ {
		for tn := 1; tn <= 3; tn++ {
			throws++
			for i := len(universes) - 1; i >= 0; i-- {
				throw(i)
			}
			// collapse universes
			m := make(map[universe]int)
			for _, wu := range universes {
				m[wu.universe] += wu.weight
			}
			universes = universes[:0]
			for u, w := range m {
				universes = append(universes, wuniverse{u, w})
			}
			if len(universes) == 0 {
				break outer
			}
		}
	}
	if wins[0] > wins[1] {
		return wins[0]
	}
	return wins[1]
}

func Day21b() {
	result := day21b(day21file())
	fmt.Println("day 21b:", result)
}
