package cathoderaytube

import (
	"fmt"
	"strconv"
	"strings"
)

type Clock struct {
	C        int
	Callback func(int)
}

func NewClock() *Clock {
	return &Clock{}
}

func (c *Clock) AddCycle(n int) {
	for i := 0; i < n; i++ {
		c.Cycle()
	}
}

func (c *Clock) Cycle() {
	c.C++
	if c.Callback != nil {
		c.Callback(c.C)
	}
}

func (c *Clock) RegisterCallback(cb func(int)) {
	c.Callback = cb
}

type CPU struct {
	Registers map[string]int
	Clock     *Clock
}

func NewCPU() *CPU {
	return &CPU{
		Registers: map[string]int{"X": 1},
		Clock:     NewClock(),
	}
}

func (c *CPU) Run(input string) error {
	split := strings.Split(input, " ")
	if split[0] == "addx" {
		n, err := strconv.Atoi(split[1])
		if err != nil {
			return err
		}
		c.Clock.AddCycle(2)
		c.Registers["X"] += n
		return nil
	}
	if split[0] == "noop" {
		c.Clock.AddCycle(1)
		return nil
	}
	return fmt.Errorf("unknown instruction %s", split[0])
}
