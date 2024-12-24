package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	input := readInputFile("input.txt")

	all := []string{}
	for _, p := range input {
		if !slices.Contains(all, p[0]) {
			all = append(all, p[0])
		}
		if !slices.Contains(all, p[1]) {
			all = append(all, p[1])
		}
	}

	con := map[string][]string{}
	for _, p := range input {
		con[p[0]] = append(con[p[0]], p[1])
		con[p[1]] = append(con[p[1]], p[0])
	}

	pairs := [][3]string{}
	for _, a1 := range all {
		for _, a2 := range all {
			for _, a3 := range all {
				if slices.Contains(con[a1], a2) &&
					slices.Contains(con[a1], a3) &&
					slices.Contains(con[a2], a3) {

					t := []string{a1, a2, a3}
					sort.Strings(t)
					u := [3]string{t[0], t[1], t[2]}

					if !slices.Contains(pairs, u) {
						pairs = append(pairs, u)
					}

				}
			}
		}
	}

	count := 0
	for _, p := range pairs {
		t := false
		for _, e := range p {
			if e[0] == 't' {
				t = true
			}
		}
		if t {
			count++
		}
	}

	fmt.Println(count)

	cliques := [][]string{}
	P := map[string]bool{}
	for node := range con {
		P[node] = true
	}

	R := map[string]bool{}
	X := map[string]bool{}

	bronKerbosch(con, R, P, X, &cliques)

	longest := []string{}
	for _, c := range cliques {
		if len(c) > len(longest) {
			longest = c
		}
	}
	sort.Strings(longest)

	fmt.Println(strings.Join(longest, ","))
}

func bronKerbosch(graph map[string][]string, R, P, X map[string]bool, cliques *[][]string) {
	if len(P) == 0 && len(X) == 0 {
		clique := make([]string, 0, len(R))
		for node := range R {
			clique = append(clique, node)
		}
		*cliques = append(*cliques, clique)
		return
	}

	for node := range P {
		neighborNodes := graph[node]
		newR := copySet(R)
		newR[node] = true

		newP := intersectSets(P, neighborNodes)
		newX := intersectSets(X, neighborNodes)

		bronKerbosch(graph, newR, newP, newX, cliques)

		delete(P, node)
		X[node] = true
	}
}

func copySet(s map[string]bool) map[string]bool {
	copied := make(map[string]bool, len(s))
	for k := range s {
		copied[k] = true
	}
	return copied
}

func intersectSets(P map[string]bool, neighbors []string) map[string]bool {
	result := map[string]bool{}
	for _, neighbor := range neighbors {
		if _, found := P[neighbor]; found {
			result[neighbor] = true
		}
	}
	return result
}

func readInputFile(filename string) [][2]string {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading file")
	}
	lines := strings.Split(string(f), "\n")
	output := [][2]string{}
	for _, l := range lines {
		parts := strings.Split(l, "-")
		output = append(output, [2]string{parts[0], parts[1]})
	}
	return output
}
