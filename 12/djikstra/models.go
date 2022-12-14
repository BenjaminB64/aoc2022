package djikstra

type HeightMap struct {
	Width  int
	Height int
	Values [][]*Point
	Start  *Point
	End    *Point
	Best   float64
}

type Point struct {
	X, Y              int
	Visited           bool
	Height            int
	DistanceFromStart float64
	Neighbours        []*Point
}

// pool of heightmaps
type Pool struct {
	HeightMaps []*HeightMap
}

func (p *Pool) Get() *HeightMap {
	if len(p.HeightMaps) == 0 {
		return &HeightMap{}
	}
	hm := p.HeightMaps[0]
	p.HeightMaps = p.HeightMaps[1:]
	return hm
}

func (hm *HeightMap) reset() {
	hm.Best = 0
	for i := 0; i < hm.Height; i++ {
		for j := 0; j < hm.Width; j++ {
			hm.Values[i][j].Visited = false
		}
	}
}

func (hm *HeightMap) copy() *HeightMap {
	hm2 := &HeightMap{
		Width:  hm.Width,
		Height: hm.Height,
		Start:  hm.Start,
		End:    hm.End,
		Best:   hm.Best,
	}
	for i := 0; i < hm.Height; i++ {
		hm2.Values = append(hm2.Values, []*Point{})
		for j := 0; j < hm.Width; j++ {
			hm2.Values[i] = append(hm2.Values[i], &Point{
				X:       hm.Values[i][j].X,
				Y:       hm.Values[i][j].Y,
				Visited: hm.Values[i][j].Visited,
				Height:  hm.Values[i][j].Height,
			})
		}
	}
	return hm2
}

// print map
func (hm *HeightMap) String() string {
	s := ""
	for i := 0; i < hm.Height; i++ {
		for j := 0; j < hm.Width; j++ {
			if hm.Values[i][j] == hm.Start {
				s += "S"
				continue
			}
			if hm.Values[i][j] == hm.End {
				s += "E"
				continue
			}
			if hm.Values[i][j].Visited {
				s += string(byte(hm.Values[i][j].Height) + 'A' - 'a')
			} else {
				s += string(byte(hm.Values[i][j].Height))
			}
		}
		s += "\n"
	}
	return s
}
