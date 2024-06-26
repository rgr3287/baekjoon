package main

import (
	"fmt"
)

func main() {
	var arr [5][4]int

	score := 0 // 평가 점수 합
	max := 0   // 가장 큰 점수
	who := 0   // 몇 번째 참가자인지

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Scan(&arr[i][j])
			score += arr[i][j]
		}
		if max < score {
			max = score
			who = i + 1 // 해당 참가자 +1 해주기(인덱스는 0부터 시작이므로)
		}
		score = 0
	}

	fmt.Println(who, max)
}
