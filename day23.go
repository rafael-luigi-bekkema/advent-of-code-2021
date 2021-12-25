package main

import (
	"fmt"
	"os"
	"strings"
)

func energyPer(p *Point) int {
	switch p.letter {
	case 'A':
		return 1
	case 'B':
		return 10
	case 'C':
		return 100
	case 'D':
		return 1000
	default:
		panic("unknown char")
	}
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

func (grid *Grid) targetFree(point *Point) (int, bool) {
	for i := grid.height - 1; i >= 1; i-- {
		idx := i*grid.width + point.homeCol
		tp := grid.data[idx]
		if tp == nil {
			return idx, true
		}
		if tp.letter != point.letter {
			return 0, false
		}
	}
	return 0, false
}

func (grid *Grid) havePath(i, j int, p *Point, energy int) (bool, int) {
	if i == j {
		return grid.data[j] == nil, energy
	}
	rowi := i / grid.width
	coli := i % grid.width
	rowj := j / grid.width
	colj := j % grid.width
	if rowi > 0 && colj != coli { // move up
		if grid.data[i-grid.width] != nil {
			return false, 0
		}
		return grid.havePath(i-grid.width, j, p, energy+energyPer(p))
	}
	if coli == colj && rowj > rowi { // j is below
		if grid.data[i+grid.width] != nil {
			return false, 0
		}
		return grid.havePath(i+11, j, p, energy+energyPer(p))
	}

	if colj > coli {
		if grid.data[i+1] != nil {
			return false, 0
		}
		return grid.havePath(i+1, j, p, energy+energyPer(p))
	} else {
		if grid.data[i-1] != nil {
			return false, 0
		}
		return grid.havePath(i-1, j, p, energy+energyPer(p))
	}
}

func copyGrid(g []byte) []byte {
	grid := make([]byte, len(g))
	copy(grid, g)
	return grid
}

/*
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########
*/

func (grid Grid) move(skips []int, energy int, out chan int, mine *int) {
	if *mine > 0 && energy >= *mine {
		// fmt.Println("skip too high")
		return
	}

	// grid.render()
	// fmt.Println("energy", energy)
	// time.Sleep(time.Second * 1)

	if grid.isDone() {
		if *mine == 0 || energy < *mine {
			out <- energy
			*mine = energy
		}
	}
	// if *mine <= 13539 {
	// 	renderGrid(grid)
	// 	fmt.Println("energy", energy)
	// }
outer:
	for i, point := range grid.data {
		if point == nil || grid.homes[point.idx] {
			continue
		}
		for _, s := range skips {
			if i == s {
				continue outer
			}
		}
		if j, ok := grid.targetFree(point); ok {
			if ok, e := grid.havePath(i, j, point, 0); ok {
				grid := grid
				grid.data[j] = grid.data[i]
				grid.data[i] = nil
				grid.homes[point.idx] = true
				// fmt.Printf("target %d -> %d: %c %v %v\n", i, j, point.letter, energy+e, *mine)
				grid.move(nil, energy+e, out, mine)
				if *mine > 0 && *mine <= energy {
					continue outer
				}
				continue outer
			}
		}
		inhall := 0 <= i && i < 11
		if !inhall { // in hallway
			// if someone above, skip for now
			if i > 22 && grid.data[i-11] != nil {
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
				if ok, e := grid.havePath(i, h, point, 0); ok {
					grid := grid
					grid.data[h] = grid.data[i]
					grid.data[i] = nil
					// if *mine <= 10625 {
					// 	fmt.Printf("hall %d -> %d: %c %v %v\n", i, h, grid[h], energy+e, *mine)
					// 	renderGrid(grid)
					// }
					grid.move(nil, energy+e, out, mine)
					if *mine > 0 && *mine <= energy {
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

func (grid *Grid) render() {
	fmt.Println(strings.Repeat("#", grid.width+2))
	fmt.Print("#")
	for _, g := range grid.data[:grid.width] {
		var l byte
		if g == nil {
			l = '.'
		} else {
			l = g.letter
		}
		fmt.Printf("%c", l)
	}
	fmt.Print("#")
	for i, g := range grid.data[grid.width : grid.height*grid.width] {
		if i%grid.width == 0 {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Print("\n ")
		}
		var l byte
		if g == nil {
			if i%grid.width%2 == 1 {
				l = '#'
			} else {
				l = '.'
			}
		} else {
			if grid.homes[g.idx] {
				l = g.letter + ('h' - 'H')
			} else {
				l = g.letter
			}
		}
		fmt.Printf("%c", l)
	}
	fmt.Print(" \n  ")
	fmt.Println(strings.Repeat("#", grid.width-2))
}

func day23a(data []string) int {
	grid := day23grid(data)
	// grid.render()
	out := make(chan int)
	go func() {
		defer close(out)
		var mine int
		grid.move(nil, 0, out, &mine)
	}()
	var minval int
	for val := range out {
		if minval == 0 || val < minval {
			minval = val
			// fmt.Println(val)
		}
	}
	return minval
}

type Grid struct {
	width, height int
	data          [55]*Point
	homes         [16]bool
}

type Point struct {
	letter  byte
	homeCol int
	idx     int
}

func (grid *Grid) home(i int, p *Point) bool {
	if i < grid.width {
		return false
	}
	if i%grid.width == p.homeCol {
		for row := grid.height - 1; row > (i / grid.width); row-- {
			if p2 := grid.data[row*grid.width+p.homeCol]; p2 == nil || p2.letter != p.letter {
				return false
			}
		}
		return true
	}
	return false
}

func (grid *Grid) isDone() bool {
	for i := 0; i < (grid.height-1)*4; i++ {
		if !grid.homes[i] {
			return false
		}
	}
	return true
}

func day23grid(data []string) Grid {
	grid := Grid{width: 11, height: len(data) - 2}

	w := grid.width
	var pidx int
	for i := 2; i <= 8; i += 2 {
		for j := 1; j <= grid.height-1; j++ {
			l := data[j+1][i+1]
			hcol := int((l - 'A' + 1) * 2)
			grid.data[j*w+i] = &Point{letter: l, homeCol: hcol, idx: pidx}
			pidx++
		}
	}
	for i, d := range grid.data {
		if d == nil {
			continue
		}
		grid.homes[d.idx] = grid.home(i, d)
	}
	return grid
}

func day23file() []string {
	data, err := os.ReadFile("input/day23.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}

func day23bfile() []string {
	input := day23file()
	input = append(input, "  #D#C#B#A#", "  #D#B#A#C#")
	input[5], input[3] = input[3], input[5]
	input[6], input[4] = input[4], input[6]
	return input
}

func Day23a() {
	result := day23a(day23file())
	fmt.Println("day 23a:", result)
}

func Day23b() {
	result := day23a(day23bfile())
	fmt.Println("day 23a:", result)
}
