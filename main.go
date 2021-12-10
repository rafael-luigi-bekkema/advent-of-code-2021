package main

import (
	"flag"
	"fmt"
	"log"
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
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					runTests(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("watch error:", err)
			}
		}
	}()

	err = watcher.Add("")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
