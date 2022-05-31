package main

import (
	"fmt"
)

func main() {
	var n uint
	fmt.Scanln(&n)
	fmt.Println(1)
}

// ((x+1)*(y+1)-xy-x-y)^n => ((x+1)*(y+1)-xy-x-y)^n = (xy + x + y + 1 - xy - x - y)^n = 1^n = 1
