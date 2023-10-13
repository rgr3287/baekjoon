package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
