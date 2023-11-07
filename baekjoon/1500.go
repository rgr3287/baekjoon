package main

import (
	"fmt"
)

func main() {
	var S, K int
	fmt.Scan(&S) // 정수 S
	fmt.Scan(&K) // 합이 S인 K개의 정수
	div := S / K
	mod := S % K
	max := int64(1)

	for i := 1; i <= K; i++ {
		if i <= mod { // 나머지 갯수만큼 +1
			max *= int64(div + 1)
		} else {
			max *= int64(div)
		}
	}

	fmt.Println(max)
}
