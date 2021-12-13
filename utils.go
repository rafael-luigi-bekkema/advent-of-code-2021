package main

import (
	"bufio"
	"io"
	"log"
	"strconv"

	"github.com/fsnotify/fsnotify"
)

func atoi(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}

func sum(ints []int) (total int) {
	for _, i := range ints {
		total += i
	}
	return total
}

func minmax(ints []int) (min, max int) {
	for i, num := range ints {
		if i == 0 || num < min {
			min = num
		}
		if i == 0 || num > max {
			max = num
		}
	}
	return
}

func scanLines(input io.Reader) []string {
	var lines []string
	s := bufio.NewScanner(input)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func watchFiles(action func(fileName string), flag fsnotify.Op, paths ...string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	handler := func(event fsnotify.Event) {
		if event.Op&flag != 0 {
			action(event.Name)
		}
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				handler(event)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("watch error:", err)
			}
		}
	}()

	for _, path := range paths {
		err = watcher.Add(path)
	}
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
