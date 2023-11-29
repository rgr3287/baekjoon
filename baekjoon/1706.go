package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	arr := make([][]byte, r)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < r; i++ {
		scanner.Scan()
		row := scanner.Text()
		arr[i] = make([]byte, c)
		for j := 0; j < c; j++ {
			arr[i][j] = row[j]
		}
	}

	min := string('z' + 1)
	for i := 0; i < r; i++ {
		var sb string
		for j := 0; j <= c; j++ {
			if j == c || arr[i][j] == '#' {
				if len(sb) >= 2 && min > sb {
					min = sb
				}
				sb = ""
			} else {
				sb += string(arr[i][j])
			}
		}
	}

	for j := 0; j < c; j++ {
		var sb string
		for i := 0; i <= r; i++ {
			if i == r || arr[i][j] == '#' {
				if len(sb) >= 2 && min > sb {
					min = sb
				}
				sb = ""
			} else {
				sb += string(arr[i][j])
			}
		}
	}

	fmt.Println(min)
}
