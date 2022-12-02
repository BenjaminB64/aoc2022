package rps

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var inputTest string

func TestCalc(t *testing.T) {
	want := 15
	got := CalcRounds(inputTest)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalc2(t *testing.T) {
	want := 12
	got := CalcRounds2(inputTest)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
