package main

import (
	"fmt"
	"os"
)

type pos struct {
	x, y int
}

func (p *pos) step(velocity *pos) {
	p.x += velocity.x
	p.y += velocity.y
	if velocity.x > 0 {
		velocity.x--
	} else if velocity.x < 0 {
		velocity.x++
	}
	velocity.y--
}

type target struct {
	x1, x2, y1, y2 int
}

func (t target) hit(p pos) bool {
	return t.x1 <= p.x && p.x <= t.x2 && t.y1 <= p.y && p.y <= t.y2
}

func day17a(input string) int {
	var t target
	fmt.Sscanf(input, `target area: x=%d..%d, y=%d..%d`, &t.x1, &t.x2, &t.y1, &t.y2)

	totalmax := 0
	for x := minx(t.x1); x <= t.x2; x++ {
	fory:
		for y := t.y1; y <= abs(t.y1); y++ {
			maxy := 0
			startpos := pos{0, 0}
			vel := pos{x, y}
			for curpos := startpos; curpos.y >= t.y1 && curpos.x <= t.x2; curpos.step(&vel) {
				if curpos == startpos || curpos.y > maxy {
					maxy = curpos.y
				}
				if vel.x == 0 && curpos.x < t.x1 {
					break fory
				}
				if t.hit(curpos) {
					if maxy > totalmax {
						totalmax = maxy
					}
					break
				}
			}
		}
	}

	return totalmax
}

func Day17a() {
	data, err := os.ReadFile("input/day17.txt")
	if err != nil {
		panic(data)
	}
	result := day17a(string(data))
	fmt.Printf("day 17a: %d\n", result)
}

func minx(tox int) int {
	i := 0
	total := 0
	for {
		total += i
		if total > tox {
			return i
		}
		i++
	}
}

func day17b(input string) int {
	var t target
	fmt.Sscanf(input, `target area: x=%d..%d, y=%d..%d`, &t.x1, &t.x2, &t.y1, &t.y2)

	var velocities []pos
	for x := minx(t.x1); x <= t.x2; x++ {
	fory:
		for y := t.y1; y <= abs(t.y1); y++ {
			maxy := 0
			startpos := pos{0, 0}
			vel := pos{x, y}
			var curpos pos
			for curpos = startpos; curpos.y >= t.y1 && curpos.x <= t.x2; curpos.step(&vel) {
				if curpos == startpos || curpos.y > maxy {
					maxy = curpos.y
				}
				if t.hit(curpos) {
					velocities = append(velocities, pos{x, y})
					continue fory
				}
			}
		}
	}

	return len(velocities)
}

func Day17b() {
	data, err := os.ReadFile("input/day17.txt")
	if err != nil {
		panic(data)
	}
	result := day17b(string(data))
	fmt.Printf("day 17b: %d\n", result)
}
