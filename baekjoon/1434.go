package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	boxSizes := make([]int, n)
	itemSizes := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Scan(&boxSizes[i])
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&itemSizes[i])
	}

	totalAvailableSpace := sum(boxSizes)
	totalSpaceNeeded := sum(itemSizes)
	unusedSpace := totalAvailableSpace - totalSpaceNeeded

	fmt.Println(unusedSpace)
}

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
