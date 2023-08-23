package main

import (
	"fmt"
	"strconv"
)

func isSelfReplicating(num int) bool {
	squared := num * num
	squaredStr := strconv.Itoa(squared)
	numStr := strconv.Itoa(num)

	return squaredStr[len(squaredStr)-len(numStr):] == numStr
}

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scan(&N)

		if isSelfReplicating(N) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
