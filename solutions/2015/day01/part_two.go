package day01

import (
	"fmt"

	"github.com/haralambov/aoc/lib/input"
)

func SecondPart() string {
	input := input.GetInputAsString(2015, 1)
	currentFloor := 0
	for i, char := range input {
		if char == '(' {
			currentFloor++
		} else {
			currentFloor--
		}
		if currentFloor == -1 {
			return fmt.Sprintf("%v", i+1)
		}
	}
	return "0"
}
