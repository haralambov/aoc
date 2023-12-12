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
		trimmedLine := strings.TrimSpace(inputLine)
		if len(trimmedLine) == 0 {
			continue
		}

		for _, forbiddenString := range forbiddenStrings {
			if strings.Contains(trimmedLine, forbiddenString) {
				continue outer
			}
		}

		vowelsCount := 0
		for _, char := range trimmedLine {
			if _, ok := vowels[char]; ok {
				vowelsCount++
			}
		}
		if vowelsCount < 3 {
			continue
		}

		currentChar := trimmedLine[0]
		for i := 1; i < len(trimmedLine); i++ {
			if currentChar == trimmedLine[i] {
				totalNiceStrings++
				continue outer
			} else {
				currentChar = trimmedLine[i]
			}
		}
	}
	return fmt.Sprintf("%v", totalNiceStrings)
}
