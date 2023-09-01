package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	costs := make([][3]int, n)
	dp := make([][3]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&costs[i][0], &costs[i][1], &costs[i][2])
	}

	dp[0] = costs[0]

	for i := 1; i < n; i++ {
		dp[i][0] = costs[i][0] + int(math.Min(float64(dp[i-1][1]), float64(dp[i-1][2])))
		dp[i][1] = costs[i][1] + int(math.Min(float64(dp[i-1][0]), float64(dp[i-1][2])))
		dp[i][2] = costs[i][2] + int(math.Min(float64(dp[i-1][0]), float64(dp[i-1][1])))
	}

	result := int(math.Min(math.Min(float64(dp[n-1][0]), float64(dp[n-1][1])), float64(dp[n-1][2])))
	fmt.Println(result)
}
