package day03

import (
	"fmt"

	"github.com/haralambov/aoc/lib/input"
)

func SecondPart() string {
	directions := input.GetInputAsString(2015, 3)
	santaX, santaY, roboX, roboY := 0, 0, 0, 0

	// adding the default house
	houses[0] = make(map[int]int, 0)
	houses[0][0] = 1
	turn := 0

	for _, direction := range directions {
		switch direction {
		case '<':
			if turn%2 == 0 {
				santaX--
			} else {
				roboX--
			}
		case '>':
			if turn%2 == 0 {
				santaX++
			} else {
				roboX++
			}
		case '^':
			if turn%2 == 0 {
				santaY++
			} else {
				roboY++
			}
		case 'v':
			if turn%2 == 0 {
				santaY--
			} else {
				roboY--
			}
		default:
			continue
		}

		if turn%2 == 0 {
			addHouse(santaX, santaY, houses)
		} else {
			addHouse(roboX, roboY, houses)
		}
		turn++
	}

	return fmt.Sprintf("%v", getTotalHouses())
}
