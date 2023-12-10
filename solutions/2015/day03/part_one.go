package day03

import (
	"fmt"

	"github.com/haralambov/aoc/lib/input"
)

func FirstPart() string {
	directions := input.GetInputAsString(2015, 3)
	santaX, santaY := 0, 0

	// adding the default house
	houses[0] = make(map[int]int, 0)
	houses[0][0] = 1

	for _, direction := range directions {
		switch direction {
		case '<':
			santaX--
		case '>':
			santaX++
		case '^':
			santaY++
		case 'v':
			santaY--
		default:
			continue
		}
		addHouse(santaX, santaY, houses)
	}
	return fmt.Sprintf("%v", getTotalHouses())
}
