package supplystacks

import "fmt"

type Stacks struct {
	Stacks map[int]*Stack
}

type Stack struct {
	id     int
	Crates []byte
}

func (s *Stack) AddItem(item byte) {
	s.Crates = append(s.Crates, item)
}

func (s *Stack) PrependItem(item byte) {
	s.Crates = append([]byte{item}, s.Crates...)
}

func (s *Stacks) Move(from int, to int) {
	stackTo := s.Stacks[to]
	if stackTo == nil {
		panic("stack to does not exist")
	}
	stackFrom := s.Stacks[from]
	if stackFrom == nil {
		panic("stack from is nil")
	}
	if len(stackFrom.Crates) == 0 {
		panic("stack from is empty")
	}
	stackTo.PrependItem(stackFrom.Crates[0])
	stackFrom.Crates = stackFrom.Crates[1:]
}

func (s *Stacks) GetTopFromEachStack() string {
	var result string
	l := len(s.Stacks)
	for i := 0; i < l; i++ {
		if s, ok := s.Stacks[i]; ok && len(s.Crates) > 0 {
			result += string(s.Crates[0])
		} else {
			result += " "
		}
	}
	return result
}

// Print stacks
func (s *Stacks) Print() {
	l := len(s.Stacks)
	for i := 0; i < l; i++ {
		fmt.Print(i+1, ": ")
		if s, ok := s.Stacks[i]; ok && len(s.Crates) > 0 {
			s.Print()
		} else {
			fmt.Println()
		}
	}
}

// Print stack
func (s *Stack) Print() {
	for _, crate := range s.Crates {
		fmt.Print(string(crate))
	}
	println()
}
