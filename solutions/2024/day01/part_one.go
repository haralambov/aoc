package day01

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

var (
	leftLocationIds  = make([]int, 0)
	rightLocationIds = make([]int, 0)
	totalDistance    = 0
)

func FirstPart() string {
	inputLines := input.GetInputAsSlice(2024, 1)

	for _, inputLine := range inputLines {
		lineParts := strings.Split(inputLine, "   ")
		leftLocationId, _ := strconv.Atoi(lineParts[0])
		rightLocationId, _ := strconv.Atoi(lineParts[1])

		leftLocationIds = append(leftLocationIds, leftLocationId)
		rightLocationIds = append(rightLocationIds, rightLocationId)
	}

	sort.Ints(leftLocationIds)
	sort.Ints(rightLocationIds)

	for i := range leftLocationIds {
		distance := math.Abs(float64(leftLocationIds[i]) - float64(rightLocationIds[i]))
		totalDistance += int(distance)
	}

	return strconv.Itoa(totalDistance)
}
