package campcleanup

import (
	"fmt"
	"strings"
)

type Section struct {
	First  int
	Second int
}

func (s Section) IsIn(section Section) bool {
	return s.First >= section.First && s.Second <= section.Second
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

func (p Pair) Overlap() bool {
	return p.First.IsIn(p.Second) || p.Second.IsIn(p.First)
}

func PairFromString(s string) Pair {
	// split string on ','
	ps := strings.Split(s, ",")
	return Pair{
		First:  SectionFromString(ps[0]),
		Second: SectionFromString(ps[1]),
	}
}
