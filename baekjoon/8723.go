package main

import (
	"fmt"
	"sort"
)

func main() {
	var a [3]int
	for i := 0; i < 3; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a[:])
	if a[0] == a[2] {
		fmt.Println(2)
	} else if a[2]*a[2] == a[0]*a[0]+a[1]*a[1] {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
