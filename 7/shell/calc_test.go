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
