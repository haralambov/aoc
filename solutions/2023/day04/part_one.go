package day04

import (
	"fmt"

	"github.com/haralambov/aoc/lib/input"
)

func FirstPart() string {
	sum := 0
	for _, line := range input.GetInputAsSlice(2023, 4) {
		sum += calcPointsInScratchcard(line)
	}
	return fmt.Sprintf("%v", sum)
}
