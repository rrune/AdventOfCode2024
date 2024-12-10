package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input := readInputFile("input.txt")

	positions := map[rune][][2]int{}

	for line := range input {
		for col := range input[line] {
			if input[line][col] != '.' {
				positions[input[line][col]] = append(positions[input[line][col]], [2]int{line, col})
			}
		}
	}

	lineLength := len(input)
	colLength := len(input[0])

	keys := getKeys(positions)

	antinodes := [][2]int{}

	for _, key := range keys {

		for _, pos1 := range positions[key] {

			for _, pos2 := range positions[key] {

				if pos1 != pos2 {

					// (1 8), (2 5) -> (0 11)
					//	2 - 1 = 1
					// 8 - 5 = 3

					// 1-1 = 0
					// 8 + 3 = 11

					// for Part 1: range 1
					// couldn't bothered to calculate the actual needed range
					// vector length/length and stuff
					// just throw compute power at the problem :)
					for i := range colLength {

						antiPos := [2]int{
							pos1[0] - ((pos2[0] - pos1[0]) * i),
							pos1[1] + ((pos1[1] - pos2[1]) * i),
						}

						if (antiPos[0] >= 0 && antiPos[0] < lineLength) && (antiPos[1] >= 0 && antiPos[1] < colLength) {

							if !slices.Contains(antinodes, antiPos) {
								antinodes = append(antinodes, antiPos)
							}
						}

					}

				}

			}

		}

	}

	fmt.Println(len(antinodes))
}

func getKeys(mymap map[rune][][2]int) []rune {
	keys := make([]rune, len(mymap))

	i := 0
	for k := range mymap {
		keys[i] = k
		i++
	}
	return keys
}

func readInputFile(filename string) [][]rune {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	lines := strings.Split(string(content), "\n")
	output := [][]rune{}
	for _, line := range lines {
		output = append(output, []rune(line))
	}
	return output
}
