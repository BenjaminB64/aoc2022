package main

import (
	_ "embed"
	"fmt"

	supplystacks "github.com/BenjaminB64/aoc2022/5/supplyStacks"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(supplystacks.Calc(input, false))
	fmt.Println(supplystacks.Calc(input, true))
}
