package cathoderaytube

import "strings"

func Calc(input string) int {
	s := strings.Split(input, "\n")
	cpu := NewCPU()
	cpu.ToCatchCycles = []int{20, 60, 100, 140, 180, 220}
	for _, line := range s {
		cpu.Run(line)
	}
	return cpu.Sum
}
