package main

import (
	_ "embed"
	"fmt"

	"github.com/BenjaminB64/aoc2022/13/distressSignal"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(distressSignal.Calc(input))
	fmt.Println(distressSignal.CalcSecondPart(input))
}
