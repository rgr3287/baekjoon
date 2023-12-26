package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tc, _ := strconv.Atoi(scanner.Text())

	var result string

	for tc > 0 {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		for i := 2; i <= n; i++ {
			cnt := 0
			for n%i == 0 {
				n /= i
				cnt++
			}
			if cnt != 0 {
				result += fmt.Sprintf("%d %d\n", i, cnt)
			}
			if n == 0 {
				break
			}
		}

		tc--
	}

	fmt.Print(result)
}
