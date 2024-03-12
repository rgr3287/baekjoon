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

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)

	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	l, _ := strconv.Atoi(parts[2])

	ballCnt := make([]int, n+1)
	res := 0
	now := 1
	for {
		ballCnt[now]++
		if ballCnt[now] == m {
			break
		}

		if ballCnt[now]%2 == 1 {
			now += l
			if now > n {
				now -= n
			}
		} else {
			now -= l
			if now <= 0 {
				now += n
			}
		}
		res++
	}
	fmt.Println(res)
}
