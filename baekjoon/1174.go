package main

import (
	"fmt"
	"sort"
)

var N int
var arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
var listData []int64

func main() {
	fmt.Scan(&N)

	DFS(0, 0)
	sort.Slice(listData, func(i, j int) bool { return listData[i] < listData[j] })

	if N <= len(listData) {
		fmt.Println(listData[N-1])
	} else {
		fmt.Println(-1)
	}
}

func DFS(num int64, inx int) {
	if !contains(listData, num) {
		listData = append(listData, num)
	}

	if inx >= 10 {
		return
	}

	DFS(num*10+int64(arr[inx]), inx+1)
	DFS(num, inx+1)
}

func contains(arr []int64, num int64) bool {
	for _, item := range arr {
		if item == num {
			return true
		}
	}
	return false
}
