package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var n int
		fmt.Scan(&n)
		tot := 0
		for j := 1; j <= n; j++ {
			tot += j * (j + 1) * (j + 2) / 2
		}
		fmt.Println(tot)
	}
}
