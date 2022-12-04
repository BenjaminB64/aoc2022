package campcleanup

import (
	"fmt"
	"strings"
)

type Section struct {
	First  int
	Second int
}

func (s Section) FullyContain(section Section) bool {
	return s.First >= section.First && s.Second <= section.Second
}

func (s Section) Overlap(section Section) bool {
	return s.First <= section.First && s.Second >= section.First ||
		s.First <= section.Second && s.Second >= section.Second
}

func SectionFromString(s string) Section {
	var p Section
	fmt.Sscanf(s, "%d-%d", &p.First, &p.Second)
	return p
}

type Pair struct {
	First  Section
	Second Section
}

func (p Pair) FullyContain() bool {
	return p.First.FullyContain(p.Second) || p.Second.FullyContain(p.First)
}

func (p Pair) Overlap() bool {
	return p.First.Overlap(p.Second) || p.Second.Overlap(p.First)
}

func PairFromString(s string) Pair {
	// split string on ','
	ps := strings.Split(s, ",")
	return Pair{
		First:  SectionFromString(ps[0]),
		Second: SectionFromString(ps[1]),
	}
}
