package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func part1(lines []string) int {
	cycles := map[int]bool{20: true, 60: true, 100: true, 140: true, 180: true, 220: true}
	x := 1
	iter := 0
	instRunning := 0
	result := 0

	for currentCycle := 1; currentCycle <= 220; currentCycle++ {
		ok := cycles[currentCycle]
		if ok {
			result += x * currentCycle
		}
		if instRunning == 0 {
			if lines[iter] == "noop" {
				iter++
				continue
			}
			instRunning++
			continue
		}
		instRunning++
		if instRunning == 2 {
			instRunning = 0
			var v int
			fmt.Sscanf(lines[iter], "addx %d", &v)
			x += v
			iter++
		}
	}

	return result
}

func part2(lines []string) [][]string {
	result := [][]string{}
	pos := 1
	iter := 0
	instRunning := 0
	for i := 0; i < 6; i++ {
		line := []string{""}
		for c := 1; c <= 40; c++ {
			if instRunning == 0 {
				if lines[iter] == "noop" {
					iter++
				} else {
					instRunning++
				}
			} else {
				instRunning++
			}

			if instRunning == 2 {
				instRunning = 0
				var v int
				fmt.Sscanf(lines[iter], "addx %d", &v)
				pos += v
				iter++
			}

			if pos-1 <= c && pos+1 >= c {
				line = append(line, "#")
			} else {
				line = append(line, ".")
			}


		}
		result = append(result, line)
	}
  return result
}

func main() {
	stdin, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(strings.TrimSuffix(string(stdin), "\n"), "\n")

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Println("Part 2:")
	crtLines := part2(lines)

	for _, line := range crtLines {
		fmt.Println(line)
	}
}
