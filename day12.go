package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func in_array(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func day12a(input io.Reader) int {
	s := bufio.NewScanner(input)
	graph := map[string][]string{}

	var count int
	var scan func(cave string, visited []string)
	scan = func(cave string, visited []string) {
		if 'a' <= cave[0] && cave[0] <= 'z' && in_array(visited, cave) {
			return
		}
		visited = append(visited, cave)
		if cave == "end" {
			count++
			return
		}
		for _, adj := range graph[cave] {
			scan(adj, visited)
		}
	}

	for s.Scan() {
		parts := strings.Split(s.Text(), "-")
		graph[parts[0]] = append(graph[parts[0]], parts[1])
		graph[parts[1]] = append(graph[parts[1]], parts[0])
	}
	scan("start", []string{})
	return count
}

func Day12a() {
	f, err := os.Open("input/day12.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day12a(f)
	fmt.Printf("day 12a: %d\n", result)
}

func day12b(input io.Reader) int {
	s := bufio.NewScanner(input)
	graph := map[string][]string{}

	var count int
	var scan func(cave string, visited []string, double_remain bool)
	scan = func(cave string, visited []string, double_remain bool) {
		if 'a' <= cave[0] && cave[0] <= 'z' && in_array(visited, cave) {
			if cave == "start" {
				return
			}
			if double_remain {
				double_remain = false
			} else {
				return
			}
		}
		visited = append(visited, cave)
		if cave == "end" {
			count++
			return
		}
		for _, adj := range graph[cave] {
			scan(adj, visited, double_remain)
		}
	}

	for s.Scan() {
		parts := strings.Split(s.Text(), "-")
		graph[parts[0]] = append(graph[parts[0]], parts[1])
		graph[parts[1]] = append(graph[parts[1]], parts[0])
	}
	scan("start", []string{}, true)
	return count
}

func Day12b() {
	f, err := os.Open("input/day12.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := day12b(f)
	fmt.Printf("day 12b: %d\n", result)
}
