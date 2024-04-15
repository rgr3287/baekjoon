package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", i+1)
		if (i+1)%6 == 0 || i == n-1 {
			fmt.Print("Go! ")
		}
	}
}
