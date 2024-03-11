package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 문자열 입력
	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1] // 개행 문자 제거

	// 문자열을 문자 배열로 변환
	charArr := []rune(str)
	len := len(charArr)

	// a, b 배열 초기화
	a := make([]int, len)
	b := make([]int, len)

	for i := 0; i < len-1; i++ {
		j := len - 1 - i

		if i == 0 {
			a[i] = int(charArr[i] - '0')
			b[j] = int(charArr[j] - '0')
			continue
		}

		a[i] = a[i-1] * int(charArr[i]-'0')
		b[j] = b[j+1] * int(charArr[j]-'0')
	}

	// 유진수 찾기
	yujinsu := false
	for i := 0; i < len-1; i++ {
		if a[i] == b[i+1] {
			yujinsu = true
			break
		}
	}

	// 결과 출력
	if yujinsu {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
