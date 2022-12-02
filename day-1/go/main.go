package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// main is part two
func main() {
	content, err := ioutil.ReadFile("../list.prod") // .test answer -> 24,000
	if err != nil {
		fmt.Println("Err")
	}

	split := strings.Split(string(content), "\n")

	list := make([]int, 0)
	maxSum := 0
	sum := 0

	for _, l := range split {
		i, _ := strconv.Atoi(strings.Trim(l, "\n"))
		list = append(list, i)
	}

	for _, item := range list {
		if item == 0 {
			if sum > maxSum {
				maxSum = sum
			}
			sum = 0
			continue
		}
		sum += item
	}
	fmt.Println("Part one:", maxSum)
	partTwo(list)
}

func partTwo(items []int) {
	elfLoad := make([]int, 0) // part two
	total := 0
	sum := 0
	for _, t := range items {
		if t == 0 {
			elfLoad = append(elfLoad, int(sum)) // part two
			sum = 0
			continue
		}

		sum += t
		if t == items[len(items)-1] {
			elfLoad = append(elfLoad, int(sum)) // part two
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(elfLoad)))
	for i := 0; i < 3; i++ {
		total += elfLoad[i]
	}
	fmt.Println("part two: ", total) // test answer -> 45,000
}
