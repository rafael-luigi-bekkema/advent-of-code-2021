package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

type coord3 struct {
	x, y, z int
}

type scanner struct {
	num     int
	beacons map[coord3]struct{}
}

func day19input(input io.Reader) []scanner {
	s := bufio.NewScanner(input)
	var scnrs []scanner
	for {
		if !s.Scan() {
			break
		}
		scnr := scanner{beacons: make(map[coord3]struct{})}
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

func rotate(theta float64, ax byte, nodes map[coord3]struct{}) map[coord3]struct{} {
	theta = theta * (math.Pi / 180)
	var sinTheta = math.Sin(theta)
	var cosTheta = math.Cos(theta)
	newNodes := make(map[coord3]struct{})
	for node := range nodes {
		switch ax {
		case 'z':
			node.x = int(math.Round(float64(node.x)*cosTheta - float64(node.y)*sinTheta))
			node.y = int(math.Round(float64(node.y)*cosTheta - float64(node.x)*sinTheta))
		case 'x':
			node.y = int(math.Round(float64(node.y)*cosTheta - float64(node.z)*sinTheta))
			node.z = int(math.Round(float64(node.z)*cosTheta - float64(node.y)*sinTheta))
		case 'y':
			node.x = int(math.Round(float64(node.x)*cosTheta - float64(node.z)*sinTheta))
			node.z = int(math.Round(float64(node.z)*cosTheta - float64(node.x)*sinTheta))
		}
		newNodes[node] = struct{}{}
	}
	return newNodes
}

func day19a(input []scanner) int {
	grid := make(map[coord3]struct{})
	for k, v := range input[0].beacons {
		grid[k] = v
	}
	// input[0], input[1] = input[1], input[0]
	mainBeacons := input[0].beacons
	fmt.Println(mainBeacons)
	mainBeacons = rotate(180, 'y', mainBeacons)
	fmt.Println()
	fmt.Println(mainBeacons)
	// mainBeacons = rotateZ(90, 'y', mainBeacons)
	for _, scnr := range input[1:] {
	baconA:
		for baconA := range mainBeacons {
			//for _, rotz := range []float64{0, 90, 180, 270} {
			//beacons := rotateZ(rotz, 'z', scnr.beacons)
			beacons := scnr.beacons
			for baconBa := range beacons {
				var match int
				for baconBb := range beacons {
					if baconBa == baconBb {
						continue
					}
					deltax := baconBb.x - baconBa.x
					deltay := baconBb.y - baconBa.y
					deltaz := baconBb.z - baconBa.z

					n := coord3{x: baconA.x + deltax, y: baconA.y + deltay, z: baconA.z + deltaz}
					if _, ok := mainBeacons[n]; ok {
						match++
					}
				}
				if match > 1 {
					fmt.Println(scnr.num, match)
					break baconA
				}
			}
			//}
		}
	}
	return 0
}
