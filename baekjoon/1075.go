package main

import "fmt"

func main() {
	var n, f int
	fmt.Scan(&n, &f)
	n /= 100

	for i := 0; i < 100; i++ {
		if (n*100+i)%f == 0 {
			fmt.Printf("%02d", i)
			break
		}
	}
}
