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
	result := Calc("bvwbjplbgvbhsrlpgdmjqwftvncz", 4)
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
	result = Calc("nppdvjthqldpwncqszvftbrmjlhg", 4)
	if result != 6 {
		t.Errorf("expected 6, got %d", result)
	}
	result = Calc("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4)
	if result != 10 {
		t.Errorf("expected 10, got %d", result)
	}
	result = Calc("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4)
	if result != 11 {
		t.Errorf("expected 11, got %d", result)
	}
}

/*
mjqjpqmgbljsphdztnvjfqwrcgsmlb: first marker after character 19
bvwbjplbgvbhsrlpgdmjqwftvncz: first marker after character 23
nppdvjthqldpwncqszvftbrmjlhg: first marker after character 23
nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg: first marker after character 29
zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw: first marker after character 26
*/
func TestCalcSecondPart(t *testing.T) {
	result := Calc("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14)
	if result != 19 {
		t.Errorf("expected 19, got %d", result)
	}
	result = Calc("bvwbjplbgvbhsrlpgdmjqwftvncz", 14)
	if result != 23 {
		t.Errorf("expected 23, got %d", result)
	}
	result = Calc("nppdvjthqldpwncqszvftbrmjlhg", 14)
	if result != 23 {
		t.Errorf("expected 23, got %d", result)
	}
	result = Calc("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14)
	if result != 29 {
		t.Errorf("expected 29, got %d", result)
	}
	result = Calc("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14)
	if result != 26 {
		t.Errorf("expected 26, got %d", result)
	}
}
