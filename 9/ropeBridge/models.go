package ropeBridge

import "fmt"

// Map
type Map struct {
	VisitedPositions map[int]map[int]struct{}
	XHead            int
	YHead            int
	Followers        []*Position
}

type Position struct {
	X int
	Y int
}

// NewMap creates a new map
func NewMap(nbFollower int) *Map {
	return &Map{
		VisitedPositions: map[int]map[int]struct{}{
			0: {
				0: {},
			},
		},
		XHead:     0,
		YHead:     0,
		Followers: make([]*Position, nbFollower),
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
	m.Follow()
}

func (m *Map) Follow() {
	for i := 0; i < len(m.Followers); i++ {
		x := 0
		y := 0
		if i == 0 {
			x = m.XHead
			y = m.YHead
		} else {
			x = m.Followers[i-1].X
			y = m.Followers[i-1].Y
		}

		if m.Followers[i] == nil {
			m.Followers[i] = &Position{}
		}

		DHeight := y - m.Followers[i].Y
		DWidth := x - m.Followers[i].X
		// absolute value
		if abs(DHeight) <= 1 && abs(DWidth) <= 1 {
			continue
		}

		if DHeight > 0 {
			m.Followers[i].Y += 1
		} else if DHeight < 0 {
			m.Followers[i].Y -= 1
		}
		if DWidth > 0 {
			m.Followers[i].X += 1
		} else if DWidth < 0 {
			m.Followers[i].X -= 1
		}

		if i < len(m.Followers)-1 {
			continue
		}
		if _, ok := m.VisitedPositions[m.Followers[i].X]; !ok {
			m.VisitedPositions[m.Followers[i].X] = map[int]struct{}{}
		}
		m.VisitedPositions[m.Followers[i].X][m.Followers[i].Y] = struct{}{}
	}
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
			for k, follower := range m.Followers {
				if follower != nil && follower.X == j && follower.Y == i {
					fmt.Print(k)
					continue
				}
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
