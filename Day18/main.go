package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	mapSize = 71
	memSize = 1024
)

func main() {
	input := readInputFile("input.txt")
	corrupted := input[:memSize]

	uncorrupted := [][2]int{}
	for x := range mapSize {
		for y := range mapSize {
			pos := [2]int{x, y}
			if !slices.Contains(corrupted, pos) {
				uncorrupted = append(uncorrupted, pos)
			}
		}
	}

	_, cost := dijkstra(uncorrupted, [2]int{0, 0}, [2]int{mapSize - 1, mapSize - 1})

	fmt.Println(cost[[2]int{mapSize - 1, mapSize - 1}])

	// Part 2
	// just a bruteforce, to lazy to write a better dijstra that gets all every path

	for i := memSize; i < len(input); i++ {
		corrupted := input[:i]
		uncorrupted := [][2]int{}
		for x := range mapSize {
			for y := range mapSize {
				pos := [2]int{x, y}
				if !slices.Contains(corrupted, pos) {
					uncorrupted = append(uncorrupted, pos)
				}
			}
		}

		success, _ := dijkstra(uncorrupted, [2]int{0, 0}, [2]int{mapSize - 1, mapSize - 1})

		if !success {
			fmt.Print(input[i-1][0])
			fmt.Print(",")
			fmt.Print(input[i-1][1])
			break
		}
	}
}

func dijkstra(uncorrupted [][2]int, start [2]int, end [2]int) (bool, map[[2]int]int) {
	visited := [][2]int{}
	queue := [][2]int{}
	cost := map[[2]int]int{}
	queue = append(queue, start)
	found := false

	for _, u := range uncorrupted {
		cost[u] = math.MaxInt
	}
	cost[start] = 0

	for !found {
		if len(queue) == 0 {
			return false, cost
		}

		pos := queue[0]

		if pos == end {
			found = true
		}

		neighbours := getNeighbours(uncorrupted, pos)

		for _, n := range neighbours {
			if cost[n] > cost[pos]+1 {
				cost[n] = cost[pos] + 1
			}
			//cost[n] = cost[pos] + 1

			if !slices.Contains(visited, n) {
				visited = append(visited, n)
				queue = append(queue, n)
			}
		}

		queue = queue[1:]
	}

	return true, cost
	//fmt.Println(cost[[2]int{mapSize - 1, mapSize - 1}])
}

func getNeighbours(mp [][2]int, pos [2]int) [][2]int {
	output := [][2]int{}

	if slices.Contains(mp, [2]int{pos[0] - 1, pos[1]}) {
		output = append(output, [2]int{pos[0] - 1, pos[1]})
	}

	if slices.Contains(mp, [2]int{pos[0] + 1, pos[1]}) {
		output = append(output, [2]int{pos[0] + 1, pos[1]})
	}

	if slices.Contains(mp, [2]int{pos[0], pos[1] - 1}) {
		output = append(output, [2]int{pos[0], pos[1] - 1})
	}

	if slices.Contains(mp, [2]int{pos[0], pos[1] + 1}) {
		output = append(output, [2]int{pos[0], pos[1] + 1})
	}

	return output
}

func printMap(uncorrupted [][2]int, visited [][2]int) {
	for y := range mapSize {
		for x := range mapSize {
			pos := [2]int{x, y}
			if slices.Contains(visited, pos) {
				fmt.Print("O")
			} else if slices.Contains(uncorrupted, pos) {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func readInputFile(filename string) [][2]int {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	output := [][2]int{}

	lines := strings.Split(string(f), "\n")
	for _, line := range lines {
		p := strings.Split(line, ",")

		p1, _ := strconv.Atoi(p[0])
		p2, _ := strconv.Atoi(p[1])

		output = append(output, [2]int{p1, p2})
	}
	return output
}
