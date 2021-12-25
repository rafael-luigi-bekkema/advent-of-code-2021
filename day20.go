package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func day20(algo, inputImg string, iters int) int {
	width := strings.Index(inputImg, "\n")
	img := []byte(strings.ReplaceAll(inputImg, "\n", ""))
	for iter := 0; iter < iters; iter++ {
		newwidth := width + 2
		newimg := make([]byte, newwidth*newwidth)
		for i := 0; i < len(newimg); i++ {
			row := i/newwidth - 1
			col := i%newwidth - 1
			var code int
			for j := 0; j < 9; j++ {
				oldrow := row + (j/3 - 1)
				oldcol := col + (j%3 - 1)
				code *= 2
				if oldrow < 0 || oldrow > width-1 || oldcol < 0 || oldcol > width-1 {
					if iter%2 == 1 && algo[0] == '#' {
						code += 1
					}
					continue
				}
				if img[oldrow*width+oldcol] == '#' {
					code++
				}
			}
			newimg[i] = algo[code]
		}
		img = newimg
		width = newwidth
	}
	return bytes.Count(img, []byte{'#'})
}

func Day20() {
	data, err := os.ReadFile("input/day20.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.SplitN(string(data), "\n", 3)
	result := day20(lines[0], lines[2], 2)
	fmt.Println("day 20a:", result)

	result = day20(lines[0], lines[2], 50)
	fmt.Println("day 20b:", result)
}
