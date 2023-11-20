package main

import (
	"fmt"
)

var white [1001]int
var black [1001]int
var mapD [1001][16][16]int

func main() {
	var index int

	for {
		var w, b int
		_, err := fmt.Scan(&w, &b)
		if err != nil {
			break
		}

		white[index] = w
		black[index] = b
		index++
	}

	fmt.Println(solution(0, 0, 0, index))
}

func solution(i, wIndex, bIndex, N int) int {
	if wIndex == 15 && bIndex == 15 {
		return 0
	}
	if i == N {
		return 0
	}

	if mapD[i][wIndex][bIndex] != 0 {
		return mapD[i][wIndex][bIndex]
	}

	// 선택 안했을 경우
	ans := solution(i+1, wIndex, bIndex, N)
	// white
	if wIndex < 15 {
		ans = maxF(ans, solution(i+1, wIndex+1, bIndex, N)+white[i])
	}
	if bIndex < 15 {
		ans = maxF(ans, solution(i+1, wIndex, bIndex+1, N)+black[i])
	}
	// black
	mapD[i][wIndex][bIndex] = ans
	return mapD[i][wIndex][bIndex]
}

func maxF(a, b int) int {
	if a > b {
		return a
	}
	return b
}
