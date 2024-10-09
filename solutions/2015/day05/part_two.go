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
		containsPair := false
		for i := 1; i < len(inputLine); i++ {
			pair := string(inputLine[i-1]) + string(inputLine[i])
			stringToCheck := inputLine[i+1:]
			if strings.Contains(stringToCheck, pair) {
				containsPair = true
				break
			}
		}

		if !containsPair {
			continue
		}

		currentChar := inputLine[0]
		for i := 2; i < len(inputLine); i++ {
			if currentChar == inputLine[i] && currentChar == inputLine[i-2] {
				totalNiceStrings++
				continue outer
			} else {
				currentChar = inputLine[i-1]
			}
		}
	}
	return fmt.Sprintf("%v", totalNiceStrings)
}
