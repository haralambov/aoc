package day05

import (
	"fmt"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

func FirstPart() string {
	inputLines := input.GetInputAsSlice(2015, 5)
	totalNiceStrings := 0
	forbiddenStrings := []string{"ab", "cd", "pq", "xy"}
	vowels := map[rune]int{
		'a': 0,
		'e': 0,
		'i': 0,
		'o': 0,
		'u': 0,
	}
outer:
	for _, inputLine := range inputLines {
		for _, forbiddenString := range forbiddenStrings {
			if strings.Contains(inputLine, forbiddenString) {
				continue outer
			}
		}

		vowelsCount := 0
		for _, char := range inputLine {
			if _, ok := vowels[char]; ok {
				vowelsCount++
			}
		}
		if vowelsCount < 3 {
			continue
		}

		currentChar := inputLine[0]
		for i := 1; i < len(inputLine); i++ {
			if currentChar == inputLine[i] {
				totalNiceStrings++
				continue outer
			} else {
				currentChar = inputLine[i]
			}
		}
	}
	return fmt.Sprintf("%v", totalNiceStrings)
}
