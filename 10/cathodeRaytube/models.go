package cathoderaytube

import (
	"fmt"
	"strconv"
	"strings"
)

type CPU struct {
	Registers     map[string]int
	Cycle         int
	ToCatchCycles []int
	Sum           int
}

func NewCPU() *CPU {
	return &CPU{
		Registers: map[string]int{"X": 1},
	}
}

func (c *CPU) Run(input string) error {
	split := strings.Split(input, " ")
	if split[0] == "addx" {
		n, err := strconv.Atoi(split[1])
		if err != nil {
			return err
		}
		c.AddCycle(2)
		c.Registers["X"] += n
		return nil
	}
	if split[0] == "noop" {
		c.AddCycle(1)
		return nil
	}
	return fmt.Errorf("unknown instruction %s", split[0])
}

func (c *CPU) AddCycle(n int) {
	c.Cycle += n

	if len(c.ToCatchCycles) > 0 && c.Cycle >= c.ToCatchCycles[0] {
		c.Sum += c.ToCatchCycles[0] * c.Registers["X"]
		if len(c.ToCatchCycles) == 1 {
			c.ToCatchCycles = []int{}
			return
		}
		c.ToCatchCycles = append([]int{}, c.ToCatchCycles[1:]...)
	}
}
