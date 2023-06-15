package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	digits := strings.Split(input, "")

	numbers := make([]int, len(digits))
	for i, digit := range digits {
		num, _ := strconv.Atoi(digit)
		numbers[i] = num
	}

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	sortedStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), ""), "[]")
	fmt.Println(sortedStr)
}
