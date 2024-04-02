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

	for i := 0; i < 3; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		strArr := strings.Split(input, " ")

		arr := make([]int, 3)
		arr2 := make([]int, 3)
		arr3 := make([]int, 3)

		for j := 0; j < 3; j++ {
			arr[j], _ = strconv.Atoi(strArr[j])
		}
		for j := 0; j < 3; j++ {
			arr2[j], _ = strconv.Atoi(strArr[j+3])
		}

		for j := 2; j >= 0; j-- {
			if arr2[j] < arr[j] {
				arr2[j-1] -= 1
				arr2[j] += 60
			}
			arr3[j] = arr2[j] - arr[j]
		}

		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", arr3[j])
		}
		fmt.Println()
	}
}
