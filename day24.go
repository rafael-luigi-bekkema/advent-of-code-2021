package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day24file() []string {
	f := must(os.Open("input/day24.txt"))
	defer f.Close()

	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

type ALU struct {
	w, x, y, z int
	zes        [14]int
	step       int
	program    []string
}

func (alu *ALU) String() string {
	return fmt.Sprintf("w=%d,x=%d,y=%d,z=%d", alu.w, alu.x, alu.y, alu.z)
}

func (alu *ALU) run(inputs []int) bool {
	valp := func(s string) *int {
		switch s {
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
	val := func(s string) int {
		if v := valp(s); v != nil {
			return *v
		}
		return must(strconv.Atoi(s))
	}
	inpcount := 0
	for _, line := range alu.program {
		if line[0] == '#' {
			continue
		}
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "inp":
			alu.step++
			*valp(parts[1]) = inputs[inpcount]
			inpcount++
		case "add":
			*valp(parts[1]) += val(parts[2])
		case "mul":
			*valp(parts[1]) *= val(parts[2])
		case "div":
			*valp(parts[1]) /= val(parts[2])
		case "mod":
			*valp(parts[1]) %= val(parts[2])
		case "eql":
			if val(parts[1]) == val(parts[2]) {
				*valp(parts[1]) = 1
			} else {
				*valp(parts[1]) = 0
			}
		default:
			panic("unknown instruction: " + line)
		}
	}

	return alu.z == 0
}

func intToInputs(val int) ([14]int, bool) {
	var inputs [14]int
	for i, r := range fmt.Sprint(val) {
		if r == '0' {
			return inputs, false
		}
		inputs[i] = int(r - '0')
	}
	return inputs, true
}

func incInput(input []int) []int {
	for i, v := range input {
		if v == 9 {
			input[i] = 1
			continue
		}
		input[i]++
		break
	}
	return input
}

func day24try(program []string, inputs []int, min bool) (int, []int) {
	if len(inputs) == 14 {
		alu := ALU{program: program}
		if v := alu.run(inputs); v {
			return 0, inputs
		} else {
			return alu.step, nil
		}
	}
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
	var ints []int
	if min {
		for i := mini; i <= maxi; i++ {
			ints = append(ints, i)
		}
	} else {
		for i := maxi; i >= mini; i-- {
			ints = append(ints, i)
		}
	}
	for _, i := range ints {
		inputs := append(inputs, i)
		failstep, result := day24try(program, inputs, min)
		if result != nil {
			return failstep, result
		}
		if failstep < len(inputs) {
			return failstep, nil
		}
	}
	return len(inputs), nil
}

func day24a(program []string, min bool) int {
	_, result := day24try(day24file(), nil, min)
	var val int
	for i, num := range result {
		val += num * pow(10, len(result)-i-1)
	}
	return val
}
