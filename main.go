package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

func main() {
	testWatcherP := flag.Bool("watch", false, "Run test watcher")
	flag.Parse()

	if *testWatcherP {
		testWatcher()
		return
	}

	fmt.Print("advent of code 2021\n\n")
	funcErrs := []func() error{
		Day1a,
		Day1b,
		Day2a,
		Day2b,
		Day3a,
		Day3b,
	}
	for _, f := range funcErrs {
		if err := f(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	// Switched to just using panics
	// no reason to do proper error handling

	Day4a()
	Day4b()
	Day5a()
	Day5b()
	Day6a()
	Day6b()
	Day7a()
	Day7b()

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
