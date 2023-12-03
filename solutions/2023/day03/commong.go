package day03

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	numbersCollection = make([]schematicElement, 0)
	symbolsCollection = make(map[int]map[int]int)
	numbersRe         = regexp.MustCompile(`[0-9]+`)
)

type schematicElement struct {
	row           int
	startPosition int
	length        int
	value         int
}

func extractNumbersAndSymbols(inputLines []string, symbolsRe *regexp.Regexp) {
	for rowIndex, line := range inputLines {
		if trimmedLine := strings.TrimSpace(line); len(trimmedLine) > 0 {
			numberMatches := numbersRe.FindAllStringIndex(trimmedLine, -1)
			addNumbersToCollection(numberMatches, trimmedLine, rowIndex, true)
			symbolMatches := symbolsRe.FindAllStringIndex(trimmedLine, -1)
			addSymbolsToCollection(symbolMatches, rowIndex)
		}
	}
}

func addNumbersToCollection(matches [][]int, line string, rowIndex int, praseNumber bool) {
	for _, match := range matches {
		startIndex, endIndex := match[0], match[1]
		numbersCollection = append(numbersCollection,
			schematicElement{
				row:           rowIndex,
				startPosition: startIndex,
				length:        endIndex - startIndex,
				value:         parseNumber(startIndex, endIndex, line, praseNumber),
			},
		)
	}
}

func addSymbolsToCollection(symbols [][]int, rowIndex int) {
	for _, symbol := range symbols {
		startIndex := symbol[0]
		if _, ok := symbolsCollection[rowIndex]; !ok {
			symbolsCollection[rowIndex] = make(map[int]int)
		}
		symbolsCollection[rowIndex][startIndex] = 1
	}
}

func parseNumber(startIndex, endIndex int, line string, parseNumber bool) int {
	if !parseNumber {
		return 0
	}
	number, _ := strconv.ParseInt(line[startIndex:endIndex], 10, 32)
	return int(number)
}

func isNumberAdjacentToSymbolIndex(number schematicElement, symbolIndex int) bool {
	if symbolIndex >= number.startPosition-1 &&
		symbolIndex <= (number.startPosition+number.length) {
		return true
	}
	return false
}
