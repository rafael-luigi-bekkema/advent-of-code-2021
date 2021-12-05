package main

import (
	"fmt"
	"os"
)

func main() {
	funcs := []func() error{
		Day1a,
		Day1b,
		Day2a,
		Day2b,
		Day3a,
		Day3b,
	}
	for _, f := range funcs {
		if err := f(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
