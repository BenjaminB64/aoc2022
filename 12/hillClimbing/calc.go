package hillclimbing

import (
	"fmt"
	"strings"
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
			p := &Point{
				X:                 j,
				Y:                 i,
				DistanceFromStart: 999,
			}
			if s[i][j] == 'S' {
				hm.Start = p
				p.DistanceFromStart = 0
				e = 'a'
			}
			if s[i][j] == 'E' {
				hm.End = p
				e = 'z'
			}
			p.Height = int(e)
			p.Neighbours = []*Point{}

			hm.Values[i] = append(hm.Values[i], p)
		}
	}

	for i := 0; i < hm.Height; i++ {
		for j := 0; j < hm.Width; j++ {
			p := hm.Values[i][j]
			if i > 0 {
				p.Neighbours = append(p.Neighbours, hm.Values[i-1][j])
			}
			if i < hm.Height-1 {
				p.Neighbours = append(p.Neighbours, hm.Values[i+1][j])
			}
			if j > 0 {
				p.Neighbours = append(p.Neighbours, hm.Values[i][j-1])
			}
			if j < hm.Width-1 {
				p.Neighbours = append(p.Neighbours, hm.Values[i][j+1])
			}
		}
	}

	return hm, nil
}

func Calc(input string) int {
	hm, err := ParseInput(input)
	if err != nil {
		panic(err)
	}

	hm.Best = 999

	djikstra(hm, hm.Start, hm.End)
	fmt.Println(call, hm.Best)
	fmt.Println(hm)
	return int(hm.Best)
}

var call = int64(0)

func djikstra(hm *HeightMap, p *Point, pDest *Point) {
	heap := Heap{}
	heap.Push(p)
	for heap.Size() > 0 {
		call++
		p = heap.Pop()
		if p.Visited {
			continue
		}
		p.Visited = true
		if p == pDest {
			if p.DistanceFromStart < hm.Best {
				hm.Best = p.DistanceFromStart
			}
		}
		for _, p2 := range p.Neighbours {
			if p2.Visited {
				continue
			}
			if p2.DistanceFromStart > p.DistanceFromStart+1 && p2.Height-p.Height <= 1 {
				p2.DistanceFromStart = p.DistanceFromStart + 1
				heap.Push(p2)
			}
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
