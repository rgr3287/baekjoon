package main

import "fmt"

func main() {
	var M, N int
	fmt.Scan(&M, &N)

	mapArr := make([][]int, M)
	for i := range mapArr {
		mapArr[i] = make([]int, N)
	}

	isVisited := make([][]bool, M)
	for i := range isVisited {
		isVisited[i] = make([]bool, N)
	}

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	cnt := 0
	sum := 0

	idx := 0
	nowX := 0
	nowY := 0
	isVisited[nowX][nowY] = true
	sum = 1

	for idx < 4 {
		nx := nowX + dx[idx]
		ny := nowY + dy[idx]

		if sum == N*M {
			fmt.Println(cnt)
			return
		}

		if nx >= 0 && ny >= 0 && nx < M && ny < N && !isVisited[nx][ny] {
			sum++
			isVisited[nx][ny] = true
			nowX = nx
			nowY = ny
		} else {
			idx++
			cnt++
		}

		if idx >= 4 {
			idx = 0
		}
	}
}
