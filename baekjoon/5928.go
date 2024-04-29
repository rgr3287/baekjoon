package main

import (
	"fmt"
)

func main() {
	var d1, h1, m1, total int

	fmt.Scan(&d1)
	fmt.Scan(&h1)
	fmt.Scan(&m1)

	if d1 < 11 {
		fmt.Println("-1")
	} else if d1 == 11 {
		if h1 < 11 {
			fmt.Println("-1")
		} else if h1 == 11 {
			if m1 >= 11 {
				total += m1 - 11
				fmt.Println(total)
			} else {
				fmt.Println("-1")
			}
		} else if h1 > 11 {
			if m1 >= 11 {
				total += (h1 - 11) * 60 + m1 - 11
				fmt.Println(total)
			} else {
				total += (h1 - 12) * 60 + 60 - 11 + m1
				fmt.Println(total)
			}
		}
	} else if d1 > 11 {
		if h1 > 11 {
			if m1 >= 11 {
				total += (d1-11)*24*60 + (h1-11)*60 + m1 - 11
				fmt.Println(total)
			} else if m1 < 11 {
				total += (d1-11)*24*60 + (h1-12)*60 + 60 - 11 + m1
				fmt.Println(total)
			}
		} else if h1 == 11 {
			if m1 >= 11 {
				total += (d1-11)*24*60 + m1 - 11
				fmt.Println(total)
			} else if m1 < 11 {
				total += (d1-12)*24*60 + 23*60 + 60 - 11 + m1
				fmt.Println(total)
			}
		} else if h1 < 11 {
			if m1 >= 11 {
				total += (d1-12)*24*60 + (24-11+h1)*60 + m1 - 11
				fmt.Println(total)
			} else if m1 < 11 {
				total += (d1-12)*24*60 + (24-12+h1)*60 + 60 - 11 + m1
				fmt.Println(total)
			}
		}
	}
}
