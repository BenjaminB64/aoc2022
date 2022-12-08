package main

import (
	_ "embed"
	"fmt"

	"github.com/BenjaminB64/aoc2022/8/trees"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(trees.Calc(input))
	fmt.Println(trees.CalcSecondPart(input))
}
