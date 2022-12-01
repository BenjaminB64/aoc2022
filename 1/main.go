package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	// for each line in input, sum the line and add to total, then keep highest
	var highest uint
	var actualSum uint
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			if actualSum > highest {
				highest = actualSum
			}
			actualSum = 0
			continue
		}
		actualLineNumber, _ := strconv.Atoi(line)
		actualSum += uint(actualLineNumber)
	}
	println(highest)
}
