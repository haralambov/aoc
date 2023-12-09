package day06

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

func SecondPart() string {
	inputLines := input.GetInputAsSlice(2023, 6)
	if len(inputLines) < 2 {
		return "Invalid input data"
	}
	timeInput := strings.TrimSpace(inputLines[0])
	distanceInput := strings.TrimSpace(inputLines[1])

	re := regexp.MustCompile("[0-9]+")
	times := re.FindAllString(timeInput, -1)
	distances := re.FindAllString(distanceInput, -1)

	timeString := ""
	maxDistanceString := ""

	for timeIndex, time := range times {
		timeString += time
		maxDistanceString += distances[timeIndex]
	}

	time, _ := strconv.Atoi(timeString)
	maxDistance, _ := strconv.Atoi(maxDistanceString)
	numberOfWaysToWin := 0

	for ms := 0; ms < time; ms++ {
		distance := ms * (time - ms)
		if distance > maxDistance {
			numberOfWaysToWin++
		}
	}

	return fmt.Sprintf("%v", numberOfWaysToWin)
}
