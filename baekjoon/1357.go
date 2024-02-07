package main

import "fmt"

func Rev(N int) int {
	rev := 0
	for N > 0 {
		rev = rev*10 + N%10
		N /= 10
	}
	return rev
}

func main() {
	var X, Y int
	fmt.Scan(&X)
	fmt.Scan(&Y)

	x := Rev(X)
	y := Rev(Y)
	res := Rev(x + y)

	fmt.Println(res)
}
