package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Cheat struct {
	StartPos  [2]int
	CheatPos  [2]int
	StartTime int
}

func main() {
	input := readInputFile("input.txt")

	start := [2]int{}
	end := [2]int{}
	valid := [][2]int{}
	for row := range input {
		for col := range input[row] {
			if input[row][col] != '#' {
				valid = append(valid, [2]int{row, col})
			}
			if input[row][col] == 'S' {
				start = [2]int{row, col}
			}
			if input[row][col] == 'E' {
				end = [2]int{row, col}
			}
		}
	}
	invalid := [][2]int{}
	for row := range input {
		for col := range input[row] {
			if input[row][col] == '#' {
				invalid = append(invalid, [2]int{row, col})
			}
		}
	}

	_, path := bfs(valid, start, end)

	pathTimes := map[[2]int]int{}
	for i, p := range path {
		pathTimes[p] = i
	}

	neighbours := [][2]int{}
	for _, p := range path {
		ne := getNeighbours(invalid, p)
		for _, n := range ne {
			if !slices.Contains(neighbours, n) {
				neighbours = append(neighbours, n)
			}
		}
	}

	count := 0

	for _, n := range neighbours {
		ne := getNeighbours(valid, n)

		lowest := math.MaxInt
		hightest := 0

		for _, nei := range ne {
			if pathTimes[nei] < lowest {
				lowest = pathTimes[nei]
			}
			if pathTimes[nei] > hightest {
				hightest = pathTimes[nei]
			}
		}

		if hightest > lowest {
			total := (hightest - lowest - 2)

			if total >= 100 {
				count += 1
			}
		}
	}

	fmt.Println(count)

	count = 0

	for _, p1 := range path {
		for _, p2 := range path {
			distance := absDiffInt(p1[0], p2[0]) + absDiffInt(p1[1], p2[1])

			if distance <= 20 && pathTimes[p1] < pathTimes[p2] {
				total := absInt(pathTimes[p1] - pathTimes[p2])

				if total-distance >= 100 {
					count += 1
				}
			}
		}
	}
	fmt.Println(count)
}

func print(path [][2]int, ne [][2]int) {
	for row := range 140 {
		for col := range 140 {
			if slices.Contains(path, [2]int{row, col}) {
				fmt.Print(".")
			} else if slices.Contains(ne, [2]int{row, col}) {
				fmt.Print("+")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func bfs(grid [][2]int, start [2]int, end [2]int) (bool, [][2]int) {
	visited := [][2]int{}
	queue := [][2]int{}
	parent := map[[2]int][2]int{}
	queue = append(queue, start)
	found := false

	for _, u := range grid {
		parent[u] = [2]int{-1, -1}
	}
	parent[start] = start

	for !found {
		if len(queue) == 0 {
			return false, [][2]int{}
		}

		pos := queue[0]

		if pos == end {
			found = true
		}

		neighbours := getNeighbours(grid, pos)

		for _, n := range neighbours {
			if parent[n] == [2]int{-1, -1} {
				parent[n] = pos
			}

			if !slices.Contains(visited, n) {
				visited = append(visited, n)
				queue = append(queue, n)
			}
		}

		queue = queue[1:]
	}

	path := [][2]int{}
	current := end
	for current != start {
		path = append(path, parent[current])
		current = parent[current]
	}

	slices.Reverse(path)
	path = append(path, end)

	return true, path
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

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func readInputFile(filename string) [][]rune {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	lines := strings.Split(string(f), "\n")
	mp := [][]rune{}
	for _, l := range lines {
		mp = append(mp, []rune(l))
	}
	return mp
}
