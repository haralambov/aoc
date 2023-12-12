package day05

import (
	"fmt"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

func SecondPart() string {
	inputLines := input.GetInputAsSlice(2015, 5)
	totalNiceStrings := 0
outer:
	for _, inputLine := range inputLines {
		trimmedLine := strings.TrimSpace(inputLine)
		if len(trimmedLine) == 0 {
			continue
		}

		containsPair := false
		for i := 1; i < len(trimmedLine); i++ {
			pair := string(trimmedLine[i-1]) + string(trimmedLine[i])
			stringToCheck := trimmedLine[i+1:]
			if strings.Contains(stringToCheck, pair) {
				containsPair = true
				break
			}
		}

		if !containsPair {
			continue
		}

		currentChar := trimmedLine[0]
		for i := 2; i < len(trimmedLine); i++ {
			if currentChar == trimmedLine[i] && currentChar == trimmedLine[i-2] {
				totalNiceStrings++
				continue outer
			} else {
				currentChar = trimmedLine[i-1]
			}
		}
	}
	return fmt.Sprintf("%v", totalNiceStrings)
}
