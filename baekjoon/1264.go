package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		str := scanner.Text()

		if str == "#" {
			break
		}

		str = strings.ToLower(str)
		ans := 0

		for i := 0; i < len(str); i++ {
			t := str[i]
			if t == 'a' || t == 'e' || t == 'i' || t == 'o' || t == 'u' {
				ans++
			}
		}

		fmt.Println(ans)
	}
}
