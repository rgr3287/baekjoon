package main

import "fmt"

func gcdFunc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcmFunc(a, b int) int {
	return a * b / gcdFunc(a, b)
}

func main() {
	var T, A, B int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&A, &B)
		result := lcmFunc(A, B)
		fmt.Println(result)
	}
}
