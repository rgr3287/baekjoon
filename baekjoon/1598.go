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

	num1, _ := strconv.Atoi(tokens[0])
	num2, _ := strconv.Atoi(tokens[1])

	var x1, x2, y1, y2 int

	if num1%4 == 0 {
		x1 = 3
		y1 = num1/4 - 1
	} else {
		x1 = num1%4 - 1
		y1 = num1 / 4
	}

	if num2%4 == 0 {
		x2 = 3
		y2 = num2/4 - 1
	} else {
		x2 = num2%4 - 1
		y2 = num2 / 4
	}

	result := 0

	if x1 > x2 {
		result += x1 - x2
	} else {
		result += x2 - x1
	}

	if y1 > y2 {
		result += y1 - y2
	} else {
		result += y2 - y1
	}

	fmt.Println(result)
}
