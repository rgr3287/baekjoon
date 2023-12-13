package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1] // Remove newline character

	var lenStack []int
	var mulStack []int

	cnt := 0
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c == '(' {
			cnt -= 1
			mulNum, _ := strconv.Atoi(string(str[i-1]))
			lenStack = append(lenStack, cnt)
			mulStack = append(mulStack, mulNum)
			cnt = 0
		} else if c == ')' {
			val := mulStack[len(mulStack)-1]
			mulStack = mulStack[:len(mulStack)-1]
			val *= cnt
			plus := lenStack[len(lenStack)-1]
			lenStack = lenStack[:len(lenStack)-1]
			cnt = plus + val
		} else {
			cnt++ // 숫자
		}
	}
	fmt.Println(cnt)
}
