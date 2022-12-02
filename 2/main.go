package main

import (
	_ "embed"
	"fmt"

	"github.com/BenjaminB64/aoc2022/2/rps"
)

//go:embed input.txt
var input string

func main() {
	// Part 1
	fmt.Println(rps.CalcRounds(input))

	// Part 2
	fmt.Println(rps.CalcRounds2(input))
}
