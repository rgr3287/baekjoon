package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var group = [8][2]rune{{'A', 'C'}, {'D', 'F'}, {'G', 'I'}, {'J', 'L'}, {'M', 'O'}, {'P', 'S'}, {'T', 'V'}, {'W', 'Z'}}

func main() {
	var p, w int
	fmt.Scan(&p, &w)

	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	before := -1
	cnt := 0
	flag := false
	ret := 0

	for _, temp := range str {
		if unicode.IsLetter(temp) {
			for j := 0; j < 8; j++ {
				if temp >= group[j][0] && temp <= group[j][1] {
					if before == j {
						flag = true
					}
					cnt = (int(temp) - int(group[j][0]) + 1) * p
					before = j
				}
			}
		} else {
			cnt = p
			before = -1
		}

		if flag {
			cnt += w
		}
		ret += cnt

		flag = false
	}

	fmt.Println(ret)
}
