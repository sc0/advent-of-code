package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func part1(lines []string) int {
	result := 0

	for _, line := range lines {
		half1 := line[:len(line)/2]
		half2 := line[len(line)/2:]

		for _, c := range half1 {
			if strings.ContainsRune(half2, c) {
				if int(c) >= int('a') && int(c) <= int('z') {
					result += int(c) - int('a') + 1
				} else {
					result += int(c) - int('A') + 27
				}
				half2 = strings.ReplaceAll(half2, string(c), "")
			}
		}
	}

	return result
}

func part2(lines []string) int {
	result := 0

	for i := 0; i < len(lines)/3; i++ {
		line1 := lines[i*3]
		line2 := lines[i*3+1]
		line3 := lines[i*3+2]

		for _, c := range line1 {
			if strings.ContainsRune(line2, c) && strings.ContainsRune(line3, c) {
				if int(c) >= int('a') && int(c) <= int('z') {
					result += int(c) - int('a') + 1
				} else {
					result += int(c) - int('A') + 27
				}
				break
			}
		}
	}

	return result
}

func main() {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSuffix(string(stdin), "\n"), "\n")

	fmt.Printf("part 1: %d\n", part1(lines))
	fmt.Printf("part 2: %d", part2(lines))
}
