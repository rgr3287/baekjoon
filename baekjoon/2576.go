package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	min := 100

	for i := 0; i < 7; i++ {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())

		// 입력값 홀수 여부 확인
		if N%2 == 1 {
			sum += N
			if N < min {
				min = N
			}
		}
	}

	// 합이 없으면 홀수가 없는 것이니 -1 출력
	if sum == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(sum)
		fmt.Println(min)
	}
}
