// Worst code I've ever written

package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input := readInputFile("input.txt")

	var initPos [2]int
	var pos [2]int
	for line := range input {
		for col := range input[line] {
			if input[line][col] == '^' {
				pos[0] = line
				pos[1] = col

				initPos[0] = line
				initPos[1] = col
			}
		}
	}

	numberRotated := 0
	input[pos[0]][pos[1]] = '1'

	/* newObsCount := 0 */

	for pos[0] != 0 {
		if input[pos[0]-1][pos[1]] == '#' {

			input, pos = rotateMap(input, pos)
			numberRotated += 1

		} else {
			pos[0] = pos[0] - 1
			input[pos[0]][pos[1]] = '1'
		}

		/* 		for col := range input[pos[0]] {
			if col > pos[1] && input[pos[0]][col] == '4' {
				if input[pos[0]][col+1] == '#' {

					pos1 := pos[0] - 1
					pos2 := pos[1]

					for range numberRotated {
						temp := pos1
						pos1 = pos2
						pos2 = len(input) - temp - 1
					}

					fmt.Println(pos1, pos2)
					newObsCount++
				}
			}
		} */
		//printMap(input)

	}
	/*	xCount := 0
		for line := range input {
			for col := range input[line] {
				if input[line][col] == 'X' {
					xCount += 1
				}
			}
		}
			for numberRotated%4 != 0 {
				input, pos = rotateMap(input, pos)
				numberRotated += 1
			} */

	xCount := 0
	for line := range input {
		for col := range input[line] {
			if input[line][col] == '1' || input[line][col] == '2' || input[line][col] == '3' || input[line][col] == '4' {
				xCount += 1
			}
		}
	}
	fmt.Println(xCount)
	// fmt.Print(newObsCount)

	// Part 2 brute force

	newObsCount := 0

	input[pos[0]][pos[1]] = 'X'

	for line := range input {
		for col := range input[line] {

			input = readInputFile("input.txt")
			numberRotated = 0

			pos = [2]int{}
			for line := range input {
				for col := range input[line] {
					if input[line][col] == '^' {
						pos[0] = line
						pos[1] = col

						initPos[0] = line
						initPos[1] = col
					}
				}
			}

			//fmt.Println("a")
			if input[line][col] == '.' {
				input[line][col] = '#'

				end := false
				for !end {
					if pos[0] == 0 {
						end = true
					} else if input[pos[0]][pos[1]] == '1' {
						end = true
						newObsCount++
						//fmt.Println(line, col)
					} else {
						if input[pos[0]-1][pos[1]] == '#' {

							input, pos = rotateMap(input, pos)
							numberRotated += 1

						} else {
							input[pos[0]][pos[1]] = 'X'
							pos[0] = pos[0] - 1
						}
					}
					//printMap(input)
				}
				input[line][col] = '.'
			}
		}

	}
	fmt.Print(newObsCount)

}

func saveLast3ObsPos(currPos [2]int, last3Obspos [3][3]int, numberRot int, length int) [3][3]int {
	newPos := [3]int{
		currPos[0],
		currPos[1],
		numberRot,
	}

	for i, pos := range last3Obspos {
		for range numberRot - pos[2] {
			temp := pos[0]
			pos[0] = length - pos[1] - 1
			pos[1] = temp
		}
		pos[2] = numberRot
		last3Obspos[i] = pos
	}

	last3Obspos[0] = last3Obspos[1]
	last3Obspos[1] = last3Obspos[2]
	last3Obspos[2] = newPos

	return last3Obspos

}

func rotateMap(input [][]rune, pos [2]int) ([][]rune, [2]int) {
	newMap := [][]rune{}
	newPos := [2]int{}

	for range input[0] {
		newMap = append(newMap, []rune{})
	}

	for line := range input {
		for col := len(input[line]) - 1; col >= 0; col-- {
			newSym := input[line][col]
			if newSym == 'X' {
				newSym = '2'
			} else if newSym == '2' {
				newSym = '3'
			} else if newSym == '3' {
				newSym = '4'
			} else if newSym == '4' {
				newSym = '1'
			} else if newSym == '1' {
				newSym = '2'
			}

			newMap[col] = append(newMap[col], newSym)
		}
	}

	slices.Reverse[[][]rune](newMap)

	newPos[0] = len(input) - pos[1] - 1
	newPos[1] = pos[0]

	return newMap, newPos
}

func printMap(input [][]rune) {
	fmt.Print(" ")
	for col := range input[0] {
		fmt.Print(col)
	}
	fmt.Println()
	for line := range input {
		fmt.Print(line)
		for col := range input[line] {
			fmt.Print(string(input[line][col]))
		}
		fmt.Println()
	}
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
