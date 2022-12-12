package hillclimbing

import (
	"fmt"
	"math"
	"strings"
	"sync/atomic"
)

func ParseInput(input string) (*HeightMap, error) {
	s := strings.Split(input, "\n")
	hm := &HeightMap{}
	for i := 0; i < len(s); i++ {
		if i == 0 {
			hm.Width = len(s[i])
			hm.Height = len(s)
		}
		if len(s[i]) != hm.Width {
			return nil, fmt.Errorf("line %d has wrong length", i)
		}

		hm.Values = append(hm.Values, []*Point{})
		for j := 0; j < len(s[i]); j++ {
			e := s[i][j]
			if s[i][j] == 'S' {
				hm.Start = Point{
					X: j,
					Y: i,
				}
				e = 'a'
			}
			if s[i][j] == 'E' {
				hm.End = Point{
					X: j,
					Y: i,
				}
				e = 'z'
			}
			hm.Values[i] = append(hm.Values[i], &Point{
				X:      j,
				Y:      i,
				Height: int(e - 'a'),
			})
		}
	}

	return hm, nil
}

func Calc(input string) int {
	hm, err := ParseInput(input)
	if err != nil {
		panic(err)
	}

	hm.Best = math.Inf(1)
	climb(hm, 0, hm.Start, 0, 0)
	fmt.Println(call, hm.Best)
	return int(hm.Best)
}

var call = int64(0)

func climb(hm *HeightMap, depth int, p Point, length int, h int) int {
	atomic.AddInt64(&call, 1)

	if p.X < 0 || p.X >= hm.Width || p.Y < 0 || p.Y >= hm.Height {
		return 0
	}
	if math.Abs(float64(hm.Values[p.Y][p.X].Height)-float64(h)) > 1 {
		return 0
	}

	if hm.Values[p.Y][p.X].Visited {
		return 0
	}

	if hm.Values[p.Y][p.X].X == hm.End.X && hm.Values[p.Y][p.X].Y == hm.End.Y {
		if float64(length) < hm.Best {
			hm.Best = float64(length)
		}
		return 0
	}

	if float64(length) > hm.Best {
		return 0
	}

	hm.Values[p.Y][p.X].Visited = true
	// up
	climb(hm, depth+1, Point{X: p.X, Y: p.Y - 1}, length+1, hm.Values[p.Y][p.X].Height)
	// down
	climb(hm, depth+1, Point{X: p.X, Y: p.Y + 1}, length+1, hm.Values[p.Y][p.X].Height)
	// left
	climb(hm, depth+1, Point{X: p.X - 1, Y: p.Y}, length+1, hm.Values[p.Y][p.X].Height)
	// right
	climb(hm, depth+1, Point{X: p.X + 1, Y: p.Y}, length+1, hm.Values[p.Y][p.X].Height)

	hm.Values[p.Y][p.X].Visited = false

	return 0
}
