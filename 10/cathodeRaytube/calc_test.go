package cathoderaytube

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var input string

//go:embed resultSecondPart.txt
var resultSecondPart string

func TestCalc(t *testing.T) {
	result := Calc(input)
	if result != 13140 {
		t.Errorf("expected 13140, got %d", result)
	}
}

func TestCalcSecondPart(t *testing.T) {
	result := CalcSecondPart(input)
	if result != resultSecondPart {
		t.Errorf("expected \n%s\n, got \n%s", resultSecondPart, result)
	}
}
