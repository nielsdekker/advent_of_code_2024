package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed test_input.txt
var test_input string

//go:embed puzzle_input.txt
var puzzle_input string

func main() {
	in := parseInput(puzzle_input)
	safe := solvePartTwo(in)
	fmt.Printf("Safe: %d\r\n", safe)
}

func solvePartOne(input [][]int) int {
	safe := 0

	for i := range input {
		s, _, _ := isSafe(input[i], -1)

		if s {
			safe++
		}
	}

	return safe
}

func solvePartTwo(input [][]int) int {
	safe := 0

	for i := range input {
		s, c, n := isSafe(input[i], -1)

		if !s {
			safeC, _, _ := isSafe(input[i], c)
			safeN, _, _ := isSafe(input[i], n)

			if safeC || safeN {
				safe++
			}
		} else {
			safe++
		}
	}

	return safe
}

func isSafe(data []int, skipIndex int) (bool, int, int) {
	if len(data) < 2 {
		return true, -1, -1
	}

	c := 0
	n := 1

	if skipIndex == 0 {
		c = 1
		n = 2
	}
	if skipIndex == 1 {
		n = 2
	}

	isInc := data[c] < data[n]

	for n < len(data) {
		delta := data[c] - data[n]
		abs := delta
		if abs < 0 {
			abs *= -1
		}

		if abs == 0 || abs > 3 || (isInc && delta > 0) || (!isInc && delta < 0) {
			return false, c, n
		}

		c = n
		n++

		if n == skipIndex {
			n++
		}
	}

	return true, -1, -1
}

func parseInput(input string) [][]int {
	var out [][]int

	s := bufio.NewScanner(strings.NewReader(input))

	rootIndex := 0
	for s.Scan() {
		for len(out) < rootIndex+1 {
			out = append(out, []int{})
		}

		splitted := strings.Split(s.Text(), " ")

		for i := range splitted {
			trimmed := strings.TrimSpace(splitted[i])

			if trimmed != "" {
				asInt, err := strconv.ParseInt(trimmed, 10, 64)
				if err != nil {
					panic(err)
				}

				out[rootIndex] = append(out[rootIndex], int(asInt))
			}
		}

		rootIndex++
	}

	return out
}
