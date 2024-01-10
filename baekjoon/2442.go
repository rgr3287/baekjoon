package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for p := 0; p < n; p++ {
		for q := 0; q < n-p-1; q++ {
			fmt.Print(" ")
		}
		for q := 0; q < 2*p+1; q++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
