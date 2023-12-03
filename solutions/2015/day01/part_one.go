package day01

import (
	"fmt"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

func FirstPart() string {
	input := input.GetInputAsString(2015, 1)
	leftParenthesisCount := strings.Count(input, "(")
	rightParenthesisCount := strings.Count(input, ")")
	return fmt.Sprintf("%v", leftParenthesisCount-rightParenthesisCount)
}
