package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(assignment string) (int, int) {
	part := strings.Split(assignment, "-")
	start, _ := strconv.Atoi(part[0])
	end, _ := strconv.Atoi(part[1])
	return start, end
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	var sum int
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		pair := strings.Split(line, ",")
		start, end := parse(pair[0])
		startB, endB := parse(pair[1])
		if (start <= startB && end >= endB) || (start >= startB && end <= endB) {
			sum += 1
		}

	}
	fmt.Println(sum)
}
