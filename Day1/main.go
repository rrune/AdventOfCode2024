package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/* var Slice1 = []int{3, 4, 2, 1, 3, 3}
var Slice2 = []int{4, 3, 5, 3, 9, 3} */

func main() {
	// read the input
	Slice1, Slice2 := readInputFile("input.txt")

	// Sort lists
	sort.Slice(Slice1, func(i, j int) bool {
		return Slice1[i] < Slice1[j]
	})
	sort.Slice(Slice2, func(i, j int) bool {
		return Slice2[i] < Slice2[j]
	})

	// calculate distance
	distanceSlice := []int{}
	for i := 0; i < len(Slice1); i++ {
		distanceSlice = append(distanceSlice, int(math.Abs(float64(Slice1[i]-Slice2[i]))))
	}

	// add up the distances
	addedUpDistance := 0
	for _, distance := range distanceSlice {
		addedUpDistance = addedUpDistance + distance
	}
	fmt.Println(addedUpDistance)

	// Part 2

	// calculate similarity
	occuranceMap := numberOfTimesInSlice(Slice2)
	similaritySlice := []int{}
	for _, id := range Slice1 {
		similaritySlice = append(similaritySlice, id*occuranceMap[id])
	}

	// add up the similarity
	addedUpSimilarity := 0
	for _, similarity := range similaritySlice {
		addedUpSimilarity = addedUpSimilarity + similarity
	}
	fmt.Println(addedUpSimilarity)
}

func readInputFile(filename string) (Slice1 []int, Slice2 []int) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}
	pairSlice := strings.Split(string(content), "\n")
	for _, pair := range pairSlice {
		leftSide, err := strconv.Atoi(pair[0:5])
		if err != nil {
			fmt.Println(err)
			panic("Error converting right side")
		}
		rightSide, err := strconv.Atoi(pair[8:13])
		if err != nil {
			fmt.Println(err)
			panic("Error converting left side")
		}
		Slice1 = append(Slice1, leftSide)
		Slice2 = append(Slice2, rightSide)
	}
	return
}

func numberOfTimesInSlice(slc []int) map[int]int {
	result := map[int]int{}
	for _, el := range slc {
		result[el] = result[el] + 1
	}
	return result
}
