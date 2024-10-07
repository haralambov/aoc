package day07

import (
	"fmt"

	"github.com/haralambov/aoc/lib/input"
)

func SecondPart() string {
	inputLines := input.GetInputAsSlice(2015, 7)
	for _, inputLine := range inputLines {
		loadInputLine(inputLine)
	}

	calculateAllSignals()

	wires["b"].value = wires["a"].value

	for wireKey, wire := range wires {
		if wireKey != "b" {
			wire.hasValue = false
			wiresToCalculate[wireKey] = true
		}
	}

	calculateAllSignals()
	return fmt.Sprintf("%d", wires["a"].value)
}
