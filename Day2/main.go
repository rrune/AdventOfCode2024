package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

/* var levelsSlice = [][]int{
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
} */

func main() {
	levelsSlice := readInputFile("input.txt")

	safeCount := 0
	for _, levels := range levelsSlice {
		if isLevelSliceSafe(levels) {
			safeCount = safeCount + 1
		} else {
			validWithOneRemoved := false
			for i := 0; i < len(levels); i++ {
				levelsCopy := append([]int{}, levels...)

				if isLevelSliceSafe(remove(levelsCopy, i)) {
					validWithOneRemoved = true
				}
			}
			if validWithOneRemoved {
				safeCount = safeCount + 1
			}
		}
	}

	fmt.Println(safeCount)
}

func isLevelSliceSafe(levels []int) (safe bool) {
	for i := 1; i < len(levels); i++ {
		diff := math.Abs(float64(levels[i] - levels[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}
	}

	ascending := slices.IsSorted(levels)

	reversedLevels := slices.Clone(levels)
	slices.Reverse(reversedLevels)
	decending := slices.IsSorted(reversedLevels)

	if ascending || decending {
		return true
	}

	return false
}

func readInputFile(filename string) (output [][]int) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		levelsStringSlice := strings.Split(line, " ")
		levelsIntSlice := []int{}
		for _, levelsStr := range levelsStringSlice {
			levelsInt, err := strconv.Atoi(levelsStr)
			if err != nil {
				fmt.Println(err)
				panic("Error converting level string")
			}
			levelsIntSlice = append(levelsIntSlice, levelsInt)
		}
		output = append(output, levelsIntSlice)
	}
	return
}

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
