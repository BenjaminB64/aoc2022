package main

import (
	_ "embed"
	"fmt"

	cathoderaytube "github.com/BenjaminB64/aoc2022/10/cathodeRaytube"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(cathoderaytube.Calc(input))
	fmt.Println(cathoderaytube.CalcSecondPart(input))
}
