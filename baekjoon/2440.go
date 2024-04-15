package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		for j := i + 1; j <= num; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
