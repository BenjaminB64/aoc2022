package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var actualSum uint
	// Keep 3 top values
	var topValues [3]uint = [3]uint{0, 0, 0}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			for i, value := range topValues {
				if actualSum > value {
					topValues[i] = actualSum
					// sort topValues in ascending order
					sort.Slice(topValues[:], func(i, j int) bool {
						return topValues[i] < topValues[j]
					})
					fmt.Println("New top value:", topValues)
					break
				}
			}
			actualSum = 0
			continue
		}
		actualLineNumber, _ := strconv.Atoi(line)
		actualSum += uint(actualLineNumber)
	}
	println(topValues[0], topValues[1], topValues[2])
	println(topValues[0] + topValues[1] + topValues[2])
}
