package monkey

import "fmt"

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
	Items           []int
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

func (o Operation) Apply(item int) int {
	operand := o.Operand
	if o.OperandIsItself {
		operand = item
	}
	switch o.Operator {
	case Plus:
		return item + operand
	case Minus:
		return item - operand
	case Mult:
		return item * operand
	case Div:
		return item / operand
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
	return Monkey{
		Items:           items,
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
			newWorry = newWorry / 3

			if newWorry%monkey.TestDivisibleBy == 0 {
				m.Monkeys[monkey.ThrowToIfTrue].Items = append(m.Monkeys[monkey.ThrowToIfTrue].Items, newWorry)
			} else {
				m.Monkeys[monkey.ThrowToIfFalse].Items = append(m.Monkeys[monkey.ThrowToIfFalse].Items, newWorry)
			}
			monkey.NbInspectedItem++
		}
		m.Monkeys[i].Items = []int{}
	}
}

func (m Monkeys) Print() {
	for i, monkey := range m.Monkeys {
		fmt.Printf("Monkey %d (%d): %v\n", i, monkey.NbInspectedItem, monkey.Items)
	}
}
