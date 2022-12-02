package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	scores := map[string][]int{
		"A X": {1 + 3, 3 + 0}, "A Y": {2 + 6, 1 + 3}, "A Z": {3 + 0, 2 + 6},
		"B X": {1 + 0, 1 + 0}, "B Y": {2 + 3, 2 + 3}, "B Z": {3 + 6, 3 + 6},
		"C X": {1 + 6, 2 + 0}, "C Y": {2 + 0, 3 + 3}, "C Z": {3 + 3, 1 + 6},
	}

	score, scoreB := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		score += scores[s][0]
		scoreB += scores[s][1]
	}

	fmt.Println(score)
	fmt.Println(scoreB)
}
