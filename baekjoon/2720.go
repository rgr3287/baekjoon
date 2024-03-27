package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	quarter := 25 // 쿼터 0.25$
	dime := 10    // 다임 0.10$
	nickel := 5   // 니켈 0.05$
	penny := 1    // 페니 0.01$

	var sb strings.Builder
	for i := 0; i < t; i++ {
		scanner.Scan()
		c, _ := strconv.Atoi(scanner.Text())

		sb.WriteString(strconv.Itoa(c/quarter) + " ")
		c %= quarter
		sb.WriteString(strconv.Itoa(c/dime) + " ")
		c %= dime
		sb.WriteString(strconv.Itoa(c/nickel) + " ")
		c %= nickel
		sb.WriteString(strconv.Itoa(c/penny) + "\n")
	}

	fmt.Print(sb.String())
}
