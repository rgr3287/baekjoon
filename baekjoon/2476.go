package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(line))

	// 최대값 저장할 변수 선언
	max := 0

	// 참가 인원 수 만큼 반복하는 for문
	for i := 0; i < N; i++ {
		// 사람이 바뀔 때마다 금액을 초기화
		money := 0
		line, _ := reader.ReadString('\n')
		inputs := strings.Fields(line)
		A, _ := strconv.Atoi(inputs[0])
		B, _ := strconv.Atoi(inputs[1])
		C, _ := strconv.Atoi(inputs[2])

		// 나오는 눈에 따른 금액 저장
		if A == B && B == C {
			money = 10000 + (A * 1000)
		} else if A != B && B != C && A != C {
			money = maxOf(A, maxOf(B, C)) * 100
		} else {
			if A == B || A == C {
				money = 1000 + (A * 100)
			} else {
				money = 1000 + (B * 100)
			}
		}
		// 최대값 저장
		if money > max {
			max = money
		}
	}
	fmt.Println(max)
}

func maxOf(a, b int) int {
	if a > b {
		return a
	}
	return b
}
