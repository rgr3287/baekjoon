package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	fmt.Scan(&T) // 테스트 케이스

	for i := 0; i < T; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)

		distance := Y - X // 거리

		max := int(math.Sqrt(float64(distance))) // 소수점 버림

		if max*max == distance {
			fmt.Println(max*2 - 1)
		} else if distance <= max*max+max {
			fmt.Println(max * 2)
		} else {
			fmt.Println(max*2 + 1)
		}
	}
}
