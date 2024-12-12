package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Plot struct {
	Char rune
	Pos  [][2]int
}

func main() {
	input := readInputFile("input.txt")

	visited := [][2]int{}
	plots := [][][2]int{}
	for line := range input {
		for col := range input[line] {

			if !slices.Contains(visited, [2]int{line, col}) {

				plot, _ := recGetPlots(input, [2]int{line, col}, [][2]int{})
				plots = append(plots, plot)
				visited = append(visited, plot...)

			}

		}
	}

	// Part 1
	count1 := 0

	for _, plot := range plots {
		outsides := 0

		for _, cords := range plot {
			// up
			if !slices.Contains(plot, [2]int{cords[0] - 1, cords[1]}) {
				outsides += 1
			}
			// down
			if !slices.Contains(plot, [2]int{cords[0] + 1, cords[1]}) {
				outsides += 1
			}
			// left
			if !slices.Contains(plot, [2]int{cords[0], cords[1] - 1}) {
				outsides += 1
			}
			// right
			if !slices.Contains(plot, [2]int{cords[0], cords[1] + 1}) {
				outsides += 1
			}
		}

		count1 += len(plot) * outsides
	}

	fmt.Println(count1)

	// Part 2

	count2 := 0
	for _, plot := range plots {
		cornerCount := 0
		for _, cords := range plot {
			l := cords[0]
			c := cords[1]

			// convex

			// .B
			// BA
			if !slices.Contains(plot, [2]int{l, c - 1}) && !slices.Contains(plot, [2]int{l - 1, c}) {
				cornerCount += 1
			}

			// B.
			// AB
			if !slices.Contains(plot, [2]int{l, c + 1}) && !slices.Contains(plot, [2]int{l - 1, c}) {
				cornerCount += 1
			}

			// AB
			// B.
			if !slices.Contains(plot, [2]int{l, c + 1}) && !slices.Contains(plot, [2]int{l + 1, c}) {
				cornerCount += 1
			}

			// BA
			// .B
			if !slices.Contains(plot, [2]int{l, c - 1}) && !slices.Contains(plot, [2]int{l + 1, c}) {
				cornerCount += 1
			}

			// concave

			// BA
			// AA
			if slices.Contains(plot, [2]int{l - 1, c}) && slices.Contains(plot, [2]int{l, c - 1}) && !slices.Contains(plot, [2]int{l - 1, c - 1}) {
				cornerCount += 1
			}

			// AB
			// AA
			if slices.Contains(plot, [2]int{l - 1, c}) && slices.Contains(plot, [2]int{l, c + 1}) && !slices.Contains(plot, [2]int{l - 1, c + 1}) {
				cornerCount += 1
			}

			// AA
			// AB
			if slices.Contains(plot, [2]int{l + 1, c}) && slices.Contains(plot, [2]int{l, c + 1}) && !slices.Contains(plot, [2]int{l + 1, c + 1}) {
				cornerCount += 1
			}

			// AA
			// BA
			if slices.Contains(plot, [2]int{l + 1, c}) && slices.Contains(plot, [2]int{l, c - 1}) && !slices.Contains(plot, [2]int{l + 1, c - 1}) {
				cornerCount += 1
			}

		}

		count2 += len(plot) * cornerCount
	}
	fmt.Println(count2)
}

func recGetPlots(input [][]rune, pos [2]int, visited [][2]int) ([][2]int, [][2]int) {
	visited = append(visited, pos)
	l := pos[0]
	c := pos[1]
	char := input[l][c]

	plot := [][2]int{pos}

	if l > 0 && input[l-1][c] == char && !slices.Contains(visited, [2]int{l - 1, c}) {
		pl := [][2]int{}
		pl, visited = recGetPlots(input, [2]int{l - 1, c}, visited)
		plot = append(plot, pl...)
	}
	if l < len(input)-1 && input[l+1][c] == char && !slices.Contains(visited, [2]int{l + 1, c}) {
		pl := [][2]int{}
		pl, visited = recGetPlots(input, [2]int{l + 1, c}, visited)
		plot = append(plot, pl...)
	}
	if c > 0 && input[l][c-1] == char && !slices.Contains(visited, [2]int{l, c - 1}) {
		pl := [][2]int{}
		pl, visited = recGetPlots(input, [2]int{l, c - 1}, visited)
		plot = append(plot, pl...)
	}
	if c < len(input[0])-1 && input[l][c+1] == char && !slices.Contains(visited, [2]int{l, c + 1}) {
		pl := [][2]int{}
		pl, visited = recGetPlots(input, [2]int{l, c + 1}, visited)
		plot = append(plot, pl...)
	}

	return plot, visited
}

func readInputFile(filename string) [][]rune {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	var output [][]rune
	lines := strings.Split(string(f), "\n")
	for _, line := range lines {
		output = append(output, []rune(line))
	}
	return output
}
