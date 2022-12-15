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
	l := p
	r := right
	if l.IsLeaf && !r.IsLeaf {
		l = &Packet{Values: []*Packet{{Value: p.Value, IsLeaf: true}}, IsLeaf: false}
	}
	if !l.IsLeaf && r.IsLeaf {
		r = &Packet{Values: []*Packet{{Value: right.Value, IsLeaf: true}}, IsLeaf: false}
	}

	if l.IsLeaf && r.IsLeaf {
		fmt.Println("Compare int:", l.Value, r.Value)
		if l.Value < r.Value {
			return RightOrder
		}
		if l.Value == r.Value {
			return Continue
		}
		return NotRightOrder
	}

	for i, v := range l.Values {
		if len(r.Values) <= i {
			fmt.Println("Right out :", len(l.Values), len(r.Values))
			return NotRightOrder
		}

		d := v.IsRightOrder(r.Values[i])
		if d == NotRightOrder {
			fmt.Println("Not Right Order")
			return NotRightOrder
		}
		if d == RightOrder {
			fmt.Println("Right Order")
			return RightOrder
		}
	}
	if len(l.Values) < len(r.Values) {
		return RightOrder
	}
	if len(l.Values) > len(r.Values) {
		return NotRightOrder
	}

	fmt.Println("Right Order")
	return Continue
}

func (p *Pair) IsRightOrder() bool {
	return p.Left.IsRightOrder(p.Right) == RightOrder
}

type PacketSlice []*Packet

// implement sort.Interface

func (p PacketSlice) Len() int {
	return len(p)
}

func (p PacketSlice) Less(i, j int) bool {
	return p[i].IsRightOrder(p[j]) == RightOrder
}

func (p PacketSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
