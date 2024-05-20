package main

import "fmt"

func main() {
	var t, ans int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		ans = 0
		for j := 0; j < 5; j++ {
			var x int
			fmt.Scan(&x)
			if x > ans {
				ans = x
			}
		}
		fmt.Printf("Case #%d: %d\n", i, ans)
	}
}
