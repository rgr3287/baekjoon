package main

import (
	"fmt"
)

type Pair struct {
	y, x int
}

func main() {
	var N int
	fmt.Scan(&N)

	pairs := make([]Pair, N)

	for i := 0; i < N; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		pair := Pair{y, x}
		pairs[i] = pair
	}

	result := solution(N, pairs)

	fmt.Println(result)
}

// 가능한 삼각형 넓이 최댓값 구하기
// N^3
func solution(N int, pairs []Pair) float64 {
	result := 0.0

	for first := 0; first < N-2; first++ {
		for second := first + 1; second < N-1; second++ {
			for third := second + 1; third < N; third++ {
				tmpWidth := calcTriangle(pairs[first], pairs[second], pairs[third])
				result = max(result, tmpWidth)
			}
		}
	}

	return result
}

// 3 점으로 삼각형 넓이 계산하기
func calcTriangle(a, b, c Pair) float64 {
	var result float64

	first := float64(a.x*b.y + b.x*c.y + c.x*a.y)
	second := float64(b.x*a.y + c.x*b.y + a.x*c.y)

	result = 0.5 * Abs(first-second)

	return result
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Abs returns the absolute value of x.
func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
