package main

import (
	_ "embed"
	"fmt"

	"github.com/BenjaminB64/aoc2022/7/shell"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(shell.Calc(input, 100000))
}
