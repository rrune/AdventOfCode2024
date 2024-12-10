// Truly horrible code again. Works I guess. Pretty slow

package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := readInputFile("input.txt")
	decrompessed := []string{}

	// decormpessing
	for i, character := range input {
		var toAppend string
		if i%2 == 0 {
			toAppend = strconv.Itoa(i / 2)
		} else {
			toAppend = "."
		}
		for range character {
			decrompessed = append(decrompessed, toAppend)
		}
	}

	// copy for part 2
	decrompessedCopy := make([]string, len(decrompessed))
	copy(decrompessedCopy, decrompessed)

	// switching
	dotCount := 0
	for _, char := range decrompessed {
		if char == "." {
			dotCount++
		}
	}

	for range dotCount {
		switchDotWithNum(decrompessed)
	}

	// checksum
	checksum := 0
	for i, char := range decrompessed {
		if char != "." {
			charInt, err := strconv.Atoi(char)
			if err != nil {
				panic("Error converting checksum 1")
			}
			checksum += i * charInt
		}
	}

	fmt.Println(checksum)

	// Part 2
	decrompessed = decrompessedCopy

	// group the files
	files := [][]string{{}}
	last := decrompessed[0]
	for _, char := range decrompessed {
		if char == last {
			files[len(files)-1] = append(files[len(files)-1], char)
		} else {
			files = append(files, []string{char})
			last = char
		}
	}

	// sort
	files = recSwitchFiles(files, len(files)-1)

	result := []string{}
	for _, file := range files {
		result = append(result, file...)
	}

	// checksum
	checksum2 := 0
	for i, char := range result {
		if char != "." {
			charInt, err := strconv.Atoi(char)
			if err != nil {
				panic("Error converting checksum 2")
			}
			checksum2 += i * charInt
		}
	}

	fmt.Println(checksum2)
}

func switchDotWithNum(decrompessed []string) {
	for i, char := range decrompessed {
		if char == "." {
			for j := len(decrompessed) - 1; j >= 0; j-- {
				if decrompessed[j] != "." {
					decrompessed[i] = decrompessed[j]
					decrompessed[j] = "."
					return
				}
			}
		}
	}
}

func recSwitchFiles(files [][]string, i int) [][]string {
	if i != 0 {
		containsDot := false
		for _, char := range files[i] {
			if strings.Contains(char, ".") {
				containsDot = true
			}
		}
		if containsDot {
			return recSwitchFiles(files, i-1)
		}

		for j, file := range files {

			onlyDots := true
			for _, str := range file {
				for _, char := range str {
					if char != '.' {
						onlyDots = false
					}
				}
			}

			if onlyDots {
				if len(file) >= len(files[i]) && j < i {
					dotPart1 := file[:len(files[i])]
					dotPart2 := file[len(files[i]):]

					files[j] = files[i]
					files[i] = dotPart1
					if len(dotPart2) != 0 {
						files = slices.Insert(files, j+1, dotPart2)
						return recSwitchFiles(files, i)
					} else {
						return recSwitchFiles(files, i-1)
					}

				}
			}
		}
		return recSwitchFiles(files, i-1)
	}
	return files
}

func readInputFile(filename string) []int {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	charactersStr := strings.Split(string(content), "")
	output := []int{}
	for _, character := range charactersStr {
		i, err := strconv.Atoi(character)
		if err != nil {
			panic("Error converting character")
		}
		output = append(output, i)
	}

	return output
}
