package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
)

type Nodes = map[coord3]struct{}

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
	beacons Nodes
}

func day19input(input io.Reader) []scanner {
	s := bufio.NewScanner(input)
	var scnrs []scanner
	for {
		if !s.Scan() {
			break
		}
		scnr := scanner{beacons: make(Nodes)}
		fmt.Sscanf(s.Text(), "--- scanner %d ---", &scnr.num)

		for s.Scan() {
			if s.Text() == "" {
				break
			}
			var bacon coord3
			fmt.Sscanf(s.Text(), "%d,%d,%d", &bacon.x, &bacon.y, &bacon.z)
			scnr.beacons[bacon] = struct{}{}
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
	newNodes := make(Nodes)
	for node := range nodes {
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
		newNodes[newnode] = struct{}{}
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

type rotCache struct {
	rot [2]rot
	num int
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
	for _, angA := range []int{0, 180} {
		for _, angB := range []int{90, -90} {
			result = append(result, [2]rot{{axX, angA}, {axY, angB}})
			result = append(result, [2]rot{{axX, angA}, {axZ, angB}})

			result = append(result, [2]rot{{axY, angA}, {axX, angB}})
			result = append(result, [2]rot{{axY, angA}, {axZ, angB}})

			result = append(result, [2]rot{{axZ, angA}, {axX, angB}})
			result = append(result, [2]rot{{axZ, angA}, {axY, angB}})
		}
	}
	return
}

func day19a(input []scanner) int {
	grid := make(Nodes)
	for k, v := range input[0].beacons {
		grid[k] = v
	}

	// cache := map[rotCache]Nodes{}

	left := input[1:]
	for i := 0; i < len(left); i++ {
		scnr := left[i]
	baconA:
		for baconA := range grid {
			for _, combo := range combos() {
				beacons := rotates(scnr.beacons, combo)
				for baconBa := range beacons {
					match := 1
					var others []coord3
					for baconBb := range beacons {
						if baconBa == baconBb {
							continue
						}
						deltax := baconBb.x - baconBa.x
						deltay := baconBb.y - baconBa.y
						deltaz := baconBb.z - baconBa.z

						n := coord3{x: baconA.x + deltax, y: baconA.y + deltay, z: baconA.z + deltaz}
						if _, ok := grid[n]; ok {
							match++
						} else {
							others = append(others, n)
						}
					}
					if match >= 12 {
						// for _, other := range others {
						// 	grid[other] = struct{}{}
						// }
						left[i], left[len(left)-1] = left[len(left)-1], left[i]
						left = left[:len(left)-1]
						i = 0
						fmt.Printf("num: %v %v\n", scnr.num, match)
						break baconA
					}
				}
			}
		}
	}
	fmt.Println("left", len(left))
	return 0
}
