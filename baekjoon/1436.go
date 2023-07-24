package main

import (
	"fmt"
	"strconv"
)

func hasThreeConsecutiveSixes(num int) bool {
	strNum := strconv.Itoa(num)
	for i := 0; i < len(strNum)-2; i++ {
		if strNum[i] == '6' && strNum[i+1] == '6' && strNum[i+2] == '6' {
			return true
		}
	}
	return false
}

func findNthMovieTitle(N int) int {
	count := 0
	num := 665

	for count < N {
		num++
		if hasThreeConsecutiveSixes(num) {
			count++
		}
	}

	return num
}

func main() {
	var N int
	fmt.Scan(&N)

	result := findNthMovieTitle(N)
	fmt.Println(result)
}
