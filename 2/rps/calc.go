package rps

import (
	"strings"
)

// Calulate the winner of a round
func calculateMyScore(opponentChoice, myChoice Choice) int {
	if opponentChoice == myChoice {
		return 3
	}
	if myChoice.Beats(opponentChoice) {
		return 6
	}
	return 0
}

func CalcRounds(input string) int {
	var total int
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		opponentChoice := Choice(line[0])
		myChoice := FromMyChoice(line[2])
		score := calculateMyScore(opponentChoice, myChoice) + myChoice.Score()
		total += score
	}
	return total
}

// 2nd part
func CalcRounds2(input string) int {
	var total int
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		opponentChoice := Choice(line[0])
		neededEnd := NeededEnd(line[2])
		myChoice := opponentChoice.ChoiceToMake(neededEnd)
		score := calculateMyScore(opponentChoice, myChoice) + myChoice.Score()
		total += score
	}
	return total
}
