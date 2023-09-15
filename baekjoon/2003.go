package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		nums[i] = num
	}

	count := 0
	sum := 0
	left, right := 0, 0

	for left < N {
		if sum < M {
			if right == N {
				break
			}
			sum += nums[right]
			right++
		} else {
			if sum == M {
				count++
			}
			sum -= nums[left]
			left++
		}
	}

	fmt.Println(count)
}
