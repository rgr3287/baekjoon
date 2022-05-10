package main

import "fmt"

func main() {
	var a, b, c uint
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	c = 2*b - a
	fmt.Println(c)
}

// 입력값이 a b 일때 c의 값을 출력하는 문제
// 5와 10일때 15
// 12와 15일때 18

// a = 5 b = 10  ==> c = 15
// a = 12 b = 15 ==> c = 18 라는 힌트를 줌

// a = 7 b = -2  ==> c = 3
// a = 17 b = 22  ==> c = 32
//  b*2 - a = c 라는 것을 간단하게 알 수 있다.
