package main

import (
	"fmt"
	"math"
)

func main() {
	var target, m int
	fmt.Scan(&target)
	fmt.Scan(&m)

	broken := make([]bool, 10)
	for i := 0; i < m; i++ {
		var n int
		fmt.Scan(&n)
		broken[n] = true
	}

	result := int(math.Abs(float64(target - 100))) // 초기값 설정
	for i := 0; i <= 999999; i++ {
		str := fmt.Sprintf("%d", i)
		len := len(str)

		isBreak := false
		for j := 0; j < len; j++ {
			if broken[int(str[j]-'0')] { // 고장난 버튼을 눌러야 하면
				isBreak = true
				break // 더 이상 탐색하지 않고 빠져나온다.
			}
		}
		if !isBreak { // i를 누를 때 고장난 버튼을 누르지 않는다면
			min := int(math.Abs(float64(target-i)) + float64(len)) // i를 누른 후 (len) target까지 이동하는 횟수 (target - i)
			result = int(math.Min(float64(min), float64(result)))
		}
	}
	fmt.Println(result)
}
