package distressSignal

import (
	"fmt"
	"sort"
	"strings"
)

func ParseInput(input string) (*Packets, error) {
	p := &Packets{}
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		pair := Pair{}
		pair.Left = parseLine(lines[i])
		pair.Right = parseLine(lines[i+1])
		p.Pairs = append(p.Pairs, &pair)
		i += 2
	}

	return p, nil
}

func Calc(input string) int {
	packets, err := ParseInput(input)
	if err != nil {
		panic(err)
	}
	s := 0
	for i, pair := range packets.Pairs {
		pair.Print()
		if pair.IsRightOrder() {
			s += i + 1
			fmt.Println("Pair", i, "is right order", s)
		} else {
			fmt.Println("Pair", i, "is not right order", s)
		}
	}

	return s
}

func CalcSecondPart(input string) int {
	lines := strings.Split(input, "\n")
	packets := PacketSlice{}
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		packets = append(packets, parseLine(lines[i]))
	}
	two := parseLine("[[2]]")
	six := parseLine("[[6]]")
	packets = append(packets, two, six)

	sort.Sort(packets)

	twoIndex := -1
	sixIndex := -1

	for i, p := range packets {
		if p == two {
			twoIndex = i + 1
		}
		if p == six {
			sixIndex = i + 1
		}
		p.Print()
		fmt.Println()
	}

	return twoIndex * sixIndex
}

func parseLine(line string) *Packet {
	//fmt.Println("ParseLine", line)
	root := &Packet{}
	p := root
	for i := 0; i < len(line); i++ {
		if line[i] == '[' {
			if i == 0 {
				continue
			}
			tmpP := &Packet{Parent: p}
			p.Values = append(p.Values, tmpP)
			p = tmpP
			continue
		}
		if line[i] == ']' {
			p = p.Parent
			if p == nil {
				return root
			}
			continue
		}

		if line[i] == ',' {
			continue
		}
		// parse the number
		n := 0
		j := i
		for ; j < len(line); j++ {
			if line[j] == ',' || line[j] == ']' || line[j] == '[' {
				break
			}
			n *= 10
			n += int(line[j] - '0')
		}
		i = j - 1
		p.Values = append(p.Values, &Packet{Parent: p, Value: n, IsLeaf: true})
		//fmt.Println("append", n)
	}
	return root
}
