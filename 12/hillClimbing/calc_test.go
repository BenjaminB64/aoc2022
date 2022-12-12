package hillclimbing

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var input string

func TestCalc(t *testing.T) {
	result := Calc(input)
	if result != 31 {
		t.Errorf("expected 31, got %d", result)
	}
}
