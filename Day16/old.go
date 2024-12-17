package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	UP    [2]int = [2]int{-1, 0}
	DOWN  [2]int = [2]int{1, 0}
	LEFT  [2]int = [2]int{0, -1}
	RIGHT [2]int = [2]int{0, 1}
	//DIRS  [][2]int = [][2]int{UP, DOWN, LEFT, RIGHT}
)

type Path struct {
	P [][2]int
	R int
}

type Pos struct {
	P [2]int
	D [2]int
}

func main() {
	input := readInputFile("input.txt")

	SPos := [2]int{}
	for line := range input {
		for col := range input[line] {
			if input[line][col] == 'S' {
				SPos = [2]int{line, col}
			}
		}
	}

	test, cul := recMapTrav(input, SPos, RIGHT, Path{})
	fmt.Println(test, cul)

}

func recMapTrav(mp [][]rune, pos [2]int, direction [2]int, path Path) (int, bool) {
	//printMap(mp, pos)
	dirs := getDirs(direction)

	checkPos := []Pos{}
	for _, dir := range dirs {
		checkPos = append(checkPos, Pos{[2]int{pos[0] + dir[0], pos[1] + dir[1]}, dir})
	}

	pathNums := []int{}
	culdesacNow := true
	culdesacUp := false
	for _, p := range checkPos {
		if !slices.Contains(path.P, p.P) {
			if mp[p.P[0]][p.P[1]] == '.' {
				culdesacNow = false
				newPath := deepCopy(path)
				if p.D != direction {
					newPath.R += 1
				}
				newPath.P = append(newPath.P, pos)

				num := 0
				num, culdesacUp = recMapTrav(mp, p.P, p.D, newPath)
				pathNums = append(pathNums, num)
			}

			if mp[p.P[0]][p.P[1]] == 'E' {
				path.P = append(path.P, pos)
				r := strconv.Itoa(path.R)
				l := strconv.Itoa(len(path.P))

				for len(l) < 3 {
					l = "0" + l
				}

				i, _ := strconv.Atoi(r + l)

				return i, false
			}
		}
	}

	if culdesacNow == true || culdesacUp == true {
		//return math.MaxInt, true
	}

	//fmt.Println(path)

	smallest := math.MaxInt
	for _, num := range pathNums {
		if num < smallest {
			smallest = num
		}
	}

	return smallest, false
}

func getDirs(dir [2]int) [][2]int {
	if dir == UP {
		return [][2]int{UP, LEFT, RIGHT}
	}
	if dir == DOWN {
		return [][2]int{DOWN, LEFT, RIGHT}
	}
	if dir == LEFT {
		return [][2]int{UP, LEFT, DOWN}
	}
	if dir == RIGHT {
		return [][2]int{UP, DOWN, RIGHT}
	}
	return [][2]int{}
}

func deepCopy(path Path) Path {
	newp := [][2]int{}
	for _, p := range path.P {
		newp = append(newp, [2]int{p[0], p[1]})
	}
	return Path{newp, path.R}
}

func printMap(input [][]rune, pos [2]int) {
	newMp := [][]rune{}
	for _, l := range input {
		cpL := make([]rune, len(l))
		copy(cpL, l)
		newMp = append(newMp, cpL)
	}

	newMp[pos[0]][pos[1]] = '@'
	/* 			fmt.Print(" ")
	for col := range input[0] {
		fmt.Print(col)
	}
	fmt.Println() */
	for line := range newMp {
		//fmt.Print(line)
		for col := range newMp[line] {
			fmt.Print(string(newMp[line][col]))
		}
		fmt.Println()
	}
}

func readInputFile(filename string) [][]rune {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	lines := strings.Split(string(f), "\n")
	output := [][]rune{}
	for _, line := range lines {
		output = append(output, []rune(line))
	}
	return output
}
