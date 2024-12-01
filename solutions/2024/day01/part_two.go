package day01

import (
	"strconv"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

var (
	locationIds           = make([]int, 0)
	locationIdOccurrences = make(map[int]int)
	similarityScore       = 0
)

func SecondPart() string {
	inputLines := input.GetInputAsSlice(2024, 1)

	for _, inputLine := range inputLines {
		lineParts := strings.Split(inputLine, "   ")
		leftLocationId, _ := strconv.Atoi(lineParts[0])
		rightLocationId, _ := strconv.Atoi(lineParts[1])

		leftLocationIds = append(leftLocationIds, leftLocationId)
		locationIdOccurrences[rightLocationId]++
	}

	for i := range leftLocationIds {
		leftSideMultiplier := leftLocationIds[i]
		similarityMultiplier, ok := locationIdOccurrences[leftSideMultiplier]

		if ok {
			similarityScore += similarityMultiplier * leftSideMultiplier
		}
	}

	return strconv.Itoa(similarityScore)
}
