package monkey

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var input string

func TestCalc(t *testing.T) {
	result := Calc(input)
	if result != 10605 {
		t.Errorf("expected 10605, got %d", result)
	}
}

func TestCalcSecondPart(t *testing.T) {
	result := CalcSecondPart(input)
	if result != 2713310158 {
		t.Errorf("expected 2713310158, got %d", result)
	}
}
