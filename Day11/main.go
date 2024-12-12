package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PrevNum struct {
	Num    int
	Blinks int
}

func main() {
	input := readInputFile("input.txt")

	var cache = make(map[PrevNum]int)
	total := 0
	for _, num := range input {
		total += recStone(num, 75, cache)
	}
	fmt.Println(total)
}

func recStone(num int, blinks int, cache map[PrevNum]int) int {
	if blinks == 0 {
		return 1
	}

	prev := cache[PrevNum{Num: num, Blinks: blinks}]
	if prev != 0 {
		return prev
	}

	numStr := strconv.Itoa(num)

	var total int

	if num == 0 {
		total = recStone(1, blinks-1, cache)
	} else if len(numStr)%2 == 0 {
		numLeftStr := numStr[:len(numStr)/2]
		numRightStr := numStr[len(numStr)/2:]

		numLeft, err := strconv.Atoi(numLeftStr)
		if err != nil {
			fmt.Println(err)
			panic("Error converting numLeft")
		}
		numRight, err := strconv.Atoi(numRightStr)
		if err != nil {
			fmt.Println(err)
			panic("Error converting numRight")
		}

		total = recStone(numLeft, blinks-1, cache) + recStone(numRight, blinks-1, cache)
	} else {
		total = recStone(num*2024, blinks-1, cache)
	}

	cache[PrevNum{Num: num, Blinks: blinks}] = total

	return total
}

func readInputFile(filename string) []int {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	parts := strings.Split(string(content), " ")
	output := []int{}
	for _, part := range parts {
		i, err := strconv.Atoi(part)
		if err != nil {
			panic("Error converting line")
		}
		output = append(output, i)
	}
	return output
}
