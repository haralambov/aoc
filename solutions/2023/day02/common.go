package day02

import (
	"strconv"
	"strings"
)

func GetCubesForGame(game string) map[string]int {
	cubes := strings.Split(game, ",")
	result := make(map[string]int, len(cubes))
	for _, cube := range cubes {
		cube := strings.Trim(cube, " ")
		cubeParts := strings.Split(cube, " ")
		color, number := cubeParts[1], cubeParts[0]
		parsedNumber, _ := strconv.ParseInt(number, 10, 32)
		result[color] = int(parsedNumber)
	}
	return result
}

func GetGameRounds(gamesPart string) []string {
	return strings.Split(gamesPart, ";")
}
