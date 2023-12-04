package day04

import "strings"

var (
	playerNumbers  = make(map[string]string)
	winningNumbers = make(map[string]string)
)

func calcPointsInScratchcard(inputLine string) int {
	defer resetCollections()
	parseAndBuildInputData(inputLine)
	return calcPoints()
}

func calcPoints() int {
	sum := 0
	isFirst := true
	for playerNumber := range playerNumbers {
		if _, ok := winningNumbers[playerNumber]; ok {
			if isFirst {
				sum = 1
				isFirst = false
			} else {
				sum *= 2
			}
		}
	}
	return sum
}

func parseAndBuildInputData(inputLine string) {
	lineParts := strings.Split(inputLine, ":")
	numbersLinePart := lineParts[1]
	numbersParts := strings.Split(numbersLinePart, "|")
	leftSideNumbers, rightSideNumbers := numbersParts[0], numbersParts[1]
	fillCollection(leftSideNumbers, false)
	fillCollection(rightSideNumbers, true)
}

func resetCollections() {
	playerNumbers = make(map[string]string)
	winningNumbers = make(map[string]string)
}

func fillCollection(numbersToAdd string, addToPlayerNumbers bool) {
	numbers := strings.Split(numbersToAdd, " ")
	for _, number := range numbers {
		number := strings.Trim(number, " ")
		if len(number) == 0 {
			continue
		}
		if addToPlayerNumbers {
			playerNumbers[number] = number
		} else {
			winningNumbers[number] = number
		}
	}
}
