package campcleanup

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var testInput string

func TestCalc(t *testing.T) {
	result := Calc(testInput)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestCalcSecondPart(t *testing.T) {
	result := CalcSecondPart(testInput)
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}
