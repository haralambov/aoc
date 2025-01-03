package runner

import (
	"fmt"

	d01y15 "github.com/haralambov/aoc/solutions/2015/day01"
	d02y15 "github.com/haralambov/aoc/solutions/2015/day02"
	d03y15 "github.com/haralambov/aoc/solutions/2015/day03"
	d04y15 "github.com/haralambov/aoc/solutions/2015/day04"
	d05y15 "github.com/haralambov/aoc/solutions/2015/day05"
	d06y15 "github.com/haralambov/aoc/solutions/2015/day06"
	d07y15 "github.com/haralambov/aoc/solutions/2015/day07"
	d08y15 "github.com/haralambov/aoc/solutions/2015/day08"

	d01y23 "github.com/haralambov/aoc/solutions/2023/day01"
	d02y23 "github.com/haralambov/aoc/solutions/2023/day02"
	d03y23 "github.com/haralambov/aoc/solutions/2023/day03"
	d04y23 "github.com/haralambov/aoc/solutions/2023/day04"
	d06y23 "github.com/haralambov/aoc/solutions/2023/day06"
	d07y23 "github.com/haralambov/aoc/solutions/2023/day07"
	d08y23 "github.com/haralambov/aoc/solutions/2023/day08"
	d09y23 "github.com/haralambov/aoc/solutions/2023/day09"

	d01y24 "github.com/haralambov/aoc/solutions/2024/day01"
)

func RunPart(year, day, part int) string {
	switch year {
	case 2024:
		return runDay2024(day, part)
	case 2023:
		return runDay2023(day, part)
	case 2015:
		return runDay2015(day, part)
	default:
		return fmt.Sprintf("Year %v is not implemented yet", year)
	}
}

func runDay2024(day, part int) string {
	switch day {
	case 1:
		if part == FirstPart {
			return d01y24.FirstPart()
		}
		return d01y24.SecondPart()
	default:
		return fmt.Sprintf("Day %v is not implemented yet", day)
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
	case 6:
		if part == FirstPart {
			return d06y23.FirstPart()
		}
		return d06y23.SecondPart()
	case 7:
		if part == FirstPart {
			return d07y23.FirstPart()
		}
		return d07y23.SecondPart()
	case 8:
		if part == FirstPart {
			return d08y23.FirstPart()
		}
		return d08y23.SecondPart()
	case 9:
		if part == FirstPart {
			return d09y23.FirstPart()
		}
		return d09y23.SecondPart()
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
	case 3:
		if part == FirstPart {
			return d03y15.FirstPart()
		}
		return d03y15.SecondPart()
	case 4:
		if part == FirstPart {
			return d04y15.FirstPart()
		}
		return d04y15.SecondPart()
	case 5:
		if part == FirstPart {
			return d05y15.FirstPart()
		}
		return d05y15.SecondPart()
	case 6:
		if part == FirstPart {
			return d06y15.FirstPart()
		}
		return d06y15.SecondPart()
	case 7:
		if part == FirstPart {
			return d07y15.FirstPart()
		}
		return d07y15.SecondPart()
	case 8:
		if part == FirstPart {
			return d08y15.FirstPart()
		}
		return d08y15.SecondPart()
	default:
		return fmt.Sprintf("Day %v is not implemented yet", day)
	}
}
