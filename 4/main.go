package main

import (
	_ "embed"
	"fmt"

	campcleanup "github.com/BenjaminB64/aoc2022/4/campCleanup"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(campcleanup.Calc(input))
	fmt.Println(campcleanup.CalcSecondPart(input))
}
