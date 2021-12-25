package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

var energyM = map[byte]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

func energyPer(c byte) int {
	if v, ok := energyM[c]; ok {
		return v
	}
	panic("unknown char")
}

func slot(c byte) [2]int {
	switch c {
	case 'A':
		return [2]int{1*11 + 2, 2*11 + 2}
	case 'B':
		return [2]int{1*11 + 4, 2*11 + 4}
	case 'C':
		return [2]int{1*11 + 6, 2*11 + 6}
	case 'D':
		return [2]int{1*11 + 8, 2*11 + 8}
	default:
		panic(fmt.Sprintf("unknown thing: %c", c))
	}
}

func home(grid []byte, i int) bool {
	s := slot(grid[i])
	if s[1] == i {
		return true
	}
	return s[0] == i && grid[s[1]] == grid[i]
}

func targetFree(grid []byte, c byte) (int, bool) {
	s := slot(c)
	if grid[s[0]] == '.' {
		if grid[s[1]] == '.' {
			return s[1], true
		}
		if grid[s[1]] != byte(c) {
			return 0, false
		}
		return s[0], true
	}
	return 0, false
}

func havePath(grid []byte, i, j int, c byte, energy int) (bool, int) {
	if i == j {
		return grid[j] == '.', energy
	}
	rowi := i / 11
	coli := i % 11
	rowj := j / 11
	colj := j % 11
	if rowi > 0 && rowj < rowi { // move up
		if grid[i-11] != '.' {
			return false, 0
		}
		return havePath(grid, i-11, j, c, energy+energyPer(c))
	}
	if coli == colj && rowj > rowi { // j is below
		if grid[i+11] != '.' {
			return false, 0
		}
		return havePath(grid, i+11, j, c, energy+energyPer(c))
	}

	if colj > coli {
		if grid[i+1] != '.' {
			return false, 0
		}
		return havePath(grid, i+1, j, c, energy+energyPer(c))
	} else {
		if grid[i-1] != '.' {
			return false, 0
		}
		return havePath(grid, i-1, j, c, energy+energyPer(c))
	}
}

func copyGrid(g []byte) []byte {
	grid := make([]byte, len(g))
	copy(grid, g)
	return grid
}

func isDone(grid []byte) bool {
	for _, c := range []byte{'A', 'B', 'C', 'D'} {
		if s := slot(c); grid[s[0]] != c || grid[s[1]] != c {
			return false
		}
	}
	return true
}

/*
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########
*/

func move(grid []byte, skips []int, energy int, out chan int, mine *int) {
	// renderGrid(grid)
	// fmt.Println("energy", energy)
	// time.Sleep(time.Second * 5)

	if isDone(grid) {
		if *mine == 0 || energy < *mine {
			out <- energy
			*mine = energy
		}
	}

	if *mine > 0 && energy >= *mine {
		// fmt.Println("skip too high")
		return
	}
	// if *mine <= 13539 {
	// 	renderGrid(grid)
	// 	fmt.Println("energy", energy)
	// }
outer:
	for i := range grid {
		for _, s := range skips {
			if i == s {
				continue outer
			}
		}
		if isletter := 'A' <= grid[i] && grid[i] <= 'D'; !isletter || home(grid, i) {
			continue
		}
		if j, ok := targetFree(grid, grid[i]); ok {
			if ok, e := havePath(grid, i, j, grid[i], 0); ok {
				grid := copyGrid(grid)
				grid[j] = grid[i]
				grid[i] = '.'
				// if *mine <= 10625 {
				// 	fmt.Printf("target %d -> %d: %c %v %v\n", i, j, grid[j], energy+e, *mine)
				// 	renderGrid(grid)
				// }
				move(grid, nil, energy+e, out, mine)
				if *mine <= energy {
					continue outer
				}
				continue outer
			}
		}
		inhall := 0 <= i && i < 11
		if !inhall { // in hallway
			// if someone above, skip for now
			if i > 22 && grid[i-11] != '.' {
				continue
			}
			// if i > 22 {
			// 	grid := copyGrid(grid)
			// 	grid[i-11] = grid[i]
			// 	grid[i] = '.'
			// 	move(grid, skips, energy+energyMap[grid[i]], out, mine)
			// }
			// try all hallway positions, and no move at all
			for _, h := range []int{1, 3, 5, 7, 9, 10, 0} {
				if ok, e := havePath(grid, i, h, grid[i], 0); ok {
					grid := copyGrid(grid)
					grid[h] = grid[i]
					grid[i] = '.'
					// if *mine <= 10625 {
					// 	fmt.Printf("hall %d -> %d: %c %v %v\n", i, h, grid[h], energy+e, *mine)
					// 	renderGrid(grid)
					// }
					move(grid, nil, energy+e, out, mine)
					if *mine <= energy {
						continue outer
					}
				}
			}
			// renderGrid(grid)
			// fmt.Printf("no move %d: %c\n", i, grid[i])
			// skips := append(skips, i)
			// move(grid, skips, energy, out, mine)
		}
	}
}

func renderGrid(grid []byte) {
	width := 11
	fmt.Println(strings.Repeat("#", width+2))
	fmt.Print("#")
	for _, g := range grid[:width] {
		fmt.Printf("%c", g)
	}
	fmt.Print("#")
	for i, g := range grid[width:] {
		if i%width == 0 {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Print("\n ")
		}
		fmt.Printf("%c", g)
	}
	fmt.Print(" \n  ")
	fmt.Println(strings.Repeat("#", width-2))
}

func day23a(input string) int {
	width := 11
	data := strings.Split(input, "\n")
	grid := bytes.Repeat([]byte{'#'}, width*3)
	for i := 0; i < 11; i++ {
		grid[i] = '.'
	}
	for i := 2; i <= 8; i += 2 {
		grid[1*width+i] = data[2][i+1]
		grid[2*width+i] = data[3][i+1]
	}
	renderGrid(grid)
	out := make(chan int)
	go func() {
		defer close(out)
		var mine int
		move(grid, nil, 0, out, &mine)
		fmt.Println("done!")
	}()
	var minval int
	for val := range out {
		if minval == 0 || val < minval {
			minval = val
			fmt.Println(val)
		}
	}
	return minval
}

func day23file() string {
	data, err := os.ReadFile("input/day23.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func Day23a() {
	result := day23a(day23file())
	fmt.Println("day 23a:", result)
}
