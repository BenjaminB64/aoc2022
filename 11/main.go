package main

import (
	_ "embed"
	"fmt"

	"github.com/BenjaminB64/aoc2022/11/monkey"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(monkey.Calc(input))
	//fmt.Println(monkey.CalcSecondPart(input))
}
