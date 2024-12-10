package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var alphabat string = "+*|"

func main() {
	input := readInputFile("input.txt")

	max := 0
	for _, entry := range input {
		if len(entry)-1 > max {
			max = len(entry) - 1
		}
	}

	combs := generateCombs(max)

	sum := 0

	for _, entry := range input {

		correct := false
		key := entry[0]
		values := entry[1:]

		for _, comb := range combs[len(values)-1] {

			total := values[0]
			for i := range len(values) - 1 {

				switch comb[i] {
				case '*':
					total = total * values[i+1]

				case '+':
					total = total + values[i+1]

				case '|':
					temp, err := strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(values[i+1]))
					if err != nil {
						panic("Error convertig |")
					}
					total = temp
				}

			}

			if total == key {
				correct = true
			}

		}

		if correct {
			sum += key
		}

	}
	fmt.Println(sum)
}

func generateCombs(length int) map[int][]string {
	combs := map[int][]string{
		1: {"*", "+", "|"},
	}

	for i := range length - 1 {
		i++

		for _, el := range combs[i] {

			combs[i+1] = append(combs[i+1], el+"+")
			combs[i+1] = append(combs[i+1], el+"*")
			combs[i+1] = append(combs[i+1], el+"|")
		}

	}

	return combs
}

func readInputFile(filename string) [][]int {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	lines := strings.Split(string(content), "\n")
	output := [][]int{}

	for _, line := range lines {
		parts := strings.Split(line, ":")
		key, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("Error convertig key")
		}
		valuesStr := strings.Split(parts[1][1:], " ")

		sl := []int{key}
		for _, valueStr := range valuesStr {
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				panic("Error convertig value")
			}
			sl = append(sl, value)
		}
		output = append(output, sl)

	}
	return output
}
