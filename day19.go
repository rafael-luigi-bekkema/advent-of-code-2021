package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"runtime/pprof"
	"sort"
)

type nodeType int

const (
	ntScanner nodeType = iota
	ntBacon
)

type Nodes = map[coord3]nodeType

type ax string

const (
	axX ax = "x"
	axY ax = "y"
	axZ ax = "z"
)

type coord3 struct {
	x, y, z int
}

type scanner struct {
	num     int
	beacons []coord3
}

func day19input(input io.Reader) []scanner {
	s := bufio.NewScanner(input)
	var scnrs []scanner
	for {
		if !s.Scan() {
			break
		}
		scnr := scanner{}
		fmt.Sscanf(s.Text(), "--- scanner %d ---", &scnr.num)

		for s.Scan() {
			if s.Text() == "" {
				break
			}
			var bacon coord3
			fmt.Sscanf(s.Text(), "%d,%d,%d", &bacon.x, &bacon.y, &bacon.z)
			scnr.beacons = append(scnr.beacons, bacon)
		}

		scnrs = append(scnrs, scnr)
	}
	return scnrs
}

func rotate(deg int, ax ax, nodes Nodes) Nodes {
	if deg == 0 || deg%360 == 0 {
		return nodes
	}
	// https://www.khanacademy.org/computing/computer-programming/programming-games-visualizations/programming-3d-shapes/a/rotating-3d-shapes
	theta := float64(deg) * (math.Pi / 180)
	var sinTheta = math.Sin(theta)
	var cosTheta = math.Cos(theta)
	newNodes := make(Nodes, len(nodes))
	for node, typ := range nodes {
		var newnode coord3
		switch ax {
		case axZ:
			newnode.x = int(math.Round(float64(node.x)*cosTheta - float64(node.y)*sinTheta))
			newnode.y = int(math.Round(float64(node.y)*cosTheta + float64(node.x)*sinTheta))
			newnode.z = node.z
		case axX:
			newnode.y = int(math.Round(float64(node.y)*cosTheta - float64(node.z)*sinTheta))
			newnode.z = int(math.Round(float64(node.z)*cosTheta + float64(node.y)*sinTheta))
			newnode.x = node.x
		case axY:
			newnode.x = int(math.Round(float64(node.x)*cosTheta + float64(node.z)*sinTheta))
			newnode.z = int(math.Round(float64(node.z)*cosTheta - float64(node.x)*sinTheta))
			newnode.y = node.y
		default:
			panic("unknown ax")
		}
		newNodes[newnode] = typ
	}
	return newNodes
}

type rot struct {
	ax    ax
	angle int
}

func (r rot) String() string {
	return fmt.Sprintf("rot: %s: %v", r.ax, r.angle)
}

func rotates(nodes Nodes, rots [2]rot) Nodes {
	for _, rot := range rots {
		nodes = rotate(rot.angle, rot.ax, nodes)
	}
	return nodes
}

func toSlice(nodes Nodes) (result []coord3) {
	for node := range nodes {
		result = append(result, node)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].x != result[j].x {
			return result[i].x < result[j].x
		}
		if result[i].y != result[j].y {
			return result[i].y < result[j].y
		}
		return result[i].z < result[j].z
	})
	return
}

func combos() (result [][2]rot) {
	// no idea why this works
	// too hard
	// but it works
	for i := 0; i < 3; i++ {
		result = append(result, [2]rot{{axX, 0}, {axY, 0}})
		result = append(result, [2]rot{{axX, 0}, {axY, 90}})
		result = append(result, [2]rot{{axZ, 90}, {axY, 0}})

		result = append(result, [2]rot{{axX, 180}, {axY, 0}}) // ax flip
		result = append(result, [2]rot{{axX, 0}, {axY, 90}})
		result = append(result, [2]rot{{axZ, 90}, {axY, 0}})
	}
	return
}

func distance(items []coord3) int {
	dist := 0
	for i, itemA := range items {
		for _, itemB := range items[i+1:] {
			val := abs(itemA.x-itemB.x) + abs(itemA.y-itemB.y) + abs(itemA.z-itemB.z)
			if val > dist {
				dist = val
			}
		}
	}
	return dist
}

func day19a(input []scanner) (int, int) {
	f, _ := os.Create("/tmp/aoc_profile.txt")
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	grid := Nodes{coord3{0, 0, 0}: ntScanner}
	for _, v := range input[0].beacons {
		grid[v] = ntBacon
	}

	rotateCount := 0
	others := make([]coord3, 0, len(grid))
	left := input[1:]
outer:
	for len(left) > 0 {
		change := false
		for _, combo := range combos() {
			grid = rotates(grid, combo)
			for i, scnr := range left {
				for baconA, typ := range grid {
					if typ == ntScanner {
						continue
					}
					beacons := scnr.beacons
					for _, baconBa := range beacons {
						match := 1
						others = others[:0]
						for _, baconBb := range beacons {
							if baconBa == baconBb {
								continue
							}
							deltax := baconBb.x - baconBa.x
							deltay := baconBb.y - baconBa.y
							deltaz := baconBb.z - baconBa.z

							n := coord3{x: baconA.x + deltax, y: baconA.y + deltay,
								z: baconA.z + deltaz}
							if _, ok := grid[n]; ok {
								match++
							} else {
								others = append(others, n)
							}
						}
						if match >= 12 {
							for _, other := range others {
								grid[other] = ntBacon
							}
							left[i], left[len(left)-1] = left[len(left)-1], left[i]
							left = left[:len(left)-1]
							fmt.Printf("num: %v %v %v left: %v\n", scnr.num, match, combo, len(left))
							scanner := coord3{x: baconA.x - baconBa.x, y: baconA.y - baconBa.y,
								z: baconA.z - baconBa.z}
							change = true
							grid[scanner] = ntScanner
							continue outer
						}
					}
				}
			}
		}
		if !change {
			break
		}
	}
	fmt.Println("rotates:", rotateCount)
	var scanners []coord3
	for node, typ := range grid {
		if typ == ntScanner {
			scanners = append(scanners, node)
		}
	}
	if len(left) > 0 {
		panic("unmatched beacons")
	}
	return len(grid) - len(input), distance(scanners)
}

func Day19() {
	f, err := os.Open("input/day19.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	resultA, resultB := day19a(day19input(f))
	fmt.Printf("day 19: beacons: %d distance: %d\n", resultA, resultB)
}
