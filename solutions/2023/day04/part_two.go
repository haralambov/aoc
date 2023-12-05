package day04

import (
	"fmt"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

func SecondPart() string {
	for _, line := range input.GetInputAsSlice(2023, 4) {
		if trimmedLine := strings.TrimSpace(line); len(trimmedLine) > 0 {
			processInputLine(trimmedLine)
		}
	}
	return fmt.Sprintf("%v", getScratchcardsCount())
}
