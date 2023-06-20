package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	result := 0
	for _, line := range lines {
		ranges := strings.Split(line, ",")
		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")

		r1_start, _ := strconv.Atoi(range1[0])
		r1_end, _ := strconv.Atoi(range1[1])
		r2_start, _ := strconv.Atoi(range2[0])
		r2_end, _ := strconv.Atoi(range2[1])

		if (r1_start <= r2_start && r1_end >= r2_end) ||
			(r2_start <= r1_start && r2_end >= r1_end) {
			result++
		}
	}
	return result
}

func part2(lines []string) int {
	result := 0
	for _, line := range lines {
		ranges := strings.Split(line, ",")
		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")

		r1_start, _ := strconv.Atoi(range1[0])
		r1_end, _ := strconv.Atoi(range1[1])
		r2_start, _ := strconv.Atoi(range2[0])
		r2_end, _ := strconv.Atoi(range2[1])

		if (r2_start >= r1_start && r2_start <= r1_end) ||
			(r1_start >= r2_start && r1_start <= r2_end) {
			result++
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

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
