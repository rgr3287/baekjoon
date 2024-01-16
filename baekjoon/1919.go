package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	alphabet := make([]int, 26)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	first := []rune(scanner.Text())

	scanner.Scan()
	second := []rune(scanner.Text())

	for _, c := range first {
		alphabet[c-97] += 1
	}
	for _, c := range second {
		alphabet[c-97] -= 1
	}

	result := 0
	for _, v := range alphabet {
		result += abs(v)
	}

	fmt.Println(result)
}

//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
