package day01

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/haralambov/aoc/lib/input"
)

func FirstPart() string {
	inputLines := input.GetInputAsSlice(2023, 1)
	re := regexp.MustCompile(`[0-9]`)
	var sum int
	for _, line := range inputLines {
		matches := re.FindAllString(line, -1)
		number := matches[0] + matches[len(matches)-1]
		convertedNumber, _ := strconv.ParseInt(number, 10, 32)
		sum += int(convertedNumber)
	}
	return fmt.Sprintf("%v", sum)
}
