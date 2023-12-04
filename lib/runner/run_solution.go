package runner

import (
	"fmt"

	d01y15 "github.com/haralambov/aoc/solutions/2015/day01"
	d02y15 "github.com/haralambov/aoc/solutions/2015/day02"
	d01y23 "github.com/haralambov/aoc/solutions/2023/day01"
	d02y23 "github.com/haralambov/aoc/solutions/2023/day02"
	d03y23 "github.com/haralambov/aoc/solutions/2023/day03"
	d04y23 "github.com/haralambov/aoc/solutions/2023/day04"
)

func RunPart(year, day, part int) string {
	switch year {
	case 2023:
		return runDay2023(day, part)
	case 2015:
		return runDay2015(day, part)
	default:
		return fmt.Sprintf("Year %v is not implemented yet", year)
	}
}

func runDay2023(day, part int) string {
	switch day {
	case 1:
		if part == FirstPart {
			return d01y23.FirstPart()
		}
		return d01y23.SecondPart()
	case 2:
		if part == FirstPart {
			return d02y23.FirstPart()
		}
		return d02y23.SecondPart()
	case 3:
		if part == FirstPart {
			return d03y23.FirstPart()
		}
		return d03y23.SecondPart()
	case 4:
		if part == FirstPart {
			return d04y23.FirstPart()
		}
		return d04y23.SecondPart()
	default:
		return fmt.Sprintf("Day %v is not implemented yet", day)
	}
}

func runDay2015(day, part int) string {
	switch day {
	case 1:
		if part == FirstPart {
			return d01y15.FirstPart()
		}
		return d01y15.SecondPart()
	case 2:
		if part == FirstPart {
			return d02y15.FirstPart()
		}
		return d02y15.SecondPart()
	default:
		return fmt.Sprintf("Day %v is not implemented yet", day)
	}
}
