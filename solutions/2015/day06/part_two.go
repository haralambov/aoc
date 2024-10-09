package day06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

func SecondPart() string {
	totalLitLight := 0
	lightsGrid := make(map[int]map[int]int, 0)
	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			if _, ok := lightsGrid[i]; ok {
				lightsGrid[i][j] = 0
			} else {
				lightsGrid[i] = make(map[int]int, 0)
			}
		}
	}
	executeCommand := func(row, col int, command string) {
		switch command {
		case "turn on":
			lightsGrid[row][col] += 1
		case "turn off":
			lightsGrid[row][col] -= 1
			if lightsGrid[row][col] < 0 {
				lightsGrid[row][col] = 0
			}
		case "toggle":
			lightsGrid[row][col] += 2
		}
	}

	instructions := input.GetInputAsSlice(2015, 6)
	for _, instruction := range instructions {
		command := getCommand(instruction)
		instruction = strings.TrimSpace(strings.Replace(instruction, command, "", 1))
		instructionsParts := strings.Split(instruction, " ") // [461,550 through 564,900]

		startCoordignates := strings.Split(instructionsParts[0], ",")
		endCoordinates := strings.Split(instructionsParts[2], ",")
		startRow, _ := strconv.Atoi(startCoordignates[0])
		startCol, _ := strconv.Atoi(startCoordignates[1])
		endRow, _ := strconv.Atoi(endCoordinates[0])
		endCol, _ := strconv.Atoi(endCoordinates[1])

		for row := startRow; row <= endRow; row++ {
			for col := startCol; col <= endCol; col++ {
				executeCommand(row, col, command)
			}
		}
	}
	for row := 0; row <= 999; row++ {
		for col := 0; col <= 999; col++ {
			totalLitLight += lightsGrid[row][col]
		}
	}
	return fmt.Sprintf("%v", totalLitLight)
}
