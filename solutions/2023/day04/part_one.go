package day04

import (
	"fmt"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

func FirstPart() string {
	sum := 0
	for _, line := range input.GetInputAsSlice(2023, 4) {
		if trimmedLine := strings.TrimSpace(line); len(trimmedLine) > 0 {
			sum += calcPointsInScratchcard(trimmedLine)
		}
	}
	return fmt.Sprintf("%v", sum)
}
