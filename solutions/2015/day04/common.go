package day04

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

func findMinNumber(prefix string) int {
	secretKey := input.GetInputAsString(2015, 4)
	secretKey = strings.TrimSpace(secretKey)
	startNumber := 1
	for {
		stringToHash := secretKey + strconv.Itoa(startNumber)
		hash := md5.Sum([]byte(stringToHash))
		if strings.HasPrefix(fmt.Sprintf("%x", hash), prefix) {
			break
		}
		startNumber++
	}
	return startNumber
}
