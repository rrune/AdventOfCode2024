package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/* var input = []string{
	"x",
	"mul(2,4)%&",
	"mul[3,7]!@^do_not_",
	"mul(5,5)+",
	"mul(32,64]then(",
	"mul(11,8)",
	"mul(8,5))]",
} */

var regexpStr = "mul\\([0-9]{1,3},[0-9]{1,3}\\)"

func main() {
	input := readInputFile("input.txt")

	r, err := regexp.Compile(regexpStr)
	if err != nil {
		panic("Error at regexp")
	}

	do := true

	results := []int{}
	for _, str := range input {

		if r.MatchString(str) && do {
			rl := strings.Split(str, ",")
			right := rl[0][4:]
			left := strings.Split(rl[1], ")")[0]

			rightInt, err := strconv.Atoi(right)
			if err != nil {
				panic("Error at right")
			}

			leftInt, err := strconv.Atoi(left)
			if err != nil {
				panic("Error at left")
			}

			results = append(results, (rightInt * leftInt))
		}

		if strings.Contains(str, "do()") {
			do = true
		}
		if strings.Contains(str, "don't()") {
			do = false
		}

	}
	result := 0
	for _, el := range results {
		result += el
	}

	fmt.Println(result)
}

func readInputFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	output := strings.Split(string(content), "mul")
	for i := 1; i < len(output); i++ {
		output[i] = "mul" + output[i]
	}
	return output
}
