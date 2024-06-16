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
	line = strings.TrimSpace(line)
	tokens := strings.Split(line, " ")
	n, _ := strconv.Atoi(tokens[0])
	m, _ := strconv.Atoi(tokens[1])

	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		tokens = strings.Split(line, " ")
		for j := 0; j < m; j++ {
			arr[i][j], _ = strconv.Atoi(tokens[j])
		}
	}

	var sb strings.Builder
	for i := 0; i < n; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		tokens = strings.Split(line, " ")
		for j := 0; j < m; j++ {
			val, _ := strconv.Atoi(tokens[j])
			sb.WriteString(strconv.Itoa(arr[i][j] + val))
			sb.WriteString(" ")
		}
		sb.WriteString("\n")
	}

	fmt.Print(sb.String())
}
