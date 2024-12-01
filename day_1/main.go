package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed test_input.txt
var test_input string

//go:embed puzzle_input.txt
var puzzle_input string

func main() {
	a, b := parseInput(puzzle_input)

	fmt.Println("distance: %d", solvePartOne(a, b))
	fmt.Println("similarity: %d", solvePartTwo(a, b))
}

func parseInput(input string) ([]int, []int) {
	s := bufio.NewScanner(strings.NewReader(input))

	var listA []int
	var listB []int

	for s.Scan() {
		splitted := strings.Split(s.Text(), " ")
		var sanitized []string

		for i := range splitted {
			trimmed := strings.TrimSpace(splitted[i])
			if trimmed != "" {
				sanitized = append(sanitized, trimmed)
			}
		}

		a, err := strconv.ParseInt(sanitized[0], 10, 64)
		if err != nil {
			panic(err)
		}
		b, err := strconv.ParseInt(sanitized[1], 10, 64)
		if err != nil {
			panic(err)
		}

		listA = append(listA, int(a))
		listB = append(listB, int(b))
	}

	return listA, listB
}

func solvePartOne(a []int, b []int) int {
	slices.Sort(a)
	slices.Sort(b)

	totalDistance := 0
	for i := 0; i < len(a); i++ {
		distance := a[i] - b[i]
		if distance < 0 {
			distance *= -1
		}

		totalDistance += distance
	}

	return totalDistance
}

func solvePartTwo(a []int, b []int) int {
	occurenceMap := make(map[int]int)

	for i := range b {
		occurenceMap[b[i]] = occurenceMap[b[i]] + 1
	}

	similarity := 0
	for i := range a {
		similarity += a[i] * occurenceMap[a[i]]
	}

	return similarity
}
