package distressSignal

import "fmt"

type Packets struct {
	Pairs []*Pair
}

type Pair struct {
	Left  *Packet
	Right *Packet
}

func (p *Pair) Print() {
	fmt.Print("Left:")
	p.Left.Print()
	fmt.Print("Right:")
	p.Right.Print()
	fmt.Println()
}

type Packet struct {
	Parent *Packet
	Values []*Packet
	Value  int
	IsLeaf bool
}

func (p *Packet) Print() {
	if p.IsLeaf {
		fmt.Print(p.Value)
		return
	}
	fmt.Print("[")
	for i, v := range p.Values {
		if i != 0 {
			fmt.Print(",")
		}
		v.Print()
	}
	fmt.Print("]")
}

type Decision = int

const (
	NotRightOrder Decision = -1
	Continue      Decision = 0
	RightOrder    Decision = 1
)

func (p *Packet) IsRightOrder(right *Packet) Decision {
	p.Print()
	fmt.Print(" ")
	right.Print()
	fmt.Println()
	//leftValues := p.Values
	//rightValues := right.Values
	if p.IsLeaf && !right.IsLeaf {
		p.Values = []*Packet{{Value: p.Value, IsLeaf: true}}
		p.IsLeaf = false
	}
	if !p.IsLeaf && right.IsLeaf {
		right.Values = []*Packet{{Value: right.Value, IsLeaf: true}}
		right.IsLeaf = false
	}

	if p.IsLeaf && right.IsLeaf {
		fmt.Println("Compare int:", p.Value, right.Value)
		if p.Value < right.Value {
			return RightOrder
		}
		if p.Value == right.Value {
			return Continue
		}
		return NotRightOrder
	}

	for i, v := range p.Values {
		if len(right.Values) <= i {
			fmt.Println("Right out :", len(p.Values), len(right.Values))
			return NotRightOrder
		}

		d := v.IsRightOrder(right.Values[i])
		if d == NotRightOrder {
			fmt.Println("Not Right Order")
			return NotRightOrder
		}
		if d == RightOrder {
			fmt.Println("Right Order")
			return RightOrder
		}
	}
	if len(p.Values) == len(right.Values) && len(p.Values) == 0 {
		return Continue
	}

	fmt.Println("Right Order")
	return RightOrder
}

func (p *Pair) IsRightOrder() bool {
	return p.Left.IsRightOrder(p.Right) == RightOrder
}
