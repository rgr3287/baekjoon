package main

import "fmt"

func main() {
	var a, b, c, d, e, f int
	fmt.Scan(&a, &b, &c, &d, &e, &f)

	total := a + b + c + d + e + f

	fmt.Println(total * 5)
}
