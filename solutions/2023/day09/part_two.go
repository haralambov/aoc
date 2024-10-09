package day09

import (
	"fmt"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

func SecondPart() string {
	sum := 0
	inputLines := input.GetInputAsSlice(2023, 9)
	for _, inputLine := range inputLines {
		numbers := strings.Split(inputLine, " ")
		parsedNumbers := parseNumbers(numbers)
		sum += extrapolateValue(parsedNumbers, true)
	}
	return fmt.Sprintf("%v", sum)
}
