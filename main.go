package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime/pprof"

	"github.com/fsnotify/fsnotify"
)

func main() {
	cpuprofile := flag.String("profile", "", "")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
		Day19()
		return
	}

	testWatcherP := flag.Bool("watch", false, "Run test watcher")
	flag.Parse()

	if *testWatcherP {
		testWatcher()
		return
	}

	fmt.Print("advent of code 2021\n\n")

	Day1a()
	Day1b()
	Day2a()
	Day2b()
	Day3a()
	Day3b()
	Day4a()
	Day4b()
	Day5a()
	Day5b()
	Day6a()
	Day6b()
	Day7a()
	Day7b()
	Day8a()
	Day8b()
	Day9a()
	Day9b()
	Day10a()
	Day10b()
	Day11a()
	Day11b()
	Day12a()
	Day12b()
	Day13a()
	Day13b()
	Day14a()
	Day14b()
	Day15a()
	Day15b()
	Day16a()
	Day16b()
	Day17a()
	Day17b()
	Day18a()
	Day18b()
	Day19()
	Day20()

	fmt.Print("\n\n")
}

func runTests(name string) {
	cmd := exec.Command("go", "test", "-timeout", "5s", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func testWatcher() {
	runTests("")

	watchFiles(runTests, fsnotify.Write|fsnotify.Rename|fsnotify.Remove, "")
}
