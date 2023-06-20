package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func parseInput(lines []string) ([]string, [][3]int) {
	var stacks []string
	var moves [][3]int

	movesIdx := 0
	for idx, line := range lines {
		if line[1] == '1' {
			movesIdx = idx + 2
			break
		}
		for stackIdx := 0; stackIdx*4+1 < len(line); stackIdx++ {
			if len(stacks) < stackIdx+1 {
				stacks = append(stacks, "")
			}
			if line[stackIdx*4+1] != ' ' {
				stacks[stackIdx] = stacks[stackIdx] + string(line[stackIdx*4+1])
			}
		}
	}

	for ; movesIdx < len(lines); movesIdx++ {
		var num, src, dest int
		fmt.Sscanf(lines[movesIdx], "move %d from %d to %d", &num, &src, &dest)
		moves = append(moves, [3]int{num, src, dest})
	}

	return stacks, moves
}

func part1(stacks []string, moves [][3]int) string {
	for _, move := range moves {
		for crate := 0; crate < move[0]; crate++ {
			src, dest := move[1]-1, move[2]-1
			stacks[dest] = string(stacks[src][0]) + stacks[dest]
			stacks[src] = stacks[src][1:]
		}
	}
	result := ""
	for _, stack := range stacks {
		result = result + string(stack[0])
	}
	return result
}

func part2(stacks []string, moves [][3]int) string {
	for _, move := range moves {
		num, src, dest := move[0], move[1]-1, move[2]-1
		stacks[dest] = stacks[src][0:num] + stacks[dest]
		stacks[src] = stacks[src][num:]
	}

	result := ""
	for _, stack := range stacks {
		result = result + string(stack[0])
	}
	return result
}

func main() {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(stdin), "\n"), "\n")

	stacks, moves := parseInput(lines)
	fmt.Println(part1(stacks, moves))

	stacks, moves = parseInput(lines)
	fmt.Println(part2(stacks, moves))
}
