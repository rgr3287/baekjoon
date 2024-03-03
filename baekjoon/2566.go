package main

import "fmt"

func main() {
	var i, j, max, x, y int
	max = 0
	x, y = 1, 1

	arr := [9][9]int{}

	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			if arr[i][j] > max {
				max = arr[i][j]
				x = i + 1
				y = j + 1
			}
		}
	}

	fmt.Println(max)
	fmt.Println(x, y)
}
