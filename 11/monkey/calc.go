package monkey

import (
	"sort"
	"strings"
)

func ParseInput(input string) Monkeys {
	s := strings.Split(input, "\n")
	monkeys := NewMonkeys()
	parser := NewParser()

	var monkeyLines strings.Builder
	for i := 0; i < len(s); i++ {
		line := s[i]
		if line == "" {
			// End of monkey
			_, monkey := parser.ParseMonkey(monkeyLines.String())
			monkeys.Monkeys = append(monkeys.Monkeys, &monkey)
			monkeyLines.Reset()
			continue
		}
		monkeyLines.WriteString(line)
		monkeyLines.WriteString("\n")
	}
	return monkeys
}

func Calc(input string) int {
	monkeys := ParseInput(input)

	for i := 0; i < 20; i++ {
		monkeys.Round()
		monkeys.Print()
	}
	top := []int{}
	for _, monkey := range monkeys.Monkeys {
		top = append(top, monkey.NbInspectedItem)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(top)))

	return top[0] * top[1]
}

func CalcSecondPart(input string) int {
	monkeys := ParseInput(input)

	for i := 0; i < 10000; i++ {
		switch i {
		case 20:
			monkeys.Print()
		case 1000:
			monkeys.Print()
		case 2000:
			monkeys.Print()
		}
		monkeys.RoundSecondPart()
	}
	monkeys.Print()
	top := []int{}
	for _, monkey := range monkeys.Monkeys {
		top = append(top, monkey.NbInspectedItem)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(top)))

	return top[0] * top[1]
}
