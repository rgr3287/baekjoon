package main

import "fmt"

func main() {
	var arr [9]int
	var max, index int

	for i := 0; i < 9; i++ {
		fmt.Scan(&arr[i])
	}

	for i, value := range arr {
		if value > max {
			max = value
			index = i + 1
		}
	}
	fmt.Println(max)
	fmt.Println(index)
}
