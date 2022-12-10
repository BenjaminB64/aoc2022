package ropeBridge

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var input string

//go:embed test2.txt
var input2 string

func TestCalc(t *testing.T) {
	result := Calc(input, 1)
	if result != 13 {
		t.Errorf("expected 13, got %d", result)
	}
	result = Calc(input2, 9)
	if result != 36 {
		t.Errorf("expected 36, got %d", result)
	}
}
