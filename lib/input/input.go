package input

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func GetInputAsString(year, day int) string {
	if inputIsComingFromStdin() {
		return getInputFromStdin()
	}
	filePath, _ := filepath.Abs(fmt.Sprintf("solutions/%v/day%02v/input", year, day))
	b, _ := os.ReadFile(filePath)
	return string(b)
}

func GetInputAsSlice(year, day int) []string {
	return strings.Split(GetInputAsString(year, day), "\n")
}

func inputIsComingFromStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	}
	return false
}

func getInputFromStdin() string {
	data, _ := io.ReadAll(os.Stdin)
	return string(data)
}
