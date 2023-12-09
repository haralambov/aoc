package day08

import "regexp"

var (
	re          = regexp.MustCompile("[A-Z]{3}")
	directions  = ""
	mapElements = make(map[string]mapElement, 0)
)

type mapElement struct {
	leftLocation  string
	rightLocation string
}
