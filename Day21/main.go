package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	numpad = map[rune][2]int{
		'7': {0, 3},
		'8': {1, 3},
		'9': {2, 3},
		'4': {0, 2},
		'5': {1, 2},
		'6': {2, 2},
		'1': {0, 1},
		'2': {1, 1},
		'3': {2, 1},
		'0': {1, 0},
		'A': {2, 0},
	}

	dirpad = map[rune][2]int{
		'^': {1, 1},
		'A': {2, 1},
		'<': {0, 0},
		'v': {1, 0},
		'>': {2, 0},
	}
)

type Path struct {
	s [2]int
	e [2]int
	r int
}

func main() {
	input := readInputFile("input.txt")

	cache := map[Path]int{}

	fmt.Println(solve(input, 2, cache))
	fmt.Println(solve(input, 25, cache))
}

func solve(input [][]rune, robots int, cache map[Path]int) int {
	total := 0
	for _, in := range input {
		in = append([]rune{'A'}, in...)
		nums := ""
		for i := 0; i < len(in)-1; i++ {
			start := numpad[in[i]]
			end := numpad[in[i+1]]
			seq := getSequence(start, end, true)
			nums += seq
		}

		count := 0
		prev := []rune(nums)
		prev = append([]rune{'A'}, prev...)
		for i := 0; i < len(prev)-1; i++ {
			start := dirpad[prev[i]]
			end := dirpad[prev[i+1]]
			count += recSolver(start, end, robots-1, cache)
		}

		inputNum, _ := strconv.Atoi(string(in[1:4]))
		total += inputNum * count
	}

	return total
}

func recSolver(start [2]int, end [2]int, robot int, memo map[Path]int) int {
	p := Path{start, end, robot}
	if val, ok := memo[p]; ok {
		return val
	}

	count := 0
	seq := getSequence(start, end, false)

	if robot == 0 {
		return len(seq)
	}
	seq = "A" + seq
	for i := 0; i < len(seq)-1; i++ {
		count += recSolver(dirpad[rune(seq[i])], dirpad[rune(seq[i+1])], robot-1, memo)
	}

	memo[p] = count
	return count
}

func getSequence(start [2]int, end [2]int, num bool) string {
	sequence := ""

	dx := (start[0] - end[0])
	dy := (start[1] - end[1])

	horizMovement := ""
	for range absInt(dx) {
		if dx >= 0 {
			horizMovement += "<"
		} else {
			horizMovement += ">"
		}
	}

	vertMovement := ""
	for range absInt(dy) {
		if dy >= 0 {
			vertMovement += "v"
		} else {
			vertMovement += "^"
		}
	}

	if num {

		if start[1] == 0 && end[0] == 0 {
			sequence += vertMovement + horizMovement
		} else if start[0] == 0 && end[1] == 0 {
			sequence += horizMovement + vertMovement
		} else if dx >= 0 {
			sequence += horizMovement + vertMovement
		} else {
			sequence += vertMovement + horizMovement
		}
	} else {
		if start[0] == 0 && end[1] == 1 {
			sequence += horizMovement + vertMovement
		} else if start[1] == 1 && end[0] == 0 {
			sequence += vertMovement + horizMovement
		} else if dx >= 0 {
			sequence += horizMovement + vertMovement
		} else {
			sequence += vertMovement + horizMovement
		}
	}

	return sequence + "A"
}

func absInt(x int) int {
	if x > 0 {
		return x
	}
	return x * -1

}

func readInputFile(filename string) [][]rune {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}
	lines := strings.Split(string(f), "\n")
	output := [][]rune{}
	for _, l := range lines {
		output = append(output, []rune(l))
	}
	return output
}
