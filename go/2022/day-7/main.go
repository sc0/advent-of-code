package main

import (
	"day-7/dir"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdin, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(strings.TrimSuffix(string(stdin), "\n"), "\n")[1:]

	root := dir.New(nil)
	currentDir := &root

	for _, line := range lines {
		words := strings.Split(line, " ")
		if words[0] == "$" {
			if words[1] == "cd" {
				if words[2] == ".." {
					currentDir = currentDir.Parent
					continue
				}
				currentDir = currentDir.Tree[words[2]]
				continue
			}
		} else {
			if words[0] == "dir" {
				d := dir.New(currentDir)

				currentDir.Tree[words[1]] = &d
			} else {
				fs, _ := strconv.Atoi(words[0])
				currentDir.FilesSize += fs
			}
		}
	}

  fmt.Println("Part 1", root.Part1())
  fmt.Println("Part 2", root.Part2(70_000_000 - root.CalculateSize()))
}
