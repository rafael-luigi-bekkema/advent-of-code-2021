package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ALU struct {
	w, x, y, z int
	program    []func(chan int)
}

func (alu *ALU) String() string {
	return fmt.Sprintf("w=%d,x=%d,y=%d,z=%d", alu.w, alu.x, alu.y, alu.z)
}

func (alu *ALU) ValPtr(name string) *int {
	switch name {
	case "w":
		return &alu.w
	case "x":
		return &alu.x
	case "y":
		return &alu.y
	case "z":
		return &alu.z
	}
	return nil
}

func (alu *ALU) Val(name string) int {
	if v := alu.ValPtr(name); v != nil {
		return *v
	}
	n, _ := strconv.Atoi(name)
	return n
}

func (alu *ALU) Input(into *int, inp int) {
	*into = inp
}

func (alu *ALU) Add(into *int, val int) {
	*into += val
}

func (alu *ALU) Mul(into *int, val int) {
	*into *= val
}

func (alu *ALU) Div(into *int, val int) {
	*into /= val
}

func (alu *ALU) Mod(into *int, val int) {
	*into %= val
}

func (alu *ALU) Eql(into *int, val int) {
	if *into == val {
		*into = 1
	} else {
		*into = 0
	}
}

func NewALU(inp io.Reader) *ALU {
	var alu ALU
	s := bufio.NewScanner(inp)
	for s.Scan() {
		parts := strings.Split(s.Text(), " ")
		var f func(chan int)
		switch parts[0] {
		case "inp":
			f = func(c chan int) {
				alu.Input(alu.ValPtr(parts[1]), <-c)
			}
		case "add":
			f = func(chan int) {
				alu.Add(alu.ValPtr(parts[1]), alu.Val(parts[2]))
			}
		case "mul":
			f = func(chan int) {
				alu.Mul(alu.ValPtr(parts[1]), alu.Val(parts[2]))
			}
		case "div":
			f = func(chan int) {
				alu.Div(alu.ValPtr(parts[1]), alu.Val(parts[2]))
			}
		case "mod":
			f = func(chan int) {
				alu.Mod(alu.ValPtr(parts[1]), alu.Val(parts[2]))
			}
		case "eql":
			f = func(chan int) {
				alu.Eql(alu.ValPtr(parts[1]), alu.Val(parts[2]))
			}
		default:
			panic("unknown instruction: " + parts[0])
		}
		alu.program = append(alu.program, f)

	}
	return &alu
}

func (alu *ALU) run(inputs []int) bool {
	inp := make(chan int, len(inputs))
	for _, i := range inputs {
		inp <- i
	}
	for _, f := range alu.program {
		f(inp)
	}
	return alu.z == 0
}

func day24minmax(step int, inputs []int) (int, int) {
	var mini, maxi int
	switch len(inputs) + 1 {
	case 1:
		mini = 8
		maxi = 9
	case 2:
		mini = 1
		maxi = 2
	case 3:
		mini = 5
		maxi = 9
	case 4:
		mini = 1
		maxi = 6
	case 5:
		mini = inputs[3] + 3
	case 6:
		mini = inputs[2] - 4
	case 7:
		mini = 7
		maxi = 9
	case 8:
		mini = inputs[6] - 6
	case 9:
		mini = 1
		maxi = 4
	case 10:
		mini = inputs[8] + 5
	case 11:
		mini = 1
		maxi = 7
	case 12:
		mini = inputs[10] + 2
	case 13:
		mini = inputs[1] + 7
	case 14:
		mini = inputs[0] - 7
	default:
		mini = 1
		maxi = 9
	}
	if maxi == 0 {
		maxi = mini
	}
	return mini, maxi
}

func day24try(alu *ALU, inputs []int, out chan int) {
	if len(inputs) == 14 {
		if v := alu.run(inputs); v {
			out <- intsToInt(inputs)
			return
		}
		panic("noo!")
	}
	mini, maxi := day24minmax(len(inputs)+1, inputs)
	for i := mini; i <= maxi; i++ {
		inputs := append(inputs, i)
		day24try(alu, inputs, out)
	}
}

func intsToInt(result []int) int {
	var val int
	for i := range result {
		val = val*10 + result[i]
	}
	return val
}

func day24alu() *ALU {
	f, err  :=os.Open("input/day24.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return NewALU(f)
}

func day24a(alu *ALU, min bool) int {
	out := make(chan int)
	go func() {
		defer close(out)
		day24try(alu, nil, out)
	}()
	var results []int
	for result := range out {
		results = append(results, result)
	}
	minv, maxv := minmax(results)
	if min {
		return minv
	}
	return maxv
}
