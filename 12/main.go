package main

import (
	_ "embed"
	"fmt"

	"github.com/BenjaminB64/aoc2022/12/djikstra"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(djikstra.Calc(input))
}
