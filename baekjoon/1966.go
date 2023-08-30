package main

import (
	"fmt"
)

type Document struct {
	index    int
	priority int
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n, &m)

		queue := make([]Document, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&queue[j].priority)
			queue[j].index = j
		}

		count := 1
		for {
			maxPriority := queue[0].priority
			maxIndex := 0
			for j := 1; j < len(queue); j++ {
				if queue[j].priority > maxPriority {
					maxPriority = queue[j].priority
					maxIndex = j
				}
			}

			for j := 0; j < maxIndex; j++ {
				queue = append(queue, queue[0])
				queue = queue[1:]
			}

			if queue[0].index == m {
				fmt.Println(count)
				break
			}

			count++
			queue = queue[1:]
		}
	}
}
