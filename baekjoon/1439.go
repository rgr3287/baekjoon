package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	// 연속된 숫자 그룹의 개수를 초기화합니다.
	groupCount := 1

	// 첫 번째 숫자와 비교하기 위해 첫 번째 숫자를 저장합니다.
	prev := S[0]

	for i := 1; i < len(S); i++ {
		// 현재 숫자와 이전 숫자가 다르다면 그룹 개수를 증가시킵니다.
		if S[i] != prev {
			groupCount++
			// 이전 숫자를 현재 숫자로 업데이트합니다.
			prev = S[i]
		}
	}

	// 그룹 개수의 절반을 출력하면 최소 횟수가 됩니다.
	// (한 그룹을 모두 뒤집으면 다른 그룹도 같아집니다)
	result := groupCount / 2
	fmt.Println(result)
}
