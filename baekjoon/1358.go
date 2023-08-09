package main

import (
	"fmt"
)

func main() {
	var W, H, X, Y, P int
	fmt.Scan(&W, &H, &X, &Y, &P)

	count := 0

	for i := 0; i < P; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		// 사각형 내부에 있는 경우
		if x >= X && x <= X+W && y >= Y && y <= Y+H {
			count++
		} else {
			// 원 내부에 있는 경우
			dx1, dy1 := x-X, y-(Y+H/2)
			dx2, dy2 := x-(X+W), y-(Y+H/2)
			if dx1*dx1+dy1*dy1 <= (H/2)*(H/2) || dx2*dx2+dy2*dy2 <= (H/2)*(H/2) {
				count++
			}
		}
	}

	fmt.Println(count)
}
