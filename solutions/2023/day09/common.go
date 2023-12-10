package day09

import (
	"slices"
	"strconv"
)

func parseNumbers(numbers []string) []int {
	result := make([]int, 0, len(numbers))
	for _, number := range numbers {
		parsedNumber, _ := strconv.Atoi(number)
		result = append(result, parsedNumber)
	}
	return result
}

func extrapolateValue(numbers []int, isPartTwo bool) int {
	iterations := make([][]int, 0)
	iterations = append(iterations, numbers)
	lastSlice := iterations[len(iterations)-1]
	lastSliceLenght := len(lastSlice)
	for lastSlice[lastSliceLenght-1] != 0 { // last number of last slice
		newNumbers := make([]int, 0, lastSliceLenght-1)
		for i := 0; i < lastSliceLenght-1; i++ {
			newNumbers = append(newNumbers, (lastSlice[i+1] - lastSlice[i]))
		}
		iterations = append(iterations, newNumbers)
		lastSlice = newNumbers
		lastSliceLenght = len(newNumbers)
	}
	if isPartTwo {
		slices.Reverse(iterations)
		leftMostValue := 0
		for i := 0; i < len(iterations)-1; i++ {
			leftMostValue = iterations[i+1][0] - leftMostValue
		}
		return leftMostValue
	}
	value := 0
	for _, iteration := range iterations {
		value += iteration[len(iteration)-1]
	}
	return value
}
