package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func check(a string) bool {
	result := true
	for i := 0; i < len(a); i++ {
		result = result && !strings.ContainsRune(a[0:i]+a[i+1:], rune(a[i]))
		if !result {
			return result
		}
	}
	return true
}

func main() {
	stdin, _ := io.ReadAll(os.Stdin)
	input := string(stdin)

	step := 4
	for result := step; ; result++ {
		if check(input[result-step : result]) {
			fmt.Println("Part 1", result)
			break
		}
	}

	step = 14
  for result := step; ; result++ {
		if check(input[result-step : result]) {
			fmt.Println("Part 2", result)
			break
		}
	}
}
