package main

import (
	"fmt"
)

func main() {
	var k, m, n int
	fmt.Scanf("%d", &k)
	for i := 0; i < k; i++ {
		fmt.Scanf("%d %d", &m, &n)
		fmt.Println(n * m / 2)
	}
}
