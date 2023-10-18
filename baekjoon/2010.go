package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	multiTabs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&multiTabs[i])
	}

	totalComputers := multiTabs[0] // 첫 번째 멀티탭은 뽑지 않음

	for i := 1; i < n; i++ {
		totalComputers += multiTabs[i] - 1 // 해당 멀티탭을 사용하고 1대 빼줌
	}

	fmt.Println(totalComputers)
}
