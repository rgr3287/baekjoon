package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := a; i <= b; i += 60 {
		fmt.Printf("All positions change in year %d\n", i)
	}
}
