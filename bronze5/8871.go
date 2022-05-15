package main

import "fmt"

func main() {
	var n uint
	fmt.Scanf("%d", &n)
	fmt.Printf("%d %d", (n+1)*2, (n+1)*3)
}

// 입력값에 1을 더하고(시험라운드) 2, 3을 곱해 각각 출력
