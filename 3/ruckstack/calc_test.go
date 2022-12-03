package ruckstack

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var testInput string

func TestCalc(t *testing.T) {
	expected := 157
	actual := Calc(testInput)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
