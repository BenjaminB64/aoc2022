package tuningTrouble

import "fmt"

func Calc(input string) int {
	uniqueSequenceLength := 4
	var chars []byte
	fmt.Println(input)
lookingFor:
	for i := 0; i < len(input)-uniqueSequenceLength; i++ {
		fmt.Println("--------", i, "---------")
		chars = make([]byte, 0, uniqueSequenceLength)
		for j := 0; j < uniqueSequenceLength; j++ {
			fmt.Println("looking for", string(input[i+j]), "in", string(chars))
			for _, c := range chars {
				if c == input[i+j] {
					continue lookingFor
				}
			}
			chars = append(chars, input[i+j])
		}
		return i + uniqueSequenceLength
	}
	return 0
}
