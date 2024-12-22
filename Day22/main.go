package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Seq struct {
	a, b, c, d int
}

func main() {
	input := readInputFile("input.txt")

	totalP1 := 0
	secrets := [][]int{}
	for _, s := range input {
		sec := []int{s}
		for range 2000 {
			s = calcNextSecret(s)
			sec = append(sec, s%10)
		}
		totalP1 += s
		secrets = append(secrets, sec)
	}
	fmt.Println(totalP1)

	diffs := [][]int{}
	for _, s := range secrets {
		d := []int{-1}
		for i := 1; i < len(s); i++ {
			d = append(d, s[i]-s[i-1])
		}
		diffs = append(diffs, d)
	}

	seqs := map[Seq]int{}

	for i, sec := range secrets {
		seen := []Seq{}
		for j := 4; j < len(sec); j++ {
			seq := Seq{diffs[i][j-3], diffs[i][j-2], diffs[i][j-1], diffs[i][j]}
			if !slices.Contains(seen, seq) {
				seqs[seq] += sec[j]
				seen = append(seen, seq)
			}
		}

	}

	biggestVal := 0
	biggestKey := Seq{}
	for k, v := range seqs {
		if v > biggestVal {
			biggestVal = v
			biggestKey = k
		}
	}

	fmt.Println(biggestKey, biggestVal)
}

func calcNumAtSeq(secrets []int, diffs []int, seq Seq) int {
	for i := 4; i < len(diffs); i++ {
		currseq := Seq{diffs[i-3], diffs[i-2], diffs[i-1], diffs[i]}

		if currseq == seq {
			return secrets[i]
		}
	}
	return 0
}

func calcNextSecret(secret int) int {
	temp := secret * 64
	secret = temp ^ secret
	secret = secret % 16777216

	temp = secret / 32
	secret = temp ^ secret
	secret = secret % 16777216

	temp = secret * 2048
	secret = temp ^ secret
	secret = secret % 16777216

	return secret
}

func readInputFile(filename string) []int {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}
	lines := strings.Split(string(f), "\n")

	output := []int{}
	for _, l := range lines {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic("Error converting line")
		}
		output = append(output, i)
	}
	return output
}
