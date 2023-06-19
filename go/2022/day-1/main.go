package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func max(ar []int) (int, int) {
	max := 0
	idx := 0
	for i, e := range ar {
		if max < e {
			max = e
			idx = i
		}
	}
	return idx, max
}

func sumOf3Max(ar []int) int {
  result := 0
	for i := 0; i < 3; i++ {
		idx, m := max(ar)
    ar = append(ar[:idx], ar[idx+1:]...)
    result += m
	}
  return result
}

func main() {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	str := strings.TrimSuffix(string(stdin), "\n")
	split := strings.Split(str, "\n")

	cals := []int{0}
	idx := 0

	for _, element := range split {
		if len(element) == 0 {
			idx++
			cals = append(cals, 0)
			continue
		}
		cal, err := strconv.Atoi(element)
		if err != nil {
			panic(err)
		}
		cals[idx] += cal
	}

	// part 1
	_, m := max(cals)
	fmt.Printf("Answer for part 1: %d\n", m)

	// part 2
	fmt.Printf("Answer for part 2: %d\n", sumOf3Max(cals))
}
