package campcleanup

import "strings"

func Calc(input string) int {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		// split line into pairs
		// check if pair fully contain the other pairs
		// if it is, add 1 to the sum
		pair := PairFromString(line)
		if pair.FullyContain() {
			sum++
		}
	}
	return sum
}

func CalcSecondPart(input string) int {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		// split line into pairs
		// check if pair have an overlap
		// if it is, add 1 to the sum
		pair := PairFromString(line)
		if pair.Overlap() {
			sum++
		}
	}
	return sum
}
