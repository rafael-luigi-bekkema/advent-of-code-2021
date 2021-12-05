package main

import (
	"fmt"
	"os"
)

func main() {
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
}
