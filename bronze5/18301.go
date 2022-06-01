package main

import "fmt"

func main() {
	var n1, n2, n3 uint
	fmt.Scanf("%d %d %d", &n1, &n2, &n3)

	t := (n1+1)*(n2+1)/(n3+1) - 1
	fmt.Println(t)
}

// 문제에서 식을 줌
