package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)

	// 입력 받아서 점수 배열 만들기
	scores := make([]float64, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&scores[i])
	}

	// 최댓값 찾기
	var max float64
	for i := 0; i < num; i++ {
		if scores[i] > max {
			max = scores[i]
		}
	}

	// 평균 계산하기
	var sum float64
	for i := 0; i < num; i++ {
		sum += (scores[i] / max) * 100
	}
	fmt.Println(sum / float64(num))
}
