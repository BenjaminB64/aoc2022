package djikstra

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var input string

func TestCalc(t *testing.T) {
	result, resultSecondPart := Calc(input)
	if result != 31 {
		t.Errorf("expected 31, got %d", result)
	}
	if resultSecondPart != 29 {
		t.Errorf("expected 29, got %d", resultSecondPart)
	}
}
