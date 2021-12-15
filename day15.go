package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"math"
	"os"
)

type pNode struct {
	n Node
	p float64
}
type FloatHeap []pNode

func (h FloatHeap) Len() int           { return len(h) }
func (h FloatHeap) Less(i, j int) bool { return h[i].p < h[j].p }
func (h FloatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FloatHeap) Push(x interface{}) {
	// Push and Pop use pofloat64er receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(pNode))
}

func (h *FloatHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

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

func (g *Graph) Dijkstra(source, dest Node) (dist map[Node]float64, prev map[Node]*Node) {
	inf := math.Inf(1)
	pQ := &FloatHeap{pNode{source, 0}}
	heap.Init(pQ)
	dist = make(map[Node]float64)
	prev = make(map[Node]*Node)
	visited := make(map[Node]struct{})
	for node := range g.nodes {
		dist[node] = inf
	}
	dist[source] = 0

	for pQ.Len() > 0 {
		pN := heap.Pop(pQ).(pNode)
		minn := pN.n

		if dist[minn] == inf {
			panic("nooo")
		}
		if minn == dest {
			break
		}
		visited[minn] = struct{}{}

		for edge, edgeVal := range g.edges[minn] {
			if _, ok := visited[edge]; ok {
				continue
			}

			alt := dist[minn] + edgeVal
			if alt < dist[edge] {
				dist[edge] = alt
				prev[edge] = &minn
				heap.Push(pQ, pNode{edge, alt})
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
