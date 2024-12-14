package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Robot struct {
	Pos [2]int
	Vel [2]int
}

const (
	height = 102
	width  = 100
)

func main() {
	input := readInputFile("input.txt")

	for i, r := range input {
		input[i].Pos = calcMoves(r, 100)
	}

	fmt.Println(getDanger(input))

	// Part 2
	input = readInputFile("input.txt")

	xThird := width / 3
	yThird := height / 3

	heigest := 0
	seco := 0
	for sec := range 10000 {
		for i, r := range input {
			input[i].Pos = calcMoves(r, 1)
		}
		danger := 0
		for _, r := range input {
			x := r.Pos[0]
			y := r.Pos[1]
			if x > xThird && x < xThird*2 && y > yThird && y < yThird*2 {
				danger += 1
			}
		}

		if danger > heigest {
			heigest = danger
			seco = sec + 1
		}

	}
	fmt.Println(seco)
}

func getDanger(ro []Robot) int {
	quads := [4]int{}
	xHalf := width / 2
	yHalf := height / 2

	for _, r := range ro {
		x := r.Pos[0]
		y := r.Pos[1]

		if x < xHalf && y < yHalf {
			quads[0] += 1
		}

		if x > xHalf && y < yHalf {
			quads[1] += 1
		}
		if x > xHalf && y > yHalf {
			quads[2] += 1
		}
		if x < xHalf && y > yHalf {
			quads[3] += 1
		}

	}

	return quads[0] * quads[1] * quads[2] * quads[3]
}

func calcMoves(r Robot, m int) [2]int {
	newPos := [2]int{}
	newPos[0] = r.Pos[0]
	newPos[1] = r.Pos[1]

	for range m {
		p1 := newPos[0] + r.Vel[0]
		p2 := newPos[1] + r.Vel[1]

		if p1 < 0 {
			p1 = p1 + (width + 1)
		}
		if p1 > width {
			p1 = p1 - (width + 1)
		}
		if p2 < 0 {
			p2 = p2 + (height + 1)
		}
		if p2 > height {
			p2 = p2 - (height + 1)
		}

		newPos[0] = p1
		newPos[1] = p2
	}

	return newPos
}

func printMap(r []Robot) {

	pos := [][2]int{}
	for _, ro := range r {
		pos = append(pos, ro.Pos)
	}

	for y := range height + 1 {
		for x := range width + 1 {
			if slices.Contains(pos, [2]int{x, y}) {
				fmt.Print("x")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func readInputFile(filename string) []Robot {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}
	lines := strings.Split(string(f), "\n")

	output := []Robot{}
	for _, line := range lines {
		posArr := strings.Split(strings.Split(line, "=")[1], ",")
		p1, _ := strconv.Atoi(posArr[0])
		p2, _ := strconv.Atoi(strings.Split(posArr[1], " ")[0])

		velArr := strings.Split(strings.Split(line, "=")[2], ",")
		v1, _ := strconv.Atoi(velArr[0])
		v2, _ := strconv.Atoi(velArr[1])

		output = append(output, Robot{[2]int{p1, p2}, [2]int{v1, v2}})
	}
	return output
}
