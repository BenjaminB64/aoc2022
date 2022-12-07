package shell

import (
	"strings"
)

func Calc(input string, maxDirSize int) int {
	// for each line
	lines := strings.Split(input, "\n")

	si := NewShellInterpreter()
	err := si.Interpret(lines)
	if err != nil {
		panic(err)
	}

	return si.CalcTotalSize(maxDirSize)
}

func CalcSecondPart(input string, totalSpaceAvailable, needFreeSpace int) int {
	// for each line
	lines := strings.Split(input, "\n")
	si := NewShellInterpreter()

	err := si.Interpret(lines)
	if err != nil {
		panic(err)
	}

	return si.FindDirectoryToDelete(totalSpaceAvailable, needFreeSpace)
}
