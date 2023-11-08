package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	dp := make([][][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([][]int, 3)
		for j := 0; j < 3; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	divNum := 1_000_000

	dp[1][0][0] = 1
	dp[1][1][0] = 1
	dp[1][0][1] = 1

	for i := 2; i <= N; i++ {
		dp[i][0][0] = (dp[i-1][0][0] + dp[i-1][1][0] + dp[i-1][2][0]) % divNum
		dp[i][0][1] = (dp[i-1][0][0] + dp[i-1][1][0] + dp[i-1][2][0] +
			dp[i-1][0][1] + dp[i-1][1][1] + dp[i-1][2][1]) % divNum
		dp[i][1][0] = (dp[i-1][0][0]) % divNum
		dp[i][1][1] = (dp[i-1][0][1]) % divNum
		dp[i][2][0] = (dp[i-1][1][0]) % divNum
		dp[i][2][1] = (dp[i-1][1][1]) % divNum
	}

	answer := (dp[N][0][0] + dp[N][0][1] + dp[N][1][0] + dp[N][1][1] + dp[N][2][0] + dp[N][2][1]) % divNum
	fmt.Println(answer)
}
