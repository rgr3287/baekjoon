package main

import (
	"container/list"
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	queue := list.New()
	for i := 1; i <= n; i++ {
		queue.PushBack(i)
	}

	var target, count int
	result := 0

	for i := 0; i < m; i++ {
		fmt.Scan(&target)

		targetIndex := 0
		for e := queue.Front(); e != nil; e = e.Next() {
			if e.Value.(int) == target {
				break
			}
			targetIndex++
		}

		if targetIndex <= queue.Len()/2 {
			for j := 0; j < targetIndex; j++ {
				queue.PushBack(queue.Front().Value)
				queue.Remove(queue.Front())
				count++
			}
		} else {
			for j := 0; j < queue.Len()-targetIndex; j++ {
				queue.PushFront(queue.Back().Value)
				queue.Remove(queue.Back())
				count++
			}
		}

		queue.Remove(queue.Front())
		result += count
		count = 0
	}

	fmt.Println(result)
}
