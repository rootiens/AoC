package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.Open("input.txt")
	check(err)
	defer data.Close()

	scanner := bufio.NewScanner(data)

	allElves := make([]int, 0)

	temp := 0
	for scanner.Scan() {
		if scanner.Text() != "" {
			cal, _ := strconv.Atoi(scanner.Text())
			temp += cal
		} else {
			allElves = append(allElves, temp)
			temp = 0
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(allElves)))
	//Part 1
	fmt.Println(allElves[0])
	// Part 2
	fmt.Println(allElves[0] + allElves[1] + allElves[2])
}
