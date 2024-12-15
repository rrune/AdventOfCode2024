package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	mp, instr := readInputFile("input.txt")

	robotPos := [2]int{}
	for line := range mp {
		for col := range mp[line] {
			if mp[line][col] == '@' {
				robotPos = [2]int{line, col}
			}
		}
	}

	for _, inst := range instr {
		mp, robotPos, _ = recRobotP1(mp, inst, robotPos)
	}
	//printMap(mp)

	count := 0
	for line := range mp {
		for col := range mp[line] {
			if mp[line][col] == 'O' {
				count += line*100 + col
			}
		}
	}
	fmt.Println(count)

	// Part 2

	mp, instr = readInputFile("input.txt")

	newmp := [][]rune{}
	for line := range mp {
		newmp = append(newmp, []rune{})
		for col := range mp[line] {
			if mp[line][col] == 'O' {
				newmp[line] = append(newmp[line], '[', ']')
			} else if mp[line][col] == '@' {
				newmp[line] = append(newmp[line], '@', '.')
			} else {
				newmp[line] = append(newmp[line], mp[line][col], mp[line][col])
			}
		}
	}
	mp = newmp

	robotPos = [2]int{}
	for line := range mp {
		for col := range mp[line] {
			if mp[line][col] == '@' {
				robotPos = [2]int{line, col}
			}
		}
	}

	for _, inst := range instr {
		mp, robotPos, _ = recRobotP2(mp, inst, robotPos)
	}
	//printMap(mp)

	count = 0
	for line := range mp {
		for col := range mp[line] {
			if mp[line][col] == '[' {
				count += line*100 + col
			}
		}
	}
	fmt.Println(count)

}

func recRobotP1(mp [][]rune, inst rune, pos [2]int) ([][]rune, [2]int, bool) {
	direction := [2]int{}
	switch inst {
	case '^':
		direction = [2]int{-1, 0}
	case '>':
		direction = [2]int{0, 1}
	case 'v':
		direction = [2]int{1, 0}
	case '<':
		direction = [2]int{0, -1}
	}

	checkPos := [2]int{pos[0] + direction[0], pos[1] + direction[1]}

	if mp[checkPos[0]][checkPos[1]] == '.' {
		mp[checkPos[0]][checkPos[1]] = mp[pos[0]][pos[1]]
		mp[pos[0]][pos[1]] = '.'
		return mp, checkPos, true
	}

	if mp[checkPos[0]][checkPos[1]] == '#' {
		return mp, pos, false
	}

	ok := false
	mp, _, ok = recRobotP1(mp, inst, checkPos)
	if ok {
		mp[checkPos[0]][checkPos[1]] = mp[pos[0]][pos[1]]
		mp[pos[0]][pos[1]] = '.'
		return mp, checkPos, true
	}
	return mp, pos, false

}

func recRobotP2(mp [][]rune, inst rune, pos [2]int) ([][]rune, [2]int, bool) {
	direction := [2]int{}
	switch inst {
	case '^':
		direction = [2]int{-1, 0}
	case '>':
		direction = [2]int{0, 1}
	case 'v':
		direction = [2]int{1, 0}
	case '<':
		direction = [2]int{0, -1}
	}

	checkPos := [2]int{pos[0] + direction[0], pos[1] + direction[1]}

	if mp[checkPos[0]][checkPos[1]] == '.' {
		mp[checkPos[0]][checkPos[1]] = mp[pos[0]][pos[1]]
		mp[pos[0]][pos[1]] = '.'
		return mp, checkPos, true
	}

	if mp[checkPos[0]][checkPos[1]] == '#' {
		return mp, pos, false
	}

	if inst == '^' || inst == 'v' {
		mpCopy := [][]rune{}
		for _, line := range mp {
			lCpy := make([]rune, len(line))
			copy(lCpy, line)
			mpCopy = append(mpCopy, lCpy)
		}

		if mp[checkPos[0]][checkPos[1]] == '[' {
			ok1 := false
			ok2 := false
			mp, _, ok1 = recRobotP2(mp, inst, checkPos)
			mp, _, ok2 = recRobotP2(mp, inst, [2]int{checkPos[0], checkPos[1] + 1})

			if ok1 && ok2 {
				mp[checkPos[0]][checkPos[1]] = mp[pos[0]][pos[1]]
				mp[pos[0]][pos[1]] = '.'
				return mp, checkPos, true
			} else {
				return mpCopy, pos, false
			}
		} else if mp[checkPos[0]][checkPos[1]] == ']' {
			ok1 := false
			ok2 := false
			mp, _, ok1 = recRobotP2(mp, inst, checkPos)
			mp, _, ok2 = recRobotP2(mp, inst, [2]int{checkPos[0], checkPos[1] - 1})

			if ok1 && ok2 {
				mp[checkPos[0]][checkPos[1]] = mp[pos[0]][pos[1]]
				mp[pos[0]][pos[1]] = '.'
				return mp, checkPos, true
			} else {
				return mpCopy, pos, false
			}
		}

	}

	ok := false
	mp, _, ok = recRobotP2(mp, inst, checkPos)
	if ok {
		mp[checkPos[0]][checkPos[1]] = mp[pos[0]][pos[1]]
		mp[pos[0]][pos[1]] = '.'
		return mp, checkPos, true
	}
	return mp, pos, false

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

func readInputFile(filename string) ([][]rune, []rune) {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading input")
	}

	parts := strings.Split(string(f), "\n\n")

	lines := strings.Split(parts[0], "\n")
	wh := [][]rune{}
	for _, line := range lines {
		wh = append(wh, []rune(line))
	}

	instr := []rune{}
	for _, ins := range parts[1] {
		if ins != '\n' {
			instr = append(instr, ins)
		}
	}

	return wh, instr
}
