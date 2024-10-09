package main

import (
	"fmt"
	"log"

	"github.com/haralambov/aoc/lib/runner"
)

func main() {
	result, err := runner.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
