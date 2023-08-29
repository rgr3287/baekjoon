package main

import (
	"fmt"
	"sort"
	"strconv"
)

type SerialNumber struct {
	Value     string
	SumDigits int
}

func main() {
	var n int
	fmt.Scan(&n)

	serials := make([]SerialNumber, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		serials[i].Value = s
		serials[i].SumDigits = calculateSumDigits(s)
	}

	// 정렬 조건에 따라 시리얼 번호 정렬
	sort.Slice(serials, func(i, j int) bool {
		if len(serials[i].Value) != len(serials[j].Value) {
			return len(serials[i].Value) < len(serials[j].Value)
		} else if serials[i].SumDigits != serials[j].SumDigits {
			return serials[i].SumDigits < serials[j].SumDigits
		}
		return serials[i].Value < serials[j].Value
	})

	// 정렬된 시리얼 번호 출력
	for _, s := range serials {
		fmt.Println(s.Value)
	}
}

func calculateSumDigits(s string) int {
	sum := 0
	for _, c := range s {
		if isDigit(byte(c)) {
			digit, _ := strconv.Atoi(string(c))
			sum += digit
		}
	}
	return sum
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
