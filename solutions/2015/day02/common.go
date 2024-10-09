package day02

import (
	"sort"
	"strconv"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

type calcPresentRibbon func(int, int, int) int
type calcAdditionalRibbon func(int, int, int) int

func parseDimension(dimention string) int {
	parsedNumber, _ := strconv.ParseInt(dimention, 10, 32)
	return int(parsedNumber)
}

func calcWrappingPaper(presentRibbotCalc calcPresentRibbon, additionalRibbonClac calcAdditionalRibbon) int {
	inputLines := input.GetInputAsSlice(2015, 2)
	totalWrappingPaper := 0
	for _, inputLine := range inputLines {
		dimentionStrings := strings.Split(inputLine, "x")
		l := parseDimension(dimentionStrings[0])
		w := parseDimension(dimentionStrings[1])
		h := parseDimension(dimentionStrings[2])
		dimentions := []int{l, w, h}
		sort.Ints(dimentions)
		length, width, height := dimentions[0], dimentions[1], dimentions[2]
		presentRibbon := presentRibbotCalc(length, width, height)
		additionalRibbon := additionalRibbonClac(length, width, height)
		totalWrappingPaper += presentRibbon + additionalRibbon
	}
	return totalWrappingPaper
}
