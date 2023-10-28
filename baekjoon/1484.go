package main

import (
	"fmt"
	"sort"
)

const MAX = 100000

func main() {
	var g int
	fmt.Scan(&g)

	powArr := make([]int64, MAX+1)

	// Initialize the array of squared values
	for i := 1; i <= MAX; i++ {
		powArr[i] = int64(i) * int64(i)
	}

	start := 1
	end := 2
	result := []int{}

	for end <= MAX {
		weight := powArr[end] - powArr[start]

		if weight == int64(g) {
			result = append(result, end)
		}

		if weight > int64(g) {
			start++
		}

		if weight <= int64(g) {
			end++
		}
	}

	sort.Ints(result)

	if len(result) == 0 || g == 1 {
		fmt.Println(-1)
	} else {
		for _, ele := range result {
			fmt.Println(ele)
		}
	}
}
