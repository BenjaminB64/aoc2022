package ropeBridge

import "fmt"

// Map
type Map struct {
	VisitedPositions map[int]map[int]struct{}
	XHead            int
	YHead            int
	XTail            int
	YTail            int
}

// NewMap creates a new map
func NewMap() *Map {
	return &Map{
		VisitedPositions: map[int]map[int]struct{}{
			0: {
				0: {},
			},
		},
		XHead: 0,
		YHead: 0,
		XTail: 0,
		YTail: 0,
	}
}

// Move moves the head of the rope bridge
func (m *Map) Move(direction byte) {
	switch direction {
	case 'U':
		m.YHead += 1
	case 'D':
		m.YHead -= 1
	case 'L':
		m.XHead -= 1
	case 'R':
		m.XHead += 1
	}

	DHeight := m.YHead - m.YTail
	DWidth := m.XHead - m.XTail
	// absolute value
	if abs(DHeight) <= 1 && abs(DWidth) <= 1 {
		return
	}

	if DHeight > 0 {
		m.YTail += 1
	} else if DHeight < 0 {
		m.YTail -= 1
	}
	if DWidth > 0 {
		m.XTail += 1
	} else if DWidth < 0 {
		m.XTail -= 1
	}

	if _, ok := m.VisitedPositions[m.XTail]; !ok {
		m.VisitedPositions[m.XTail] = map[int]struct{}{}
	}
	m.VisitedPositions[m.XTail][m.YTail] = struct{}{}
}

// Absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// print grid of visited positions (10x10 around head)
func (m *Map) Print() {
	for i := m.YHead + 5; i >= m.YHead-5; i-- {
		for j := m.XHead - 5; j <= m.XHead+5; j++ {
			if j == m.XHead && i == m.YHead {
				fmt.Print("H")
				continue
			}
			if j == m.XTail && i == m.YTail {
				fmt.Print("T")
				continue
			}
			if _, ok := m.VisitedPositions[j]; !ok {
				fmt.Print(".")
				continue
			}
			if _, ok := m.VisitedPositions[j][i]; !ok {
				fmt.Print(".")
				continue
			}
			fmt.Print("#")
		}
		fmt.Println()
	}
}
