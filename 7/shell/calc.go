package shell

import (
	"fmt"
	"strconv"
	"strings"
)

func Calc(input string, maxDirSize int) int {
	// for each line
	lines := strings.Split(input, "\n")
	root := Directory{
		Name:     "/",
		Parent:   nil,
		Children: make([]*Directory, 0),
		Files:    make([]*File, 0),
	}
	currentDir := &root
	for i := 0; i < len(lines); i++ {
		sLine := strings.Split(lines[i], " ")
		if len(sLine) < 2 {
			fmt.Println("line", i, "is too short")
			continue
		}
		// if line start with $, interpret as a command
		if sLine[0] == "$" {
			// if line start with "cd", change directory
			if sLine[1] == "cd" {
				if len(sLine) < 3 {
					fmt.Println("cd command on line", i, "is too short")
					continue
				}
				dir := strings.TrimSpace(sLine[2])
				switch dir {
				case "/":
					currentDir = &root
				case ".":
					// do nothing
				case "..":
					if currentDir.Parent != nil {
						currentDir = currentDir.Parent
					} else {
						fmt.Println("cannot change directory to parent, already at root")
					}
				default:
					for _, child := range currentDir.Children {
						if child.Name == dir {
							currentDir = child
							break
						}
					}
				}
			} else if sLine[1] == "ls" {
				for j := i + 1; j < len(lines) && !strings.HasPrefix(lines[j], "$"); j++ {
					subSLine := strings.Split(lines[j], " ")
					if subSLine[0] == "dir" {
						dir := Directory{
							Name:     strings.TrimSpace(lines[j][4:]),
							Parent:   currentDir,
							Children: make([]*Directory, 0),
							Files:    make([]*File, 0),
						}
						currentDir.Children = append(currentDir.Children, &dir)
					} else {
						size, _ := strconv.Atoi(subSLine[0])
						file := File{
							Name: strings.TrimSpace(subSLine[1]),
							Size: size,
						}
						currentDir.Files = append(currentDir.Files, &file)
					}
				}
			} else {
				fmt.Println("unknown command", lines[i])
			}
		}
	}
	return root.CalcSize(maxDirSize, 0)
}
