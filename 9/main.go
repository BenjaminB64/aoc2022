package main

import (
	_ "embed"
	"fmt"

	"github.com/BenjaminB64/aoc2022/9/ropeBridge"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(ropeBridge.Calc(input, 1))
	fmt.Println(ropeBridge.Calc(input, 9))
}
