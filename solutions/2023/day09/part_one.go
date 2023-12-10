package day09

import (
	"fmt"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

func FirstPart() string {
	sum := 0
	inputLines := input.GetInputAsSlice(2023, 9)
	for _, inputLine := range inputLines {
		if trimmedLine := strings.TrimSpace(inputLine); len(trimmedLine) > 0 {
			numbers := strings.Split(trimmedLine, " ")
			parsedNumbers := parseNumbers(numbers)
			sum += extrapolateValue(parsedNumbers, false)
		}
	}
	return fmt.Sprintf("%v", sum)
}
