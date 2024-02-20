package main

import "fmt"

func main() {
	var N, m, M, T, R int
	fmt.Scan(&N, &m, &M, &T, &R)

	res := 0
	move := 0
	initM := m

	for move != N {
		res++

		if m+T <= M {
			m += T
			move++
		} else {
			m -= R

			if m < initM {
				m = initM
			}
		}

		if m+T > M && m == initM {
			break
		}
	}

	if move != N {
		fmt.Println(-1)
	} else {
		fmt.Println(res)
	}
}
