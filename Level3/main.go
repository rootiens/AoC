package main

import (
	"fmt"
	"os"
	"strings"
)

func lineDivider(line string) (part1, part2 string) {
	for i := 0; i < len(line); i++ {
		if i < (len(line) / 2) {
			part1 += string(line[i])
		} else {
			part2 += string(line[i])
		}
	}
	return part1, part2
}

func findDuplicates(part1, part2 string) uint8 {
	var priority uint8
	for i := 0; i < len(part1); i++ {
		for j := 0; j < len(part2); j++ {
			if part1[i] == part2[j] {
				fmt.Println(string(part1[i]))
				if part1[i] > 90 {
					priority = part1[i] - 96
					return priority
				} else {
					priority = part1[i] - 64 + 26
					return priority
				}
			}
		}
	}
	return priority
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	var priority int
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		part1, part2 := lineDivider(line)
		fmt.Println(part1, " ", part2)
		priority += int(findDuplicates(part1, part2))
		fmt.Println("====")
	}
	fmt.Println(priority)
}
