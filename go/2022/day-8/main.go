package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(grid [][]int) int {
	visible := len(grid[0])*2 + (len(grid)-2)*2

	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[0])-1; x++ {
			val := grid[y][x]
			visibleFromSides := []bool{true, true, true, true}

			for dx := 0; dx < x; dx++ {
				if val <= grid[y][dx] {
					visibleFromSides[0] = false
				}
			}
			if visibleFromSides[0] {
				visible++
				continue
			}

			for dx := x + 1; dx < len(grid[0]); dx++ {
				if val <= grid[y][dx] {
					visibleFromSides[1] = false
				}
			}
			if visibleFromSides[1] {
				visible++
				continue
			}

			for dy := 0; dy < y; dy++ {
				if val <= grid[dy][x] {
					visibleFromSides[2] = false
				}
			}
			if visibleFromSides[2] {
				visible++
				continue
			}

			for dy := y + 1; dy < len(grid); dy++ {
				if val <= grid[dy][x] {
					visibleFromSides[3] = false
				}
			}
			if visibleFromSides[3] {
				visible++
				continue
			}
		}
	}

	return visible
}

func part2(grid [][]int) int {
	bestScenic := 0

	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[0])-1; x++ {
			val := grid[y][x]
			viewInDirection := []int{0, 0, 0, 0}
			scenicScore := 0

			for dx := x - 1; dx >= 0; dx-- {
				if val > grid[y][dx] {
					viewInDirection[0]++
					continue
				}
				viewInDirection[0]++
				break
			}

			for dx := x + 1; dx < len(grid[0]); dx++ {
				if val > grid[y][dx] {
					viewInDirection[1]++
					continue
				}
				viewInDirection[1]++
				break
			}

			for dy := y - 1; dy >= 0; dy-- {
				if val > grid[dy][x] {
					viewInDirection[2]++
					continue
				}
				viewInDirection[2]++
				break
			}

			for dy := y + 1; dy < len(grid); dy++ {
				if val > grid[dy][x] {
					viewInDirection[3]++
					continue
				}
				viewInDirection[3]++
				break
			}

			scenicScore = viewInDirection[0] * viewInDirection[1] * viewInDirection[2] * viewInDirection[3]
			bestScenic = int(math.Max(float64(scenicScore), float64(bestScenic)))
		}
	}

	return bestScenic
}

func main() {
	stdin, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(strings.TrimSuffix(string(stdin), "\n"), "\n")

	var grid [][]int

	for y, line := range lines {
		grid = append(grid, []int{})
		for _, value := range line {
			val, _ := strconv.Atoi(string(value))
			grid[y] = append(grid[y], val)
		}
	}

	fmt.Println("Part 1:", part1(grid))
	fmt.Println("Part 2:", part2(grid))
}
