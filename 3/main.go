package main

import (
	_ "embed"
	"fmt"

	"github.com/BenjaminB64/aoc2022/3/ruckstack"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(ruckstack.Calc(input))
	fmt.Println(ruckstack.CalcSecondPart(input))
}
