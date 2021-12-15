package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

type Node struct {
	x, y int
}

type Graph struct {
	nodes map[Node]struct{}
	edges map[Node]map[Node]float64
}

func (g *Graph) AddEdge(n1, n2 Node, value int) {
	if g.edges == nil {
		g.edges = make(map[Node]map[Node]float64)
	}
	if g.edges[n1] == nil {
		g.edges[n1] = make(map[Node]float64)
	}
	g.edges[n1][n2] = float64(value)
}

func (g *Graph) AddNode(n Node) {
	if g.nodes == nil {
		g.nodes = make(map[Node]struct{})
	}
	g.nodes[n] = struct{}{}
}

var inf = math.Inf(1)

type pQueue struct {
	queue map[Node]float64
}

func (pq *pQueue) addWithPriority(n Node, p float64) {
	pq.queue[n] = p
}

func (pq *pQueue) decreasePriority(n Node, p float64) {
	pq.queue[n] -= p
}

func (pq *pQueue) extractMin() Node {
	var minn Node
	minp := inf
	var first = true
	for n, p := range pq.queue {
		if first || p < minp {
			p = minp
			minn = n
			first = false
		}
	}
	delete(pq.queue, minn)
	return minn
}

func (g *Graph) Dijkstra(source, dest Node) (dist map[Node]float64, prev map[Node]*Node) {
	pQ := pQueue{queue: make(map[Node]float64, len(g.nodes))}
	dist = make(map[Node]float64, len(g.nodes))
	prev = make(map[Node]*Node)
	dist[source] = 0
	for node := range g.nodes {
		if node != source {
			dist[node] = inf
		}
		pQ.addWithPriority(node, dist[node])
	}

	for len(pQ.queue) > 0 {
		minn := pQ.extractMin()
		if dist[minn] == inf || minn == dest {
			break
		}
		for edge, edgeVal := range g.edges[minn] {
			if _, ok := pQ.queue[edge]; !ok {
				continue
			}

			alt := dist[minn] + edgeVal
			if alt < dist[edge] {
				dist[edge] = alt
				prev[edge] = &minn
				pQ.decreasePriority(edge, alt)
			}
		}
	}

	return dist, prev
}

func day15parseInput(input io.Reader) (width int, grid []int) {
	s := bufio.NewScanner(input)
	for s.Scan() {
		line := s.Text()
		if width == 0 {
			width = len(line)
		}
		for _, c := range line {
			grid = append(grid, int(c-'0'))
		}
	}
	return width, grid
}

func day15a(width int, grid []int) int {
	height := len(grid) / width
	var graph Graph
	for i := range grid {
		y := i / width
		x := i % width
		n1 := Node{x, y}
		graph.AddNode(n1)
		for _, j := range []int{1, 3, 5, 7} {
			diffy := j/3 - 1
			diffx := j%3 - 1
			edgey := y + diffy
			edgex := x + diffx
			if edgey < 0 || edgey > height-1 || edgex < 0 || edgex > width-1 {
				continue
			}
			n2 := Node{edgex, edgey}
			graph.AddEdge(n1, n2, grid[edgey*width+edgex])
		}
	}
	dest := Node{height - 1, width - 1}
	dist, _ := graph.Dijkstra(Node{0, 0}, dest)

	// fmt.Println(dist[Node{0, 2}])
	// for e := prev[dest]; e != nil; e = prev[*e] {
	// 	fmt.Printf("%v\n", *e)
	// }
	total := dist[dest]
	return int(total)
}

func Day15a() {
	f, err := os.Open("input/day15.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day15a(day15parseInput(f))
	fmt.Printf("day 15a: %d\n", result)
}

func day15b(iwidth int, igrid []int) int {
	iheight := len(igrid) / iwidth
	width := iwidth * 5

	grid := make([]int, len(igrid)*25)
	for i := range grid {
		y := i / width
		x := i % width
		origy := y % iheight
		origx := x % iwidth

		inc := y/iheight + x/iwidth

		grid[i] = (igrid[origy*iwidth+origx]+inc-1)%9 + 1
	}

	height := len(grid) / width

	var graph Graph
	for i := range grid {
		y := i / width
		x := i % width
		n1 := Node{x, y}
		graph.AddNode(n1)
		for _, j := range []int{1, 3, 5, 7} {
			diffy := j/3 - 1
			diffx := j%3 - 1
			edgey := y + diffy
			edgex := x + diffx
			if edgey < 0 || edgey > height-1 || edgex < 0 || edgex > width-1 {
				continue
			}
			n2 := Node{edgex, edgey}
			graph.AddEdge(n1, n2, grid[edgey*width+edgex])
		}
	}

	dest := Node{height - 1, width - 1}
	dist, _ := graph.Dijkstra(Node{0, 0}, dest)

	// fmt.Println(dist[Node{0, 2}])
	// for e := prev[dest]; e != nil; e = prev[*e] {
	// 	fmt.Printf("%v\n", *e)
	// }
	total := dist[dest]
	return int(total)
}

func Day15b() {
	f, err := os.Open("input/day15.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day15b(day15parseInput(f))
	fmt.Printf("day 15b: %d\n", result)
}
