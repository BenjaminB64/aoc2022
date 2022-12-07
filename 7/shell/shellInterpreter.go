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

func (si *ShellInterpreter) Interpret(lines []string) error {
	for i := 0; i < len(lines); i++ {
		sLine := strings.Split(lines[i], " ")
		if len(sLine) < 2 {
			return fmt.Errorf("line %d is too short (%s)", i, lines[i])
		}
		// if line start with $, interpret as a command
		if sLine[0] == "$" {
			// if line start with "cd", change directory
			// if line start with "ls", use list result to create directories and files on our tree
			switch sLine[1] {
			case "cd":
				err := si.InterpretCd(lines[i])
				if err != nil {
					return err
				}
			case "ls":
				var j int
				for j = i + 1; j < len(lines) && !strings.HasPrefix(lines[j], "$"); j++ {
					err := si.InterpretLs(lines[j])
					if err != nil {
						return err
					}
				}
				i = j - 1
			default:
				fmt.Println("unknown command", lines[i])
			}
		} else {
			return fmt.Errorf("line %d does not start with $ (%s)", i, lines[i])
		}
	}
	return nil
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

func (si *ShellInterpreter) FindDirectoryToDelete(totalDiskSpace int, needFreeSpace int) int {
	si.root.CalcSize(0)
	freeSpace := totalDiskSpace - si.root.Size
	minToDelete := needFreeSpace - freeSpace
	fmt.Println("need to delete", minToDelete, "bytes")
	return si.root.FindDirectoryToDelete(minToDelete)
}
