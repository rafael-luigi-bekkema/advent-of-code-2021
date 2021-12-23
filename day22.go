package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func day22a(inp io.Reader) int {
	s := bufio.NewScanner(inp)

	grid := make(map[coord3]bool)
	for s.Scan() {
		line := s.Text()
		var fromx, tox, fromy, toy, fromz, toz int
		var flip string
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &flip, &fromx, &tox, &fromy, &toy, &fromz, &toz)
		on := flip == "on"

		for _, from := range []*int{&fromx, &fromy, &fromz} {
			if *from < -50 {
				*from = -50
			}
		}
		for _, to := range []*int{&tox, &toy, &toy} {
			if *to > 50 {
				*to = 50
			}
		}

		for x := fromx; x <= tox; x++ {
			for y := fromy; y <= toy; y++ {
				for z := fromz; z <= toz; z++ {
					if x < -50 || y < -50 || z < -50 || x > 50 || y > 50 || z > 50 {
						continue
					}
					grid[coord3{x, y, z}] = on
				}
			}
		}
	}

	var count int
	for _, on := range grid {
		if on {
			count++
		}
	}

	return count
}

type Cube struct {
	from, to coord3
	on       bool
	overlaps []Cube
}

func (c Cube) size() int {
	return (c.to.x - c.from.x + 1) * (c.to.y - c.from.y + 1) * (c.to.z - c.from.z + 1)
}

// realSize calculates size of non-overlapped parts
func (c Cube) realSize() int {
	xs := []int{c.from.x, c.to.x + 1}
	ys := []int{c.from.y, c.to.y + 1}
	zs := []int{c.from.z, c.to.z + 1}
	for _, o := range c.overlaps {
		xs = append(xs, o.from.x, o.to.x+1)
		ys = append(ys, o.from.y, o.to.y+1)
		zs = append(zs, o.from.z, o.to.z+1)
	}
	sort.Ints(xs)
	sort.Ints(ys)
	sort.Ints(zs)
	var size int
	for ix, x := range xs[1:] {
		for iy, y := range ys[1:] {
		z:
			for iz, z := range zs[1:] {
				cub := Cube{
					from:     coord3{xs[ix], ys[iy], zs[iz]},
					to:       coord3{x - 1, y - 1, z - 1},
					on:       false,
					overlaps: nil,
				}
				for _, o := range c.overlaps {
					if _, ok := o.overlap(cub); ok {
						continue z
					}
				}
				size += cub.size()
			}
		}
	}
	return size
}

func (c Cube) overlap(c2 Cube) (o Cube, ok bool) {
	for i := 0; i < 2; i++ {
		if c.from.x <= c2.to.x && c.to.x >= c2.from.x &&
			c.from.y <= c2.to.y && c.to.y >= c2.from.y &&
			c.from.z <= c2.to.z && c.to.z >= c2.from.z {

			if c.from.x >= c2.from.x {
				o.from.x = c.from.x
			} else {
				o.from.x = c2.from.x
			}
			if c.to.x <= c2.to.x {
				o.to.x = c.to.x
			} else {
				o.to.x = c2.to.x
			}

			if c.from.y >= c2.from.y {
				o.from.y = c.from.y
			} else {
				o.from.y = c2.from.y
			}
			if c.to.y <= c2.to.y {
				o.to.y = c.to.y
			} else {
				o.to.y = c2.to.y
			}

			if c.from.z >= c2.from.z {
				o.from.z = c.from.z
			} else {
				o.from.z = c2.from.z
			}
			if c.to.z <= c2.to.z {
				o.to.z = c.to.z
			} else {
				o.to.z = c2.to.z
			}
			return o, true
		}
		c, c2 = c2, c
	}
	return o, false
}

func day22b(inp io.Reader) int {
	s := bufio.NewScanner(inp)

	var cubes []Cube
	for s.Scan() {
		line := s.Text()
		var cube Cube
		var flip string
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &flip, &cube.from.x, &cube.to.x,
			&cube.from.y, &cube.to.y, &cube.from.z, &cube.to.z)
		cube.on = flip == "on"

		for i := range cubes {
			if o, ok := cubes[i].overlap(cube); ok {
				cubes[i].overlaps = append(cubes[i].overlaps, o)
			}
		}

		if cube.on {
			cubes = append(cubes, cube)
		}
	}

	var count int
	for _, c := range cubes {
		count += c.realSize()
	}
	return count
}
