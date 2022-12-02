package rps

type Choice byte

const (
	Rock     Choice = 'A'
	Paper    Choice = 'B'
	Scissors Choice = 'C'
)

func (c Choice) Beats(other Choice) bool {
	switch c {
	case Rock:
		return other == Scissors
	case Paper:
		return other == Rock
	case Scissors:
		return other == Paper
	}
	return false
}

func (c Choice) Score() int {
	switch c {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	return 0
}

// FromMyChoice return the adversary choice (A, B, C) for a given letter (X, Y, Z)
func FromMyChoice(b byte) Choice {
	return Choice(b - ('X' - 'A'))
}

// 2nd part
// ChoiceToMake return the choice to make to reach the given end
func (c Choice) ChoiceToMake(end NeededEnd) Choice {
	if end == Draw {
		return c
	}
	choices := []Choice{Rock, Paper, Scissors}
	for _, choice := range choices {
		if choice == c {
			continue
		}
		if c.Beats(choice) == (end == Loss) {
			return choice
		}
	}
	return c
}

type NeededEnd byte

const (
	Loss NeededEnd = 'X'
	Draw NeededEnd = 'Y'
	Win  NeededEnd = 'Z'
)

func (n NeededEnd) String() string {
	switch n {
	case Loss:
		return "Loss"
	case Draw:
		return "Draw"
	case Win:
		return "Win"
	}
	return "unknown"
}
