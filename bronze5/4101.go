package main

import "fmt"

func main() {
	var a, b uint
	for {
		fmt.Scanf("%d %d", &a, &b)
		if a == 0 && b == 0 {
			break
		} else if a > b {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

// 수 비교 문제
