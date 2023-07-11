package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(num int) bool {
	str := strconv.Itoa(num)
	length := len(str)

	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}

	return true
}

func main() {
	var num int

	for {
		fmt.Scanln(&num)

		if num == 0 {
			break
		}

		if isPalindrome(num) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
