package ruckstack

import "strings"

func Calc(input string) int {
	total := 0 // total sum of all items
	// split input into lines
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

func CalcSecondPart(input string) int {
	total := 0 // total sum of all items
	splitted := strings.Split(input, "\n")

	groupSize := 3
	nbGroups := len(splitted) / groupSize
	for i := 0; i < nbGroups; i = i + 1 {
		lineNb := i * groupSize
		characters := strings.Split(splitted[lineNb], "")
		// for each character in the line, check if it appears in the other lines
		// if it does, add its priority to the total
	lookForCharacter:
		for c := 0; c < len(characters); c++ {
			for j := 1; j < groupSize; j++ {
				if !strings.Contains(splitted[lineNb+j], characters[c]) {
					break
				}
				if j == (groupSize - 1) {
					total += toPriority(characters[c][0])
					break lookForCharacter
				}
			}
		}
	}

	return total
}
