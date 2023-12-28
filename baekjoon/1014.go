package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SeatCount struct {
	SeatBit int
	Count   int
}

var DP [][]int
var classRoom [][]bool
var N, M int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for t := 0; t < T; t++ {
		scanner.Scan()
		params := strings.Fields(scanner.Text())
		N, _ = strconv.Atoi(params[0])
		M, _ = strconv.Atoi(params[1])

		DP = make([][]int, N+1)
		for i := range DP {
			DP[i] = make([]int, 1<<M)
		}

		classRoom = make([][]bool, N+1)
		for n := 1; n <= N; n++ {
			scanner.Scan()
			input := scanner.Text()
			classRoom[n] = make([]bool, M)
			for m := 0; m < M; m++ {
				classRoom[n][m] = input[m:m+1] == "x"
			}
		}

		solve()
	}
}

func solve() {
	ans := 0

	bitsSet := make([]SeatCount, 0)
	for bit := 0; bit < (1 << M); bit++ {
		if adjCheck(bit) {
			cnt := 0
			for i := 0; i < M; i++ {
				check := 1 << i
				if (bit & check) >= 1 {
					cnt++
				}
			}
			bitsSet = append(bitsSet, SeatCount{bit, cnt})
		}
	}

	for i := 1; i <= N; i++ {
		for _, bit := range bitsSet {
			if !seatCheck(i, bit.SeatBit) {
				continue
			}
			for _, fbit := range bitsSet {
				if bitsCheck(bit.SeatBit, fbit.SeatBit) {
					DP[i][bit.SeatBit] = max(DP[i][bit.SeatBit], DP[i-1][fbit.SeatBit]+bit.Count)
					ans = max(ans, DP[i][bit.SeatBit])
				}
			}
		}
	}

	fmt.Println(ans)
}

func adjCheck(bit int) bool {
	for i := 0; i < M-1; i++ {
		val := 3 << i
		if (bit & val) == val {
			return false
		}
	}
	return true
}

func bitsCheck(bit, fbit int) bool {
	for i := 0; i < M; i++ {
		if (1<<i)&fbit >= 1 {
			if i > 0 && (1<<(i-1)&bit) >= 1 {
				return false
			}
			if (1<<(i+1))&bit >= 1 {
				return false
			}
		}
	}
	return true
}

func seatCheck(n, bit int) bool {
	for i := 0; i < len(classRoom[n]); i++ {
		if classRoom[n][i] && (bit&(1<<i)) >= 1 {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
