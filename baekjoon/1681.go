package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var NOT_INCLUDED_NUM string

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputArr := strings.Split(input, " ")

	studentNum, _ := strconv.Atoi(inputArr[0])
	NOT_INCLUDED_NUM = inputArr[1]

	count := 0
	for studentNum > 0 {
		count++
		if !strings.Contains(strconv.Itoa(count), NOT_INCLUDED_NUM) {
			studentNum--
		}
	}

	fmt.Print(count)
}
