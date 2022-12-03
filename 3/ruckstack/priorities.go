package ruckstack

// toPriority converts an item to a priority
func toPriority(input byte) int {
	// if input is a lowercase letter, return its position in the alphabet
	// if input is a uppercase letter, return its position in the alphabet + 27
	if input >= 'a' && input <= 'z' {
		return int(input - 'a' + 1)
	} else if input >= 'A' && input <= 'Z' {
		return int(input - 'A' + 27)
	}
	return 0
}
