package monkey

import (
	"sort"
	"strings"
)

func Calc(input string) int {
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
	for i := 0; i < 20; i++ {
		monkeys.Round()
		monkeys.Print()
	}
	// monkeyBusinessLevel := 0
	top := []int{}
	for _, monkey := range monkeys.Monkeys {
		top = append(top, monkey.NbInspectedItem)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(top)))

	return top[0] * top[1]
}
