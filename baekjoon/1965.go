package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	boxes := make([]int, n)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&boxes[i])
		dp[i] = 1
	}

	maxCount := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if boxes[i] > boxes[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				if dp[i] > maxCount {
					maxCount = dp[i]
				}
			}
		}
	}

	fmt.Println(maxCount)
}
