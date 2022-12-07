package shell

import (
	"fmt"
	"strings"
)

func Calc(input string, maxDirSize int) int {
	// for each line
	lines := strings.Split(input, "\n")
	si := NewShellInterpreter()
	for i := 0; i < len(lines); i++ {
		sLine := strings.Split(lines[i], " ")
		if len(sLine) < 2 {
			fmt.Println("line", i, "is too short")
			continue
		}
		// if line start with $, interpret as a command
		if sLine[0] == "$" {
			// if line start with "cd", change directory
			// if line start with "ls", use list result to create directories and files on our tree
			switch sLine[1] {
			case "cd":
				si.InterpretCd(lines[i])
			case "ls":
				for j := i + 1; j < len(lines) && !strings.HasPrefix(lines[j], "$"); j++ {
					si.InterpretLs(lines[j])
				}
			default:
				fmt.Println("unknown command", lines[i])
			}
		} else {
			fmt.Println("line", i, "is not a command", lines[i])
		}
	}
	return si.CalcTotalSize(maxDirSize)
}
