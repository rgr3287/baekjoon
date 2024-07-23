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

	// 첫 번째 줄 입력 받기
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	// 두 번째 줄 입력 받기
	line, _ := reader.ReadString('\n')
	tokens := strings.Fields(line)

	count := 0
	for _, token := range tokens {
		num, _ := strconv.Atoi(token)
		if n == num {
			count++
		}
	}

	fmt.Println(count)
}
