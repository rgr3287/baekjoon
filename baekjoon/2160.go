package main

import (
	"fmt"
)

const MAX = 987654321

var mapArray [5][7][50]byte

func main() {
	var n, firstPicture, secondPicture, cnt, answer int
	answer = MAX

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for j := 0; j < 5; j++ {
			var str string
			fmt.Scan(&str)
			for k := 0; k < 7; k++ {
				mapArray[j][k][i] = str[k]
			}
		}
	}

	for pivot := 0; pivot < n; pivot++ {
		for target := pivot + 1; target < n; target++ {
			cnt = 0
			for i := 0; i < 5; i++ {
				for j := 0; j < 7; j++ {
					if mapArray[i][j][pivot] != mapArray[i][j][target] {
						cnt++
					}
					// If cnt is greater than the current answer, no need to continue searching
					if cnt > answer {
						break
					}
				}
				// If cnt is greater than the current answer, no need to continue searching
				if cnt > answer {
					break
				}
			}

			if cnt < answer {
				answer = cnt
				firstPicture = pivot + 1
				secondPicture = target + 1
			}
		}
	}

	fmt.Printf("%d %d\n", firstPicture, secondPicture)
}
