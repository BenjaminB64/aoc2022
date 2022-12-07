package shell

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var input string

func TestCalc(t *testing.T) {
	result := Calc(input, 100000)
	if result != 95437 {
		t.Errorf("expected 95437, got %d", result)
	}
}

func TestCalcSecondPart(t *testing.T) {
	result := CalcSecondPart(input, 70000000, 30000000)
	if result != 24933642 {
		t.Errorf("expected 24933642, got %d", result)
	}
}
