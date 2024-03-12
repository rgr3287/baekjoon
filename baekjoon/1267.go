package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var YoungSik, MinSik int

	for i := 0; i < N; i++ {
		var K int
		fmt.Scan(&K)

		YoungSik += ((K / 30) + 1) * 10
		MinSik += ((K / 60) + 1) * 15
	}

	if YoungSik == MinSik {
		fmt.Println("Y M", YoungSik)
	} else if YoungSik < MinSik {
		fmt.Println("Y", YoungSik)
	} else {
		fmt.Println("M", MinSik)
	}
}
