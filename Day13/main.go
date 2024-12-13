package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	A     [2]float64
	B     [2]float64
	Total [2]float64
}

var Offset float64 = 10000000000000

func main() {
	input := readInputFile("input.txt")

	var total1 float64 = 0
	var total2 float64 = 0

	for _, mach := range input {

		// asked my calculator
		n := -(mach.B[0]*mach.Total[1] - mach.B[1]*mach.Total[0]) / (mach.A[0]*mach.B[1] - mach.A[1]*mach.B[0])
		m := (mach.A[0]*mach.Total[1] - mach.A[1]*mach.Total[0]) / (mach.A[0]*mach.B[1] - mach.A[1]*mach.B[0])

		if isIntegral(n) && isIntegral(m) {
			total1 += n*3 + m
		}

		mach.Total[0] = mach.Total[0] + Offset
		mach.Total[1] = mach.Total[1] + Offset

		n = -(mach.B[0]*mach.Total[1] - mach.B[1]*mach.Total[0]) / (mach.A[0]*mach.B[1] - mach.A[1]*mach.B[0])
		m = (mach.A[0]*mach.Total[1] - mach.A[1]*mach.Total[0]) / (mach.A[0]*mach.B[1] - mach.A[1]*mach.B[0])

		if isIntegral(n) && isIntegral(m) {
			total2 += n*3 + m
		}
	}

	fmt.Println(int(total1))
	fmt.Println(int(total2))
}

func isIntegral(val float64) bool {
	return val == float64(int(val))
}

func readInputFile(filename string) []Machine {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}

	machines := strings.Split(string(f), "\n\n")
	output := []Machine{}

	for _, machine := range machines {

		lines := strings.Split(machine, "\n")

		aXs := lines[0][12:14]
		aYs := lines[0][18:20]

		bXs := lines[1][12:14]
		bYs := lines[1][18:20]

		t := strings.Split(lines[2], ",")
		tXs := strings.Split(t[0], "=")[1]
		tYs := t[1][3:]

		aX, _ := strconv.ParseFloat(aXs, 64)
		aY, _ := strconv.ParseFloat(aYs, 64)

		bX, _ := strconv.ParseFloat(bXs, 64)
		bY, _ := strconv.ParseFloat(bYs, 64)

		tX, _ := strconv.ParseFloat(tXs, 64)
		tY, _ := strconv.ParseFloat(tYs, 64)

		output = append(output, Machine{
			A:     [2]float64{aX, aY},
			B:     [2]float64{bX, bY},
			Total: [2]float64{tX, tY},
		})
	}

	return output
}
