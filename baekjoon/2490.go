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
	arr := []string{"D", "C", "B", "A", "E"}
	var sb strings.Builder

	for i := 0; i < 3; i++ {
		input, _ := reader.ReadString('\n')
		tokens := strings.Fields(input)
		sum := 0
		for _, token := range tokens {
			num, _ := strconv.Atoi(token)
			sum += num
		}
		sb.WriteString(arr[sum] + "\n")
	}
	fmt.Print(sb.String())
}
