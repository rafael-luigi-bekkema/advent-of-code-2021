package main

import (
	"fmt"
	"os"
)

func main() {
	funcs := []func() error{
		Day1a,
	}
	for _, f := range funcs {
		if err := f(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
