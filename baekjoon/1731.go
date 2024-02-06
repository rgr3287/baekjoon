package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	n := make([]int, num)

	for i := 0; i < len(n); i++ {
		fmt.Scan(&n[i])
	}

	if n[2]-n[1] == n[1]-n[0] {
		fmt.Println(n[len(n)-1] + (n[1] - n[0]))
	} else {
		fmt.Println(n[len(n)-1] * (n[1] / n[0]))
	}
}
