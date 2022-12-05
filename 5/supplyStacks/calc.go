package supplystacks

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Calc(input string, secondPart bool) string {
	regexDetectRowStackLine := regexp.MustCompile(`^\s*\[`)
	regexDetectStackIds := regexp.MustCompile(`^\s*(\d+)`)

	stacks := Stacks{Stacks: make(map[int]*Stack)}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		if regexDetectStackIds.MatchString(line) {
			continue
		}

		// detect if line is a stack or a move
		if regexDetectRowStackLine.MatchString(line) {
			// stack
			calcRowLine(line, &stacks)
		} else {
			// move
			stacks.Print()
			if secondPart {
				calcMoveLineSecondPart(line, &stacks)
				continue
			}
			calcMoveLine(line, &stacks)
		}
	}
	return stacks.GetTopFromEachStack()
}

func calcRowLine(line string, stacks *Stacks) {
	// detect crate id
	regexDetectCrate := regexp.MustCompile(`\[([A-Z, ]+)\]|\s?(\s{3})\s?`)

	onlySpace := regexp.MustCompile(`^\s*$`)
	stackItems := regexDetectCrate.FindAllString(line, -1)

	for i, item := range stackItems {
		if onlySpace.MatchString(item) {
			continue
		}

		if s, ok := stacks.Stacks[i]; ok {
			s.AddItem(item[1])
		} else {
			stacks.Stacks[i] = &Stack{id: i, Crates: []byte{item[1]}}
		}
	}
}

func calcMoveLine(line string, stacks *Stacks) {
	regexDetectMove := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	matches := regexDetectMove.FindStringSubmatch(line)

	if len(matches) != 4 {
		fmt.Println("error parsing move line", matches)
		return
	}
	fmt.Println("move line", matches)
	nbToMove, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])

	for i := 0; i < nbToMove; i++ {
		stacks.Move(from-1, to-1)
	}
}

func calcMoveLineSecondPart(line string, stacks *Stacks) {
	regexDetectMove := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	matches := regexDetectMove.FindStringSubmatch(line)

	if len(matches) != 4 {
		fmt.Println("error parsing move line", matches)
		return
	}
	fmt.Println("move line", matches)
	nbToMove, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])

	stacks.BatchMove(from-1, to-1, nbToMove)
}
