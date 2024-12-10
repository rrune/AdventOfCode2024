package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInputFile("input.txt")

	count1 := 0
	count2 := 0

	for line := range input {
		for col, el := range input[line] {
			if el == 0 {

				posSlc := [][2]int{}
				posSlc = append(posSlc, [2]int{line, col})

				for i := 1; i < 10; i++ {

					newPosSlc := [][2]int{}

					for _, pos := range posSlc {

						if pos[0] > 0 {
							if input[pos[0]-1][pos[1]] == i {
								newPosSlc = append(newPosSlc, [2]int{pos[0] - 1, pos[1]})
							}
						}

						if pos[0] < len(input)-1 {
							if input[pos[0]+1][pos[1]] == i {
								newPosSlc = append(newPosSlc, [2]int{pos[0] + 1, pos[1]})
							}
						}

						if pos[1] > 0 {
							if input[pos[0]][pos[1]-1] == i {
								newPosSlc = append(newPosSlc, [2]int{pos[0], pos[1] - 1})
							}
						}

						if pos[1] < len(input[0])-1 {
							if input[pos[0]][pos[1]+1] == i {
								newPosSlc = append(newPosSlc, [2]int{pos[0], pos[1] + 1})
							}
						}

					}

					posSlc = newPosSlc

				}

				count2 += len(posSlc)

				posSlc = removeDuplicateInt(posSlc)

				count1 += len(posSlc)

			}
		}
	}
	fmt.Println(count1)
	fmt.Println(count2)
}

func removeDuplicateInt(intSlice [][2]int) [][2]int {
	allKeys := make(map[[2]int]bool)
	list := [][2]int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func readInputFile(filename string) [][]int {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	lines := strings.Split(string(content), "\n")
	output := [][]int{}
	for _, line := range lines {
		lineInt := []int{}
		for _, char := range line {
			i, err := strconv.Atoi(string(char))
			if err != nil {
				panic("Error converting line")
			}
			lineInt = append(lineInt, i)
		}
		output = append(output, lineInt)
	}
	return output
}
