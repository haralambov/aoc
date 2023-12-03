package day03

import (
	"fmt"
	"regexp"

	"github.com/haralambov/aoc/lib/input"
)

func SecondPart() string {
	inputLines := input.GetInputAsSlice(2023, 3)
	extractNumbersAndSymbols(inputLines, regexp.MustCompile(`\*`))
	return fmt.Sprintf("%v", calculateGearRatiosSum())
}

func calculateGearRatiosSum() int {
	sum := 0
	for rowIndex, symbolIndexArr := range symbolsCollection {
		for symbolIndex := range symbolIndexArr {
			sum += getGearRatio(rowIndex, symbolIndex)
		}
	}
	return sum
}

func getGearRatio(rowIndex, symbolIndex int) int {
	adjacentNumbers := make([]schematicElement, 0)
	for _, number := range numbersCollection {
		if number.row == rowIndex-1 ||
			number.row == rowIndex ||
			number.row == rowIndex+1 {
			if isNumberAdjacentToSymbolIndex(number, symbolIndex) {
				adjacentNumbers = append(adjacentNumbers, number)
			}
		}
	}
	if len(adjacentNumbers) != 2 {
		return 0
	}
	ratio := 1
	for _, number := range adjacentNumbers {
		ratio *= number.value
	}
	return ratio
}
