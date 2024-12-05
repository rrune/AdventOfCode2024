package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	instructions, orders := readInputFile("input.txt")

	reverseInstrMap := map[int][]int{}
	for _, instruction := range instructions {
		reverseInstrMap[instruction[1]] = append(reverseInstrMap[instruction[1]], instruction[0])
	}

	sumCor := 0
	sumInc := 0
	for _, order := range orders {
		correct := true
		previous := []int{}

		for _, el := range order {
			used := []int{}
			for _, elem := range reverseInstrMap[el] {
				if slices.Contains(order, elem) {
					used = append(used, elem)
				}
			}

			for _, elem := range used {
				if !(slices.Contains(previous, elem)) {
					correct = false
				}
			}
			previous = append(previous, el)
		}

		if correct {
			sumCor += order[len(order)/2]
		} else {
			// PART 2
			correctOrder := []int{}
			for len(order) != 0 {
				el := order[0]

				allowed := true
				for _, elem := range order {
					if slices.Contains(reverseInstrMap[elem], el) {
						allowed = false
					}
				}
				if allowed {
					order = order[1:]
					correctOrder = append(correctOrder, el)
				} else {
					order = append(order[1:], el)
				}

			}
			sumInc += correctOrder[len(correctOrder)/2]
		}
	}

	fmt.Println(sumCor)
	fmt.Println(sumInc)

}

func readInputFile(filename string) ([][]int, [][]int) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	both := strings.Split(string(content), "\n\n")
	instrLines := strings.Split(both[0], "\n")
	ordersLines := strings.Split(both[1], "\n")

	instr := convertTo2DIntSlice(instrLines, "|")
	orders := convertTo2DIntSlice(ordersLines, ",")

	return instr, orders
}

func convertTo2DIntSlice(slice []string, sep string) (output [][]int) {
	for _, line := range slice {
		a := strings.Split(line, sep)
		b := []int{}
		for _, entry := range a {
			i, err := strconv.Atoi(entry)
			if err != nil {
				panic("Error convertig " + sep)
			}
			b = append(b, i)
		}
		output = append(output, b)
	}
	return
}
