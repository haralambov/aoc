package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

var bagContents = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func FirstPart() string {
	inputLines := input.GetInputAsSlice(2023, 2)
	sum := 0
	for _, line := range inputLines {
		lineParts := strings.Split(line, ":")
		gameIdPart, gamesPart := lineParts[0], lineParts[1]
		if isGamePossible(GetGameRounds(gamesPart)) {
			sum += extractGameId(gameIdPart)
		}
	}
	return fmt.Sprintf("%v", sum)
}

func isGamePossible(gameRounds []string) bool {
	for _, game := range gameRounds {
		cubesForGame := GetCubesForGame(game)
		for color, cubes := range cubesForGame {
			if cubes > bagContents[color] {
				return false
			}
		}
	}
	return true
}

func extractGameId(gameIdPart string) int {
	id := strings.Split(gameIdPart, " ")[1]
	convertedId, _ := strconv.ParseInt(id, 10, 32)
	return int(convertedId)
}
