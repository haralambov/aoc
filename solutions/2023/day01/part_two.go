package day01

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

func SecondPart() string {
	inputLines := input.GetInputAsSlice(2023, 1)
	re := regexp.MustCompile(`[0-9]|one|two|three|four|five|six|seven|eight|nine`)
	var sum int
	for _, line := range inputLines {
		if trimmedLine := strings.TrimSpace(strings.ToLower(line)); len(trimmedLine) > 0 {
			matches := re.FindAllString(line, -1)
			number := convertWordToDigit(matches[0]) + convertWordToDigit(matches[len(matches)-1])
			convertedNumber, _ := strconv.ParseInt(number, 10, 32)
			sum += int(convertedNumber)
		}
	}
	return fmt.Sprintf("%v", sum)
}

func convertWordToDigit(word string) string {
	digitsAsWords := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	digit, ok := digitsAsWords[word]
	if ok {
		return digit
	}
	return word
}
