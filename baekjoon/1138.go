package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	heights := make([]int, n)
	positions := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&heights[i])
	}

	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if count == heights[i] && positions[j] == 0 {
				positions[j] = i + 1
				break
			} else if positions[j] == 0 {
				count++
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", positions[i])
	}
}
