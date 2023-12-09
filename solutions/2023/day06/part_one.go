package day06

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

func FirstPart() string {
	raceData := make(map[int]int, 0)
	inputLines := input.GetInputAsSlice(2023, 6)
	if len(inputLines) < 2 {
		return "Invalid input data"
	}
	timeInput := strings.TrimSpace(inputLines[0])
	distanceInput := strings.TrimSpace(inputLines[1])

	re := regexp.MustCompile("[0-9]+")
	times := re.FindAllString(timeInput, -1)
	distances := re.FindAllString(distanceInput, -1)

	for timeIndex, time := range times {
		parsedTime, _ := strconv.Atoi(time)
		parsedDistance, _ := strconv.Atoi(distances[timeIndex])
		raceData[parsedTime] = parsedDistance
	}

	waysToWin := make(map[int]int, len(raceData))

	for time, maxDistance := range raceData {
		for ms := 0; ms < time; ms++ {
			distance := ms * (time - ms)
			if distance > maxDistance {
				if _, ok := waysToWin[time]; ok {
					waysToWin[time] = waysToWin[time] + 1
				} else {
					waysToWin[time] = 1
				}
			}
		}
	}

	numberOfWaysToWin := 1
	for _, val := range waysToWin {
		numberOfWaysToWin *= val
	}

	return fmt.Sprintf("%v", numberOfWaysToWin)
}
