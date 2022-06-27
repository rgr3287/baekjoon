package main

import "fmt"

func main() {
	var n, v int
	var a [100]int
	cnt := 0

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Scanf("%d", &v)

	for j := 0; j < n; j++ {
		if a[j] == v {
			cnt++
		}
	}
	fmt.Printf("%d", cnt)

}
