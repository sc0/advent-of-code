package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func part1(rounds []string) int {
	rock := 1
	paper := 2
	scissors := 3

	score := 0

	for _, round := range rounds {
		play := strings.Split(round, " ")
		if play[1] == "X" {
			score += rock
			if play[0] == "A" {
				score += 3
			} else if play[0] == "C" {
				score += 6
			}
		}
		if play[1] == "Y" {
			score += paper
			if play[0] == "A" {
				score += 6
			} else if play[0] == "B" {
				score += 3
			}
		}
		if play[1] == "Z" {
			score += scissors
			if play[0] == "B" {
				score += 6
			} else if play[0] == "C" {
				score += 3
			}
		}
	}
	return score
}

func part2(rounds []string) int {
	rock := 1
	paper := 2
	scissors := 3

	score := 0

	for _, round := range rounds {
		play := strings.Split(round, " ")
		if play[1] == "X" {
			if play[0] == "A" {
				score += scissors
			}
			if play[0] == "B" {
				score += rock 
			}
			if play[0] == "C" {
				score += paper
			}
		}
		if play[1] == "Y" {
			score += 3 
			if play[0] == "A" {
				score += rock 
			}
			if play[0] == "B" {
				score += paper 
			}
			if play[0] == "C" {
				score += scissors 
			}
		}
		if play[1] == "Z" {
			score += 6 
			if play[0] == "A" {
				score += paper 
			}
			if play[0] == "B" {
				score += scissors 
			}
			if play[0] == "C" {
				score += rock 
			}
		}
	}
	return score
}

func main() {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	str := strings.TrimSuffix(string(stdin), "\n")
	rounds := strings.Split(str, "\n")

	fmt.Printf("Part 1: %d\n", part1(rounds))
	fmt.Printf("Part 2: %d", part2(rounds))
}
