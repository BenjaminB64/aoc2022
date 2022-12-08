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
