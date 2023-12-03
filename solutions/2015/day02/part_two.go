package day02

import "fmt"

func SecondPart() string {
	presentRibbon := func(l, w, h int) int {
		return 2*l + 2*w
	}
	additionalRibbon := func(l, w, h int) int {
		return l * w * h
	}
	totalWrappingPaper := calcWrappingPaper(presentRibbon, additionalRibbon)
	return fmt.Sprintf("%v", totalWrappingPaper)
}
