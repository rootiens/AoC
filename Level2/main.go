package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	score := 0
	scoreB := 0
	for scanner.Scan() {
		opponent := string(scanner.Text()[0])
		me := string(scanner.Text()[2])
		switch opponent {
		case "A":
			if me == "X" {
				score += 1 + 3
				scoreB += 3 + 0
			}
			if me == "Y" {
				score += 2 + 6
				scoreB += 1 + 3
			}
			if me == "Z" {
				score += 3 + 0
				scoreB += 2 + 6
			}
		case "B":
			if me == "X" {
				score += 1 + 0
				scoreB += 1 + 0
			}
			if me == "Y" {
				score += 2 + 3
				scoreB += 2 + 3
			}
			if me == "Z" {
				score += 3 + 6
				scoreB += 3 + 6
			}
		case "C":
			if me == "X" {
				score += 1 + 6
				scoreB += 2 + 0
			}
			if me == "Y" {
				score += 2 + 0
				scoreB += 3 + 3
			}
			if me == "Z" {
				score += 3 + 3
				scoreB += 1 + 6
			}
		}
	}
	fmt.Println(score)
	fmt.Println(scoreB)
}
