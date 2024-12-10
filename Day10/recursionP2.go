package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInputFile("input.txt")

	count2 := 0

	for line := range input {
		for col, el := range input[line] {
			if el == 0 {
				routes := [][][2]int{}
				routes = append(routes, [][2]int{})
				routes[0] = append(routes[0], [2]int{line, col})

				routes = recTrailFinder(input, routes, 0)
				//fmt.Println(routes)

				correctRoutes := [][][2]int{}
				for _, route := range routes {
					if len(route) == 10 {
						correctRoutes = append(correctRoutes, route)
					}
				}
				count2 += len(correctRoutes)

			}
		}
	}
	fmt.Println(count2)
}

func recTrailFinder(input [][]int, routes [][][2]int, currentVal int) [][][2]int {
	if currentVal == 9 {
		return routes
	}

	for i, route := range routes {

		pos := route[len(route)-1]
		line := pos[0]
		col := pos[1]

		alreadyAppended := false

		if line > 0 {
			if input[line-1][col] == currentVal+1 {
				routes[i] = append(routes[i], [2]int{line - 1, col})
				alreadyAppended = true
			}
		}

		if line < len(input)-1 {
			if input[line+1][col] == currentVal+1 {

				if alreadyAppended {
					routeCpy := make([][2]int, len(route))
					copy(routeCpy, route)

					routeCpy = append(routeCpy, [2]int{line + 1, col})
					routes = append(routes, routeCpy)

				} else {
					routes[i] = append(routes[i], [2]int{line + 1, col})
					alreadyAppended = true
				}
			}
		}

		if col > 0 {
			if input[line][col-1] == currentVal+1 {

				if alreadyAppended {
					routeCpy := make([][2]int, len(route))
					copy(routeCpy, route)

					routeCpy = append(routeCpy, [2]int{line, col - 1})
					routes = append(routes, routeCpy)

				} else {
					routes[i] = append(routes[i], [2]int{line, col - 1})
					alreadyAppended = true
				}

			}
		}

		if col < len(input[0])-1 {
			if input[line][col+1] == currentVal+1 {

				if alreadyAppended {
					routeCpy := make([][2]int, len(route))
					copy(routeCpy, route)

					routeCpy = append(routeCpy, [2]int{line, col + 1})
					routes = append(routes, routeCpy)

				} else {
					routes[i] = append(routes[i], [2]int{line, col + 1})
					alreadyAppended = true
				}
			}
		}

	}

	return recTrailFinder(input, routes, currentVal+1)
}

func printMap(input [][]int) {
	fmt.Print("  ")
	for col := range input[0] {
		fmt.Print(col)
	}
	fmt.Println()
	fmt.Print("  ")
	for range input[0] {
		fmt.Print("-")
	}
	fmt.Println()
	for line := range input {
		fmt.Print(line)
		fmt.Print("|")
		for col := range input[line] {
			fmt.Print(strconv.Itoa(input[line][col]))
		}
		fmt.Println()
	}
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
