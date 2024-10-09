package day04

import (
	"fmt"

	"github.com/haralambov/aoc/lib/input"
)

func SecondPart() string {
	for _, line := range input.GetInputAsSlice(2023, 4) {
		processInputLine(line)
	}
	return fmt.Sprintf("%v", getScratchcardsCount())
}
