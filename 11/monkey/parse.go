package monkey

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

type Parser struct {
	RegexID *regexp.Regexp
}

func NewParser() Parser {
	return Parser{
		RegexID: regexp.MustCompile(`^Monkey (\d+):$`),
	}
}

func (p Parser) ParseMonkey(input string) (int, Monkey) {
	/*
		Parse a monkey from a string.
			Monkey 0:
			Starting items: 79, 98
			Operation: new = old * 19
			Test: divisible by 23
				If true: throw to monkey 2
				If false: throw to monkey 3
	*/
	id := 0
	m := Monkey{}
	var err error
	var items []int
	for i, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		switch i {
		case 0:
			// Monkey 0:
			id, err = p.parseID(line)
			if err != nil {
				panic(err)
			}
		case 1:
			// Starting items: 79, 98
			items, err = p.parseItems(line)
			m.Items = make([]*big.Int, len(items))
			for i, item := range items {
				m.Items[i] = big.NewInt(int64(item))
			}
			if err != nil {
				panic(err)
			}
		case 2:
			// Operation: new = old * 19
			m.Operation, err = p.parseOperation(line)
			if err != nil {
				panic(err)
			}
		case 3:
			// Test: divisible by 23
			m.TestDivisibleBy, err = p.parseTestDivisibleBy(line)
			if err != nil {
				panic(err)
			}
		case 4:
			// If true: throw to monkey 2
			m.ThrowToIfTrue, err = p.parseThrowTo(line)
			if err != nil {
				panic(err)
			}
		case 5:
			// If false: throw to monkey 3
			m.ThrowToIfFalse, err = p.parseThrowTo(line)
			if err != nil {
				panic(err)
			}
		default:
			panic("too many lines")
		}
	}
	return id, m
}

func (p Parser) parseID(line string) (int, error) {
	// Monkey 0:
	matches := p.RegexID.FindStringSubmatch(line)
	if len(matches) != 2 {
		return 0, fmt.Errorf("could not parse ID from line: %s", line)
	}
	return strconv.Atoi(matches[1])
}

func (p Parser) parseItems(line string) ([]int, error) {
	// Starting items: 79, 98
	line = strings.Trim(line, " ")
	items := []int{}
	for _, item := range strings.Split(line[16:], ", ") {
		i, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

func (p Parser) parseOperation(line string) (*Operation, error) {
	// Operation: new = old * 19
	line = strings.Trim(line, " ")
	s := strings.Split(line[17:], " ")

	if len(s) != 3 {
		return nil, fmt.Errorf("could not parse operation from line: %s", line)
	}

	if s[0] == s[2] {
		return &Operation{
			Operator:        Operator(s[1]),
			Operand:         0,
			OperandIsItself: true,
		}, nil
	}

	operand, err := strconv.Atoi(s[2])
	if err != nil {
		return nil, err
	}

	return &Operation{
		Operator:        Operator(s[1]),
		Operand:         operand,
		OperandIsItself: false,
	}, nil
}

func (p Parser) parseTestDivisibleBy(line string) (int, error) {
	// Test: divisible by 23
	line = strings.Trim(line, " ")

	return strconv.Atoi(line[19:])
}

func (p Parser) parseThrowTo(line string) (int, error) {
	// If true: throw to monkey 2
	return strconv.Atoi(line[len(line)-1:])
}
