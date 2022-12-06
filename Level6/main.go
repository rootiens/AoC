package main

import (
	"fmt"
	"os"
	"strings"
)

func find(s string, l int) int {
	for i := l; i < len(s); i++ {
		m := map[rune]struct{}{}
		for _, r := range s[i-l : i] {
			m[r] = struct{}{}
		}
		if len(m) >= l {
			return i
		}
	}
	return -1
}

func main() {
	input, _ := os.ReadFile("input.txt")
	data := strings.TrimSpace(string(input))
	fmt.Println(find(data, 4))
	fmt.Println(find(data, 14))
}
