package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	N, _ = strconv.Atoi(readLine(reader))
	mapData := make([][]byte, N)

	for i := 0; i < N; i++ {
		line, _ := reader.ReadString('\n')
		mapData[i] = []byte(line[:N])
	}

	answer := int(^uint(0) >> 1)

	for bit := 1; bit < (1 << uint(N)); bit++ {
		sum := 0

		for j := 0; j < N; j++ {
			back := 0

			for i := 0; i < N; i++ {
				curr := mapData[i][j]

				if (bit & (1 << uint(i))) != 0 {
					curr = reverse(curr)
				}

				if curr == 'T' {
					back++
				}
			}

			sum += minData(back, N-back)
		}

		if sum < answer {
			answer = sum
		}
	}

	fmt.Println(answer)
}

func reverse(curr byte) byte {
	if curr == 'T' {
		return 'H'
	} else {
		return 'T'
	}
}

func minData(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func readLine(reader *bufio.Reader) string {
	str, _ := reader.ReadString('\n')
	return str[:len(str)-1]
}
