/* package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

type State struct {
	x         int
	y         int
	direction int
	cost      int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func readMaze(fileName string) ([]string, Point, Point) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var maze []string
	var start, end Point
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		maze = append(maze, line)
		for x := 0; x < len(line); x++ {
			if line[x] == 'S' {
				start = Point{x, y}
			}
			if line[x] == 'E' {
				end = Point{x, y}
			}
		}
	}
	return maze, start, end
}

func isValid(maze []string, x, y int) bool {
	return x >= 0 && y >= 0 && y < len(maze) && x < len(maze[y]) && maze[y][x] != '#'
}

func bfs(maze []string, start, end Point) (int, map[Point]bool) {
	directions := []Point{
		{0, -1}, // North
		{1, 0},  // East
		{0, 1},  // South
		{-1, 0}, // West
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &State{start.x, start.y, 1, 0})

	visited := make(map[[3]int]int)
	bestPathTiles := make(map[Point]bool)
	parent := make(map[Point][]State)

	for pq.Len() > 0 {
		state := heap.Pop(pq).(*State)

		if state.x == end.x && state.y == end.y {
			bestPathTiles[state] = true
			for _, parentState := range parent[state] {
				bestPathTiles[parentState] = true
			}
			return state.cost, bestPathTiles
		}

		if v, exists := visited[[3]int{state.x, state.y, state.direction}]; exists && v <= state.cost {
			continue
		}
		visited[[3]int{state.x, state.y, state.direction}] = state.cost

		// forward
		nextX := state.x + directions[state.direction].x
		nextY := state.y + directions[state.direction].y
		if isValid(maze, nextX, nextY) {
			heap.Push(pq, &State{nextX, nextY, state.direction, state.cost + 1})
		}

		// Rotate clockwise (90 degrees)
		newDirection := (state.direction + 1) % 4
		heap.Push(pq, &State{state.x, state.y, newDirection, state.cost + 1000})

		// Rotate counter-clockwise (90 degrees)
		newDirection = (state.direction + 3) % 4
		heap.Push(pq, &State{state.x, state.y, newDirection, state.cost + 1000})
	}

	return -1
}

func main() {
	maze, start, end := readMaze("input.txt")
	minScore := bfs(maze, start, end)
	fmt.Println(minScore)
}
*/