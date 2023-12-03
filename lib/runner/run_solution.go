package runner

import (
	"fmt"

	"github.com/haralambov/aoc/solutions/2023/day01"
	"github.com/haralambov/aoc/solutions/2023/day02"
)

func RunPart(year, day, part int) string {
	switch year {
	case 2023:
		switch day {
		case 1:
			if part == FirstPart {
				return day01.FirstPart()
			}
			return day01.SecondPart()
		case 2:
			if part == FirstPart {
				return day02.FirstPart()
			}
			return day02.SecondPart()
		default:
			return fmt.Sprintf("Day %v is not implemented yet", day)
		}
	default:
		return fmt.Sprintf("Year %v is not implemented yet", year)
	}
}
