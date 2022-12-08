package trees

import (
	"fmt"
	"strings"
)

func CreateMap(input string) *Map {
	lines := strings.Split(input, "\n")
	m := NewMap()
	for _, line := range lines {
		if line == "" {
			continue
		}
		m.Height += 1
		if m.Width != 0 && len(line) != m.Width {
			panic("invalid map")
		}
		m.Width = len(line)
		m.Trees = append(m.Trees, []*Tree{})
		for _, c := range line {
			m.Trees[m.Height-1] = append(m.Trees[m.Height-1], &Tree{Height: uint(c - '0')})
		}
	}
	return m
}
func Calc(input string) int {
	m := CreateMap(input)
	visibleTrees := 0
	var v bool
	for i := 0; i < m.Height; i++ {
		fmt.Println()
		for j := 0; j < m.Width; j++ {
			v = m.Check(i, j)
			if v {
				visibleTrees += 1
				fmt.Printf("T")
				continue
			}
			fmt.Printf(".")
		}
	}
	return visibleTrees
}

func CalcSecondPart(input string) int {
	m := CreateMap(input)
	var highestScore int
	var score int
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			score = m.CalcVisibilityScore(i, j)
			if score > highestScore {
				highestScore = score
				continue
			}
		}
	}
	return highestScore
}
