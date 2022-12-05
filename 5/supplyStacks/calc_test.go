package supplystacks

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var testInput string

func TestCalc(t *testing.T) {
	result := Calc(testInput)
	if result != "CMZ" {
		t.Errorf("expected 'CMZ', got %s", result)
	}
}

func TestCalcRow(t *testing.T) {
	s := Stacks{Stacks: make(map[int]*Stack)}
	calcRowLine("   [A] [B] [C] [D] [E] [F] [G] [H] [I]     [K]", &s)
	if len(s.Stacks) != 10 {
		t.Errorf("Expected 26 stacks, got %d", len(s.Stacks))
	}
}
