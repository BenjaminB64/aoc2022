package shell

import (
	"fmt"
	"strings"
)

// Directory is a directory which contains files and directories
type Directory struct {
	Name     string
	Parent   *Directory
	Children []*Directory
	Files    []*File
	Size     int
}

func (d *Directory) CalcSize(level int) int {
	size := 0
	fmt.Printf("%s - %s (%d children(s), %d files)\n", strings.Repeat(" ", level), d.Name, len(d.Children), len(d.Files))
	if d.Children != nil {
		for _, cd := range d.Children {
			scd := cd.CalcSize(level + 1)
			size += scd
		}
	}
	if d.Files != nil {
		for _, cf := range d.Files {
			fmt.Printf("%s - %s (%d)\n", strings.Repeat(" ", level+1), cf.Name, cf.Size)
			size += cf.Size
		}
	}
	fmt.Printf("%s - %s (%d)\n", strings.Repeat(" ", level), d.Name, size)
	d.Size = size
	return size
}

func (d *Directory) TotalSize(max int) int {
	sum := 0
	if max < 0 || d.Size <= max {
		sum = d.Size
	}
	for _, cd := range d.Children {
		sum += cd.TotalSize(max)
	}
	return sum
}

type File struct {
	Name string
	Size int
}
