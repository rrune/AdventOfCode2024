package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInputFile("input.txt")

	// PART 1

	xmasCount := 0
	// lines
	for _, line := range input {
		// right to left
		for i := 0; i < len(line); i++ {
			if line[i] == 'X' {
				// forwards
				if i < (len(line) - 3) {
					if line[i+1] == 'M' && line[i+2] == 'A' && line[i+3] == 'S' {
						//fmt.Println("xmas, horizontal, forwards")
						xmasCount++
					}
				}

				// backwards
				if i >= 3 {
					if line[i-1] == 'M' && line[i-2] == 'A' && line[i-3] == 'S' {
						//fmt.Println("xmas, horizontal, backwards")
						xmasCount++
					}
				}
			}
		}
	}

	// columns
	for column := 0; column < len(input[0]); column++ {
		for i := 0; i < len(input); i++ {
			if input[i][column] == 'X' {
				// forwards
				if i < (len(input[0]) - 3) {
					if input[i+1][column] == 'M' && input[i+2][column] == 'A' && input[i+3][column] == 'S' {
						//fmt.Println("xmas, vertical, forwards")
						xmasCount++
					}
				}

				// backwards
				if i >= 3 {
					if input[i-1][column] == 'M' && input[i-2][column] == 'A' && input[i-3][column] == 'S' {
						//fmt.Println("xmas, vertical, backwards")
						xmasCount++
					}
				}
			}
		}
	}

	// diagonal right
	for line := 0; line < len(input); line++ {
		for column := 0; column < len(input[line]); column++ {
			if input[line][column] == 'X' {
				// forwards
				if (line < (len(input) - 3)) && (column < (len(input[0]) - 3)) {
					if input[line+1][column+1] == 'M' && input[line+2][column+2] == 'A' && input[line+3][column+3] == 'S' {
						//fmt.Println("xmas, diagonal right, forwards")
						xmasCount++
					}
				}

				// backwards
				if line >= 3 && column >= 3 {
					if input[line-1][column-1] == 'M' && input[line-2][column-2] == 'A' && input[line-3][column-3] == 'S' {
						//fmt.Println("xmas, diagnonal right, backwards")
						xmasCount++
					}
				}
			}
		}
	}

	// diagonal left
	for line := 0; line < len(input); line++ {
		for column := 0; column < len(input[line]); column++ {
			if input[line][column] == 'X' {
				// forwards
				if (line < (len(input) - 3)) && (column >= 3) {
					if input[line+1][column-1] == 'M' && input[line+2][column-2] == 'A' && input[line+3][column-3] == 'S' {
						//fmt.Println("xmas, diagonal left, forwards")
						xmasCount++
					}
				}

				// backwards
				if line >= 3 && (column < (len(input[0]) - 3)) {
					if input[line-1][column+1] == 'M' && input[line-2][column+2] == 'A' && input[line-3][column+3] == 'S' {
						//fmt.Println("xmas, diagnonal left, backwards")
						xmasCount++
					}
				}
			}
		}
	}

	fmt.Println(xmasCount)

	// PART 2
	masCount := 0
	for line := 0; line < len(input); line++ {
		for column := 0; column < len(input[line]); column++ {
			if input[line][column] == 'A' {
				// bounds
				if line >= 1 && column >= 1 && line < (len(input)-1) && column < (len(input[line])-1) {
					// top left - bottom right
					if (input[line-1][column-1] == 'M' && input[line+1][column+1] == 'S') || (input[line-1][column-1] == 'S' && input[line+1][column+1] == 'M') {
						// top right - bottom left
						if (input[line-1][column+1] == 'M' && input[line+1][column-1] == 'S') || (input[line-1][column+1] == 'S' && input[line+1][column-1] == 'M') {
							masCount++
						}
					}
				}
			}
		}
	}
	fmt.Println(masCount)
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
