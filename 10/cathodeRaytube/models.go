package cathoderaytube

import "strings"

type FirstPart struct {
	CPU           *CPU
	ToCatchCycles []int
	Sum           int
}

func NewFirstPart() *FirstPart {
	f := &FirstPart{
		CPU:           NewCPU(),
		ToCatchCycles: []int{20, 60, 100, 140, 180, 220},
	}
	f.CPU.Clock.RegisterCallback(f.CatchCycle)
	return f
}

// callback called by clock
func (f *FirstPart) CatchCycle(c int) {
	if len(f.ToCatchCycles) > 0 && c >= f.ToCatchCycles[0] {
		f.Sum += f.ToCatchCycles[0] * f.CPU.Registers["X"]
		if len(f.ToCatchCycles) == 1 {
			f.ToCatchCycles = []int{}
			return
		}
		f.ToCatchCycles = append([]int{}, f.ToCatchCycles[1:]...)
	}
}

type SecondPart struct {
	CPU    *CPU
	Screen []bool
}

const width = 40
const height = 6

func NewSecondPart() *SecondPart {
	s := &SecondPart{
		CPU:    NewCPU(),
		Screen: make([]bool, width*height),
	}
	s.CPU.Clock.RegisterCallback(s.CatchCycle)
	return s
}

func (s *SecondPart) CatchCycle(c int) {
	index := c - 1
	// lastRowIndex := width
	row := (index / width) % (height + 1)
	col := index - row*width

	if col > width {
		panic("col >= width")
	}
	if s.CPU.Registers["X"]-1 <= col && s.CPU.Registers["X"]+1 >= col {
		s.Screen[row*width+col] = true
	}
}

func (s *SecondPart) String() string {
	var sb strings.Builder
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if s.Screen[i*width+j] {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
