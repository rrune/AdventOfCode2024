package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	towels, designs := readInputFile("input.txt")

	count1 := 0
	count2 := 0
	for _, d := range designs {
		r := recDesign(d, towels, map[string]int{})
		if r > 0 {
			count1++
			count2 += r
		}
	}
	fmt.Println(count1)
	fmt.Println(count2)
}

func recDesign(design string, towels []string, memo map[string]int) int {
	if n, ok := memo[design]; ok {
		return n
	}
	if design == "" {
		return 1
	}

	res := 0
	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			res += recDesign(design[len(towel):], towels, memo)
		}
	}
	memo[design] = res
	return res

}

func readInputFile(filename string) ([]string, []string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	parts := strings.Split(string(f), "\n\n")

	towels := strings.Split(parts[0], ",")
	designs := strings.Split(parts[1], "\n")

	for i := 1; i < len(towels); i++ {
		towels[i] = towels[i][1:]
	}

	return towels, designs
}
