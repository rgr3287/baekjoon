package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// 포도주 양을 저장할 배열과 DP 테이블을 초기화
	grape := make([]int, n+1)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&grape[i])
	}

	// n이 1인 경우에 대한 예외 처리
	if n == 1 {
		fmt.Println(grape[1])
		return
	}

	dp[1] = grape[1]
	dp[2] = grape[1] + grape[2]

	// DP를 사용하여 최대 포도주 양을 계산
	for i := 3; i <= n; i++ {
		// i번째 잔을 마시지 않는 경우, i-1번째 잔을 마실 때와 i-2번째 잔을 마실 때 중 큰 값을 선택
		dp[i] = max(dp[i-1], max(dp[i-2]+grape[i], dp[i-3]+grape[i-1]+grape[i]))
	}

	fmt.Println(dp[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
