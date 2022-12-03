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

func findBadge(group []string) uint8 {
	var priority uint8
	for i := 0; i < len(group[0]); i++ {
		for j := 0; j < len(group[1]); j++ {
			for k := 0; k < len(group[2]); k++ {
				if group[0][i] == group[1][j] && group[0][i] == group[2][k] {
					if group[0][i] > 90 {
						priority = group[0][i] - 96
						return priority
					} else {
						priority = group[0][i] - 64 + 26
						return priority
					}
				}
			}
		}
	}
	return priority
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	var priority int
	var badgePriority int
	badge := make([]string, 0)
	for i, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		part1, part2 := lineDivider(line)
		priority += int(findDuplicates(part1, part2)) // part 1
		badge = append(badge, line)
		index := i + 1
		if index%3 == 0 {
			badgePriority += int(findBadge(badge)) // part2
			badge = nil
		}
	}

	fmt.Println("Part 1: ", priority)
	fmt.Println("Part 2: ", badgePriority)
}
