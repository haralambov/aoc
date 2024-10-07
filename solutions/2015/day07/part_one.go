package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

var (
	wires            = make(map[string]*wire)
	wiresToCalculate = make(map[string]bool)
)

type wire struct {
	leftOperand  string
	operator     string
	rightOperand string
	value        uint16
	hasValue     bool
}

func getSignalValue(target string) (uint16, bool) {
	if targetWire, ok := wires[target]; ok {
		if targetWire.hasValue {
			return targetWire.value, true
		}
	} else {
		val, _ := strconv.ParseUint(target, 10, 16)
		return uint16(val), true
	}

	return 0, false
}

func (w *wire) calculateSignal() {
	if len(w.operator) == 0 && len(w.rightOperand) == 0 {
		if value, correct := getSignalValue(w.leftOperand); correct {
			w.value = value
			w.hasValue = true
		}
	} else if w.operator == "NOT" {
		if value, correct := getSignalValue(w.rightOperand); correct {
			w.value = ^value
			w.hasValue = true
		}
	} else {
		leftValue, leftCorrect := getSignalValue(w.leftOperand)
		rightValue, rightCorrect := getSignalValue(w.rightOperand)
		if leftCorrect && rightCorrect {
			switch w.operator {
			case "AND":
				w.value = leftValue & rightValue
			case "OR":
				w.value = leftValue | rightValue
			case "LSHIFT":
				w.value = leftValue << rightValue
			case "RSHIFT":
				w.value = leftValue >> rightValue
			}
			w.hasValue = true
		}
	}
}

func newWire(signal string) *wire {
	signalParts := strings.Split(signal, " ")
	var leftOperand, operator, rightOperand string
	switch len(signalParts) {
	case 1:
		leftOperand = signalParts[0]
	case 2:
		operator, rightOperand = signalParts[0], signalParts[1]
	case 3:
		leftOperand, operator, rightOperand = signalParts[0], signalParts[1], signalParts[2]
	}

	return &wire{
		leftOperand:  leftOperand,
		operator:     operator,
		rightOperand: rightOperand,
	}
}

func calculateAllSignals() {
	for len(wiresToCalculate) > 0 {
		for wireKey, wire := range wires {
			if !wire.hasValue {
				wire.calculateSignal()
				if wire.hasValue {
					delete(wiresToCalculate, wireKey)
				}
			}
		}
	}
}

func loadInputLine(inputLine string) {
	if len(inputLine) == 0 {
		return
	}
	lineParts := strings.Split(inputLine, " -> ")
	signal, wireKey := lineParts[0], lineParts[1]
	wires[wireKey] = newWire(signal)
	wiresToCalculate[wireKey] = true
}

func FirstPart() string {
	inputLines := input.GetInputAsSlice(2015, 7)
	for _, inputLine := range inputLines {
		loadInputLine(inputLine)
	}

	calculateAllSignals()
	return fmt.Sprintf("%d", wires["a"].value)
}
