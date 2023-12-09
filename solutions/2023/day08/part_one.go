package day08

import (
	"fmt"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

var startingMapLocation = "AAA"

func FirstPart() string {
	inputLines := input.GetInputAsSlice(2023, 8)
	for inputLineIndex, inputLine := range inputLines {
		if trimmedLine := strings.TrimSpace(inputLine); len(trimmedLine) > 0 {
			if inputLineIndex == 0 { // first line
				directions = trimmedLine
				continue
			}

			matches := re.FindAllString(trimmedLine, -1)
			mapElements[matches[0]] = mapElement{
				leftLocation:  matches[1],
				rightLocation: matches[2],
			}
		}
	}

	totalSteps := 0
	found := false
	for !found {
		for _, direction := range directions {
			element := mapElements[startingMapLocation]
			if direction == 'L' {
				startingMapLocation = element.leftLocation
			} else {
				startingMapLocation = element.rightLocation
			}
			totalSteps++
			if startingMapLocation == "ZZZ" {
				found = true
				break
			}
		}
	}
	return fmt.Sprintf("%v", totalSteps)
}
