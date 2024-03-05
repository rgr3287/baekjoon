package main

import (
	"fmt"
)

func main() {
	var X, Y, N int
	fmt.Scan(&X) // X원
	fmt.Scan(&Y) // Y그램
	fmt.Scan(&N) // 편의점 갯수
	minPrice := float64(X) / float64(Y) * 1000 // 1000그램 가격

	for i := 0; i < N; i++ {
		var X1, Y1 int
		fmt.Scan(&X1) // 가격
		fmt.Scan(&Y1) // 그램
		price := float64(X1) / float64(Y1) * 1000
		if price < minPrice {
			minPrice = price
		}
	}

	fmt.Println(minPrice)
}
