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
	tokens := strings.Fields(input)

	A, _ := strconv.Atoi(tokens[0])
	B, _ := strconv.Atoi(tokens[1])
	C, _ := strconv.Atoi(tokens[2])

	// 큰 값을 구해서 max에 저장
	max := max(B-A, C-B)

	// 설명에서처럼 값에 -1을 해서 출력
	fmt.Println(max - 1)
}
