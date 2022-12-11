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
		t.Errorf("expected 0, got %d", result)
	}
}
