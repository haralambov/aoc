package day04

import (
	"strconv"
	"strings"
)

var (
	playerNumbers  = make(map[string]string)
	winningNumbers = make(map[string]string)
	totalCards     = make(map[int]int)
)

func calcPointsInScratchcard(inputLine string) int {
	defer resetNumberCollections()
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

func calcCopies() int {
	copies := 0
	for playerNumber := range playerNumbers {
		if _, ok := winningNumbers[playerNumber]; ok {
			copies++
		}
	}
	return copies
}

func parseAndBuildInputData(inputLine string) int {
	lineParts := strings.Split(inputLine, ":")
	numbersLinePart := lineParts[1]
	numbersParts := strings.Split(numbersLinePart, "|")
	leftSideNumbers, rightSideNumbers := numbersParts[0], numbersParts[1]
	fillCollection(leftSideNumbers, false)
	fillCollection(rightSideNumbers, true)
	return extractGameId(lineParts[0])
}

func extractGameId(gameIdInputPart string) int {
	gameId := strings.ReplaceAll(gameIdInputPart, "Card", "")
	gameId = strings.Trim(gameId, " ")
	parsedGameId, _ := strconv.ParseInt(gameId, 10, 32)
	return int(parsedGameId)
}

func resetNumberCollections() {
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

func processInputLine(inputLine string) {
	defer resetNumberCollections()
	gameId := parseAndBuildInputData(inputLine)
	copies := calcCopies()
	if _, ok := totalCards[gameId]; ok {
		totalCards[gameId] = totalCards[gameId] + 1
	} else {
		totalCards[gameId] = 1
	}
	calcTotalCopies(gameId, copies)
}

func calcTotalCopies(gameId int, copies int) {
	oldCopies := 1
	if _, ok := totalCards[gameId]; ok {
		oldCopies = totalCards[gameId]
	}
	for i := 0; i < oldCopies; i++ {
		for j := 1; j <= copies; j++ {
			if _, ok := totalCards[gameId+j]; ok {
				totalCards[gameId+j] = totalCards[gameId+j] + 1
			} else {
				totalCards[gameId+j] = 1
			}
		}
	}
}

func getScratchcardsCount() int {
	result := 0
	for _, cardsCount := range totalCards {
		result += cardsCount
	}
	return result
}
