package tuningTrouble

import (
	_ "embed"
	"testing"
)

/*
bvwbjplbgvbhsrlpgdmjqwftvncz: first marker after character 5
nppdvjthqldpwncqszvftbrmjlhg: first marker after character 6
nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg: first marker after character 10
zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw: first marker after character 11
*/
func TestCalc(t *testing.T) {
	result := Calc("bvwbjplbgvbhsrlpgdmjqwftvncz")
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
	result = Calc("nppdvjthqldpwncqszvftbrmjlhg")
	if result != 6 {
		t.Errorf("expected 6, got %d", result)
	}
	result = Calc("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	if result != 10 {
		t.Errorf("expected 10, got %d", result)
	}
	result = Calc("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	if result != 11 {
		t.Errorf("expected 11, got %d", result)
	}
}
