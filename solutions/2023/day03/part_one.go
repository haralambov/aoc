package day03

import (
	"fmt"
	"regexp"

	"github.com/haralambov/aoc/lib/input"
)

func FirstPart() string {
	inputLines := input.GetInputAsSlice(2023, 3)
	extractNumbersAndSymbols(inputLines, regexp.MustCompile(`[^0-9.]`))

	return fmt.Sprintf("%v", calculateSumOfPartNumbers())
}

func calculateSumOfPartNumbers() int {
	sum := 0
	for _, number := range numbersCollection {
		if isNumberAdjacent(number) {
			sum += number.value
		}
	}
	return sum
}

func isNumberAdjacent(number schematicElement) bool {
	rowIndicesToCheck := [3]int{number.row - 1, number.row, number.row + 1}
	for _, rowIndex := range rowIndicesToCheck {
		if symbolsInRow, ok := symbolsCollection[rowIndex]; ok {
			for symbolIndex := range symbolsInRow {
				if isNumberAdjacentToSymbolIndex(number, symbolIndex) {
					return true
				}
			}
		}
	}
	return false
}
