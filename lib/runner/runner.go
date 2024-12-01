package runner

import (
	"errors"
	"flag"
	"fmt"
)

const (
	MinYear    = 2015
	MaxYear    = 2024
	MinDay     = 1
	MaxDay     = 25
	FirstPart  = 1
	SecondPart = 2
)

func Run() (string, error) {
	year, day, part := parseArguments()
	err := validateArguments(*year, *day, *part)
	if err != nil {
		return "", err
	}
	return RunPart(*year, *day, *part), nil
}

func parseArguments() (*int, *int, *int) {
	year := flag.Int("year", MaxYear, fmt.Sprintf("Selected year. Must be between %v and %v", MinYear, MaxYear))
	day := flag.Int("day", MinDay, fmt.Sprintf("Selected day. Must be between %v and %v", MinDay, MaxDay))
	part := flag.Int("part", FirstPart, fmt.Sprintf("Selected part. Must be either %v or %v", FirstPart, SecondPart))
	flag.Parse()
	return year, day, part
}

func validateArguments(year, day, part int) error {
	if year < MinYear || year > MaxYear {
		return errors.New(fmt.Sprintf("Year must be between %v and %v", MinYear, MaxYear))
	}
	if day < MinDay || day > MaxDay {
		return errors.New(fmt.Sprintf("Day must be between %v and %v", MinDay, MaxDay))
	}
	if part != FirstPart && part != SecondPart {
		return errors.New(fmt.Sprintf("Part must be either %v or %v", FirstPart, SecondPart))
	}
	return nil
}
