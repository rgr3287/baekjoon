package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var expression string
	fmt.Scan(&expression)

	// '-'를 기준으로 식을 나눕니다.
	subexpressions := strings.Split(expression, "-")

	// 첫 번째 서브 식을 양수로 처리합니다.
	firstSubexpression := subexpressions[0]
	result := calculateSubexpression(firstSubexpression)

	// 나머지 서브 식을 모두 빼줍니다.
	for i := 1; i < len(subexpressions); i++ {
		subexpression := subexpressions[i]
		result -= calculateSubexpression(subexpression)
	}

	fmt.Println(result)
}

func calculateSubexpression(subexpression string) int {
	// '+'를 기준으로 서브 식을 나눕니다.
	operands := strings.Split(subexpression, "+")
	sum := 0
	for _, operand := range operands {
		num, _ := strconv.Atoi(operand)
		sum += num
	}
	return sum
}
