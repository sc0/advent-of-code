package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func moveTail(hx int, hy int, tx int, ty int) (int, int) {
	if math.Abs(float64(hx-tx)) <= 1 && math.Abs(float64(hy-ty)) <= 1 {
		return tx, ty
	}

	if hx == tx {
		if hy > ty {
			ty++
			return tx, ty
		}
		ty--
		return tx, ty
	}

	if hy == ty {
		if hx > tx {
			tx++
			return tx, ty
		}
		tx--
		return tx, ty
	}

	if hx > tx {
		tx++
	} else {
		tx--
	}

	if hy > ty {
		ty++
	} else {
		ty--
	}

	return tx, ty
}

func addPosition(positions [][2]int, tx int, ty int) [][2]int {
	for _, ar := range positions {
		if ar[0] == tx && ar[1] == ty {
			return positions
		}
	}
	return append(positions, [2]int{tx, ty})
}

func part1(lines []string) int {
	positions := [][2]int{{0, 0}}

	hx, hy := 0, 0
	tx, ty := 0, 0
	for _, line := range lines {
		var dir string
		var count int
		fmt.Sscanf(line, "%s %d", &dir, &count)

		for ; count > 0; count-- {
			switch dir {
			case "R":
				hx++
			case "U":
				hy++
			case "L":
				hx--
			case "D":
				hy--
			}

			tx, ty = moveTail(hx, hy, tx, ty)
			positions = addPosition(positions, tx, ty)
		}
	}
	return len(positions)
}

func part2(lines []string) int {
	positions := [][2]int{{0, 0}}

	rope := [10][2]int{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	for _, line := range lines {
		var dir string
		var count int
		fmt.Sscanf(line, "%s %d", &dir, &count)

		for ; count > 0; count-- {
			switch dir {
			case "R":
				rope[0][0]++
			case "U":
				rope[0][1]++
			case "L":
				rope[0][0]--
			case "D":
				rope[0][1]--
			}

			for i := 1; i < 10; i++ {
				rope[i][0], rope[i][1] = moveTail(rope[i-1][0], rope[i-1][1], rope[i][0], rope[i][1])
			}
			positions = addPosition(positions, rope[9][0], rope[9][1])
		}
	}
	return len(positions)
}

func main() {
	stdin, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(strings.TrimSuffix(string(stdin), "\n"), "\n")

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 1:", part2(lines))
}
