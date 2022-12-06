package main

import (
	_ "embed"
	"fmt"

	tuningTrouble "github.com/BenjaminB64/aoc2022/6/tuningTrouble"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(tuningTrouble.Calc(input, 4))
	fmt.Println(tuningTrouble.Calc(input, 14))
}
