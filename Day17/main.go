package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	a, b, c, program := readInputFile("input.txt")
	fmt.Println(run(a, b, c, program))

	prgmStr := []string{}
	for _, i := range program {
		prgmStr = append(prgmStr, strconv.Itoa(i))
	}
	s := strings.Join(prgmStr, ",")

	z := 1
	running := true
	for running {
		out := run(z, b, c, program)
		for !strings.Contains(s, out) {
			z += 1
			out = run(z, b, c, program)
		}

		if s == out {
			fmt.Println(z)
			return
		}

		z = z * 8
	}
}

func run(a int, b int, c int, program []int) string {
	output := ""
	ip := 0
	for ip < len(program) {
		opcode := program[ip]
		operand := program[ip+1]

		switch opcode {
		case 0:
			// adv
			a = a / powInt(2, getComboOperand(operand, a, b, c))

		case 1:
			// bxl
			b = b ^ operand

		case 2:
			// bst
			b = getComboOperand(operand, a, b, c) % 8

		case 4:
			// bxc
			b = b ^ c

		case 5:
			// out
			output += fmt.Sprint(getComboOperand(operand, a, b, c) % 8)
			output += ","

		case 6:
			// bdv
			b = a / powInt(2, getComboOperand(operand, a, b, c))

		case 7:
			// cdv
			c = a / powInt(2, getComboOperand(operand, a, b, c))

		case 3:
			// jnz
			if a != 0 {
				ip = operand - 2
			}
		}

		ip += 2
	}

	return output[:len(output)-1]
}

func getComboOperand(operand int, a int, b int, c int) int {
	switch operand {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	}
	return -1
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func readInputFile(filename string) (a int, b int, c int, program []int) {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}
	parts := strings.Split(string(f), "\n\n")

	registerLines := strings.Split(parts[0], "\n")
	a, _ = strconv.Atoi(registerLines[0][12:])
	b, _ = strconv.Atoi(registerLines[1][12:])
	c, _ = strconv.Atoi(registerLines[2][12:])

	programStr := strings.Split(parts[1][9:], ",")
	program = []int{}
	for _, s := range programStr {
		i, _ := strconv.Atoi(s)
		program = append(program, i)
	}
	return
}
