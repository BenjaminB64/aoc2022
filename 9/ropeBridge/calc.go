package ropeBridge

import (
	"fmt"
	"strconv"
	"strings"
)

func Calc(input string, snakeSize int) int {
	m := NewMap(snakeSize)
	n := 0
	lines := strings.Split(input, "\n")
	for _, c := range lines {
		if c == "" {
			continue
		}
		n, _ = strconv.Atoi(c[2:])
		for i := 0; i < n; i++ {
			fmt.Println("move", string(c[0]))
			m.Move(c[0])
			//m.Print()
		}
	}
	l := 0
	for _, c := range m.VisitedPositions {
		l += len(c)
	}
	return l
}
