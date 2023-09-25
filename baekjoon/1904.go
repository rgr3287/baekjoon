package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	// dp 배열을 생성하고 초기값을 설정
	dp := make([]int, N+1)
	dp[0] = 1
	dp[1] = 1

	// dp 배열
	for i := 2; i <= N; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 15746
	}
	fmt.Println(dp[N])
}
