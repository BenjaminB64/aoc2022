package cathoderaytube

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var input string

func TestCalc(t *testing.T) {
	result := Calc(input)
	if result != 13140 {
		t.Errorf("expected 13140, got %d", result)
	}
}
