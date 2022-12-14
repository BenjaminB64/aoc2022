package distressSignal

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var input string

func TestCalc(t *testing.T) {
	result := Calc(input)
	if result != 13 {
		t.Errorf("expected 13, got %d", result)
	}
}

func TestCalcSecondPart(t *testing.T) {
	result := CalcSecondPart(input)
	if result != 140 {
		t.Errorf("expected 13, got %d", result)
	}
}
