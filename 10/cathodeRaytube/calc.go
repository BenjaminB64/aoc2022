package cathoderaytube

import (
	"fmt"
	"strings"
)

func Calc(input string) int {
	s := strings.Split(input, "\n")
	f := NewFirstPart()
	for _, line := range s {
		f.CPU.Run(line)
	}
	return f.Sum
}

func CalcSecondPart(input string) string {
	s := strings.Split(input, "\n")
	secondPart := NewSecondPart()

	for _, line := range s {
		secondPart.CPU.Run(line)
	}

	p := secondPart.String()
	fmt.Println(p)
	return p
}
