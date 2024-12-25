package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var file = "input.txt"

func main() {
	mem, insts := readInputFile(file)

	zCount := 0

	for len(insts) != 0 {
		inst := insts[0]
		_, ok1 := mem[inst[0]]
		_, ok2 := mem[inst[2]]
		if ok1 && ok2 {
			if strings.Contains(inst[3], "z") {
				zCount++
			}
			calcInst(mem, inst)
			insts = insts[1:]
		} else {
			insts = append(insts[1:], inst)
		}
	}

	bin := ""
	for i := range zCount {
		key := fmt.Sprintf("z%02d", i)
		bin = strconv.Itoa(mem[key]) + bin
	}
	dec, _ := strconv.ParseInt(bin, 2, 64)
	fmt.Println(dec)

	// Part 2
	// this is only code to help. did the actual thing manually
	// if you visualize how each x and y get to their z its quite
	// straightforward to just check the spots with the wrong bit
	// for an error. This code gives those spots

	mem, insts = readInputFile(file)

	xCount := 0
	yCount := 0
	for k, _ := range mem {
		if strings.Contains(k, "x") {
			xCount++
		}
		if strings.Contains(k, "y") {
			yCount++
		}
	}

	binX := ""
	for i := range xCount {
		key := fmt.Sprintf("x%02d", i)
		binX = strconv.Itoa(mem[key]) + binX
	}
	binY := ""
	for i := range yCount {
		key := fmt.Sprintf("y%02d", i)
		binY = strconv.Itoa(mem[key]) + binY
	}
	decX, _ := strconv.ParseInt(binX, 2, 64)
	decY, _ := strconv.ParseInt(binY, 2, 64)
	decZ := decX + decY

	binZ := strconv.FormatInt(decZ, 2)

	fmt.Println(bin)
	fmt.Println(binZ)

	path := []string{}
	for i := range bin {
		if bin[i] != binZ[i] {
			fmt.Println(string(bin[i]), string(binZ[i]), len(bin)-i-1)
			p := recGetPath(insts, fmt.Sprintf("z%02d", len(bin)-i-1))
			p = removeDuplicateStr(p)
			path = append(path, p...)
		}
	}
	path = removeDuplicateStr(path)
}

func recGetPath(insts [][4]string, inst string) []string {
	path := []string{inst}
	if strings.Contains(inst, "x") || strings.Contains(inst, "y") {
		return []string{}
	}

	for _, i := range insts {
		if i[3] == inst {
			path = append(path, recGetPath(insts, i[0])...)
			path = append(path, recGetPath(insts, i[2])...)
		}
	}
	return path
}

func calcInst(mem map[string]int, inst [4]string) {
	switch inst[1] {
	case "AND":
		mem[inst[3]] = mem[inst[0]] & mem[inst[2]]
	case "OR":
		mem[inst[3]] = mem[inst[0]] | mem[inst[2]]
	case "XOR":
		mem[inst[3]] = mem[inst[0]] ^ mem[inst[2]]
	}
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func readInputFile(filename string) (map[string]int, [][4]string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	parts := strings.Split(string(f), "\n\n")

	mem := map[string]int{}
	lines := strings.Split(parts[0], "\n")
	for _, l := range lines {
		p := strings.Split(l, ": ")
		i, _ := strconv.Atoi(p[1])
		mem[p[0]] = i
	}

	insts := [][4]string{}
	lines = strings.Split(parts[1], "\n")
	for _, l := range lines {
		p := strings.Split(l, " -> ")
		i := strings.Split(p[0], " ")
		insts = append(insts, [4]string{i[0], i[1], i[2], p[1]})
	}
	return mem, insts
}
