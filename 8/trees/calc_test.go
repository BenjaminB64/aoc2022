package trees

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var input string

func TestCalc(t *testing.T) {
	result := Calc(input)
	if result != 21 {
		t.Errorf("expected 21, got %d", result)
	}
}

func TestCalcSecondPart(t *testing.T) {
	result := CalcSecondPart(input)
	if result != 8 {
		t.Errorf("expected 8, got %d", result)
	}
}
