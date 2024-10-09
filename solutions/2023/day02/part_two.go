package day02

import (
	"fmt"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

func SecondPart() string {
	inputLines := input.GetInputAsSlice(2023, 2)
	sum := 0
	for _, line := range inputLines {
		lineParts := strings.Split(line, ":")
		_, gamesPart := lineParts[0], lineParts[1]
		gameRounds := GetGameRounds(gamesPart)
		minimumSetOfCubes := getMinimumSetOfCubesForGame(gameRounds)
		sum += calculateProductOfCubes(minimumSetOfCubes)
	}
	return fmt.Sprintf("%v", sum)
}

func getMinimumSetOfCubesForGame(gameRounds []string) map[string]int {
	setOfCubes := make(map[string]int, 3)
	for _, game := range gameRounds {
		cubesForGame := GetCubesForGame(game)
		for color, cubesCountInGame := range cubesForGame {
			if cubesCount, ok := setOfCubes[color]; ok {
				if cubesCountInGame > cubesCount {
					setOfCubes[color] = cubesCountInGame
				}
			} else {
				setOfCubes[color] = cubesCountInGame
			}
		}
	}
	return setOfCubes
}

func calculateProductOfCubes(cubes map[string]int) int {
	product := 1
	for _, count := range cubes {
		product *= count
	}
	return product
}
