package main

import "fmt"

func peopleInApartment(k, n int) int {
	if k == 0 {
		return n
	} else {
		sum := 0
		for i := 1; i <= n; i++ {
			sum += peopleInApartment(k-1, i)
		}
		return sum
	}
}

func main() {
	var T, k, n int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&k, &n)
		fmt.Println(peopleInApartment(k, n))
	}
}
