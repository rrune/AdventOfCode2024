package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInputFile("input.txt")

	keys := [][]int{}
	locks := [][]int{}

	for _, e := range input {
		key := true
		for col := range e {
			if e[col][0] == '#' {
				key = false
			}
		}

		heightMap := []int{}

		for _, col := range e {
			height := len(strings.ReplaceAll(string(col), ".", "")) - 1
			heightMap = append(heightMap, height)
		}

		if key {
			keys = append(keys, heightMap)
		} else {
			locks = append(locks, heightMap)
		}
	}

	count := 0
	for _, key := range keys {
		for _, lock := range locks {

			fits := true
			for i := range len(key) {
				sum := key[i] + lock[i]
				if sum > 5 {
					fits = false
				}
			}
			if fits {
				count++
			}

		}
	}
	fmt.Println(count)
}

func readInputFile(filename string) [][][]rune {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}
	kl := strings.Split(string(f), "\n\n")

	output := [][][]rune{}
	for _, e := range kl {
		lines := strings.Split(e, "\n")
		m := [][]rune{}
		for _, line := range lines {
			m = append(m, []rune(line))
		}

		mr := make([][]rune, len(m[0]))

		for line := range m {
			for col := range m[line] {
				mr[col] = append(mr[col], m[line][col])
			}
		}

		output = append(output, mr)
	}
	return output
}
