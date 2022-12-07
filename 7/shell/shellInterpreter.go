package shell

import (
	"fmt"
	"strconv"
	"strings"
)

type ShellInterpreter struct {
	root       *Directory
	currentDir *Directory
}

func NewShellInterpreter() *ShellInterpreter {
	root := &Directory{
		Name:     "/",
		Parent:   nil,
		Children: []*Directory{},
		Files:    []*File{},
	}
	return &ShellInterpreter{
		root:       root,
		currentDir: root,
	}
}

func (si *ShellInterpreter) InterpretCd(line string) error {
	sLine := strings.Split(line, " ")
	if len(sLine) < 3 {
		return fmt.Errorf("cd line is too short (%s)", line)
	}
	dir := strings.TrimSpace(sLine[2])
	switch dir {
	case "/":
		si.currentDir = si.root
	case ".":
		// do nothing
	case "..":
		if si.currentDir.Parent != nil {
			si.currentDir = si.currentDir.Parent
		} else {
			fmt.Println("cannot change directory to parent, already at root")
		}
	default:
		for _, child := range si.currentDir.Children {
			if child.Name == dir {
				si.currentDir = child
				break
			}
		}
	}
	return nil
}

func (si *ShellInterpreter) InterpretLs(line string) error {
	sLine := strings.Split(line, " ")
	if sLine[0] == "dir" {
		dir := Directory{
			Name:     strings.TrimSpace(line[4:]),
			Parent:   si.currentDir,
			Children: make([]*Directory, 0),
			Files:    make([]*File, 0),
		}
		si.currentDir.Children = append(si.currentDir.Children, &dir)
	} else {
		size, _ := strconv.Atoi(sLine[0])
		file := File{
			Name: strings.TrimSpace(sLine[1]),
			Size: size,
		}
		si.currentDir.Files = append(si.currentDir.Files, &file)
	}
	return nil
}

func (si *ShellInterpreter) CalcTotalSize(maxDirSize int) int {
	si.root.CalcSize(0)
	return si.root.TotalSize(maxDirSize)
}
