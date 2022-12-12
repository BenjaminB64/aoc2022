package main

import (
	_ "embed"
	"fmt"

	hillclimbing "github.com/BenjaminB64/aoc2022/12/hillClimbing"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(hillclimbing.Calc(input))
}
