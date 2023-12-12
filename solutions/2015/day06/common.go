package day06

import "strings"

func getCommand(inputLine string) string {
	if strings.HasPrefix(inputLine, "turn on") {
		return "turn on"
	} else if strings.HasPrefix(inputLine, "turn off") {
		return "turn off"
	}
	return "toggle"
}
