package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed test_input.txt
var testInput string

//go:embed puzzle_input.txt
var puzzleInput string

type state int

const (
	COM state = iota
	LEFT
	RIGHT
	DISABLED
)

func main() {
	total := solvePartTwo(puzzleInput)

	fmt.Printf("Total %d\r\n", total)

}

func solvePartOne(in string) int {
	state := COM
	buffer := make([]byte, 0)

	left := 0
	right := 0
	total := 0

	for i := range in {
		r := in[i]

		buffer = append(buffer, r)

		switch state {
		case COM:
			// Check if the last part of the buffer matches mul(
			if bytes.HasSuffix(buffer, []byte("mul(")) {
				buffer = buffer[:0]
				state = LEFT
			}

		case LEFT:
			if r >= 48 && r <= 57 {
				left *= 10
				left += int(r - 48)
			} else if r == ',' {
				state = RIGHT
			} else {
				// Wrong char, reset
				state = COM
				left = 0
			}
		case RIGHT:
			if r >= 48 && r <= 57 {
				right *= 10
				right += int(r - 48)
			} else {
				if r == ')' {
					total += left * right
				}
				left = 0
				right = 0
				state = COM
			}
		}
	}

	return total
}

func solvePartTwo(in string) int {
	state := COM
	buffer := make([]byte, 0)

	left := 0
	right := 0
	total := 0

	for i := range in {
		r := in[i]

		switch state {
		case COM:
			buffer = append(buffer, r)
			// Check if the last part of the buffer matches mul(
			if bytes.HasSuffix(buffer, []byte("mul(")) {
				buffer = buffer[:0]
				state = LEFT
			} else if bytes.HasSuffix(buffer, []byte("don't()")) {
				state = DISABLED
				buffer = buffer[:0]
			}

		case DISABLED:
			buffer = append(buffer, r)
			if bytes.HasSuffix(buffer, []byte("do()")) {
				buffer = buffer[:0]
				state = COM
			}

		case LEFT:
			if r >= 48 && r <= 57 {
				left *= 10
				left += int(r - 48)
			} else if r == ',' {
				state = RIGHT
			} else {
				// Wrong char, reset
				state = COM
				left = 0
			}
		case RIGHT:
			if r >= 48 && r <= 57 {
				right *= 10
				right += int(r - 48)
			} else {
				if r == ')' {
					total += left * right
				}
				left = 0
				right = 0
				state = COM
			}
		}
	}

	return total
}
