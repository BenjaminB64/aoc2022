package main

import (
	_ "embed"
	"fmt"

	"github.com/BenjaminB64/aoc2022/7/shell"
)

//go:embed input.txt
var input string

func main() {
	//fmt.Println(shell.Calc(input, 100000))
	fmt.Println(shell.CalcSecondPart(input, 70000000, 30000000))
}
