package day08

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

var (
	total    = 0
	hexMatch = regexp.MustCompile(`\\x[a-zA-Z0-9]{2}`)
)

func FirstPart() string {
	return solve(false)
}

func solve(encodeFirst bool) string {
	inputLines := input.GetInputAsSlice(2015, 8)
	for _, inputLine := range inputLines {
		if encodeFirst {
			inputLine = encodeInputLine(inputLine)
		}
		total += len(inputLine) - calculateNumberOfCharacters(inputLine)
	}
	return fmt.Sprintf("%d", total)
}

func calculateNumberOfCharacters(line string) int {
	line = line[1 : len(line)-1]
	line = strings.ReplaceAll(line, "\\\\", "a")
	line = strings.ReplaceAll(line, "\\\"", "a")
	line = hexMatch.ReplaceAllString(line, "a")
	return len(line)
}

func encodeInputLine(inputLine string) string {
	inputLine = strings.ReplaceAll(inputLine, "\\", "\\\\")
	inputLine = strings.ReplaceAll(inputLine, "\"", "\\\"")
	return "\"" + inputLine + "\""
}
