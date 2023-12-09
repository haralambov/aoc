package day08

import (
	"fmt"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

var (
	startingLocations = make([]string, 0)
	stepsPerElement   = make(map[string]int, 0)
)

func SecondPart() string {
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
			if strings.HasSuffix(matches[0], "A") {
				startingLocations = append(startingLocations, matches[0])
			}
		}
	}

	for _, startingLocation := range startingLocations {
		found := false
		stepsPerLocation := 0
		currentLocation := startingLocation
		for !found {
			for _, direction := range directions {
				element := mapElements[currentLocation]
				if direction == 'L' {
					currentLocation = element.leftLocation
				} else {
					currentLocation = element.rightLocation
				}
				stepsPerLocation++
				if strings.HasSuffix(currentLocation, "Z") {
					stepsPerElement[startingLocation] = stepsPerLocation
					found = true
					break
				}
			}
		}
	}
	numbersToFindLCM := make([]int, 0)
	for _, steps := range stepsPerElement {
		numbersToFindLCM = append(numbersToFindLCM, steps)
	}
	return fmt.Sprintf("%v", LCM(numbersToFindLCM[0], numbersToFindLCM[1], numbersToFindLCM[2:]...))
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
