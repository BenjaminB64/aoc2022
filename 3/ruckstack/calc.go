package ruckstack

import "strings"

func Calc(input string) int {
	total := 0 // total sum of all items
	// Split input into lines
	for _, line := range strings.Split(input, "\n") {
		characters := strings.Split(line, "")
		middle := len(characters) / 2
		for i := 0; i < middle; i++ {
			// check if characters appear in the second half of the string
			if strings.Contains(line[middle:], characters[i]) {
				total += toPriority(characters[i][0])
				break
			}
		}
	}
	return total
}
