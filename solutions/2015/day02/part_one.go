package day02

import "fmt"

func FirstPart() string {
	presentRibbon := func(l, w, h int) int {
		return (2*l*w + 2*w*h + 2*h*l)
	}
	additionalRibbon := func(l, w, h int) int {
		return l * w
	}
	totalWrappingPaper := calcWrappingPaper(presentRibbon, additionalRibbon)
	return fmt.Sprintf("%v", totalWrappingPaper)
}
