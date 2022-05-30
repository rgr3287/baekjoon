package main

import "fmt"

func main() {
	var n uint
	fmt.Scanln(&n)
	fmt.Println(n * (n - 1))
}

// 상하의 색이 안 겹칠경우
// 상의 빨강, 노랑 하의 빨강 노랑 일때 상의를 빨강을 고르면 하의는 빨강을 뺀 나머지가됨
// 그 위의 경우의 수를 구하면 n * (n - 1)
