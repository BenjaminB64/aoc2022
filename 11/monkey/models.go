package monkey

import (
	"fmt"
	"math/big"
)

/*
Example input:
Monkey 0:

	Starting items: 79, 98
	Operation: new = old * 19
	Test: divisible by 23
	  If true: throw to monkey 2
	  If false: throw to monkey 3
*/
type Monkeys struct {
	Monkeys []*Monkey
}

func NewMonkeys() Monkeys {
	return Monkeys{}
}

func (m *Monkeys) AddMonkey(monkey *Monkey) {
	m.Monkeys = append(m.Monkeys, monkey)
}

type Monkey struct {
	Items           []*big.Int
	Operation       *Operation
	TestDivisibleBy int
	ThrowToIfFalse  int
	ThrowToIfTrue   int
	NbInspectedItem int
}

type Operation struct {
	Operator        Operator
	Operand         int
	OperandIsItself bool
}

func (o Operation) Apply(item *big.Int) *big.Int {
	operand := big.NewInt(int64(o.Operand))
	if o.OperandIsItself {
		operand = item
	}
	switch o.Operator {
	case Plus:
		return item.Add(item, operand)
	case Minus:
		return item.Sub(item, operand)
	case Mult:
		return item.Mul(item, operand)
	case Div:
		return item.Div(item, operand)
	default:
		panic("unknown operator")
	}
}

type Operator string

const (
	Plus  Operator = "+"
	Minus Operator = "-"
	Mult  Operator = "*"
	Div   Operator = "/"
)

func NewMonkey(items []int, operation *Operation, testDivisibleBy, throwToIfFalse, throwToIfTrue int) Monkey {
	bigInts := make([]*big.Int, len(items))
	for i, item := range items {
		bigInts[i] = big.NewInt(int64(item))
	}

	return Monkey{
		Items:           bigInts,
		Operation:       operation,
		TestDivisibleBy: testDivisibleBy,
		ThrowToIfFalse:  throwToIfFalse,
		ThrowToIfTrue:   throwToIfTrue,
	}
}

func (m *Monkeys) Round() {
	for i, monkey := range m.Monkeys {
		for _, item := range monkey.Items {
			newWorry := monkey.Operation.Apply(item)
			newWorry = newWorry.Div(newWorry, big.NewInt(3))
			mod := big.NewInt(0)
			if mod.Mod(newWorry, big.NewInt(int64(monkey.TestDivisibleBy))).Cmp(big.NewInt(0)) == 0 {
				m.Monkeys[monkey.ThrowToIfTrue].Items = append(m.Monkeys[monkey.ThrowToIfTrue].Items, newWorry)
			} else {
				m.Monkeys[monkey.ThrowToIfFalse].Items = append(m.Monkeys[monkey.ThrowToIfFalse].Items, newWorry)
			}
			monkey.NbInspectedItem++
		}
		m.Monkeys[i].Items = []*big.Int{}
	}
}

func (m *Monkeys) RoundSecondPart() {
	for i, monkey := range m.Monkeys {
		for _, item := range monkey.Items {
			newWorry := monkey.Operation.Apply(item)
			mod := big.NewInt(0)
			if mod.Mod(newWorry, big.NewInt(int64(monkey.TestDivisibleBy))).Cmp(big.NewInt(0)) == 0 {
				m.Monkeys[monkey.ThrowToIfTrue].Items = append(m.Monkeys[monkey.ThrowToIfTrue].Items, newWorry)
			} else {
				m.Monkeys[monkey.ThrowToIfFalse].Items = append(m.Monkeys[monkey.ThrowToIfFalse].Items, newWorry)
			}
			monkey.NbInspectedItem++
		}
		m.Monkeys[i].Items = []*big.Int{}
	}
}

func (m Monkeys) Print() {
	for i, monkey := range m.Monkeys {
		fmt.Printf("Monkey %d (%d): %v\n", i, monkey.NbInspectedItem, monkey.Items)
	}
}
