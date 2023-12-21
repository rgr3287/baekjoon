package main

import "fmt"

var dps [32]int64
var sym [32]int64

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(tileCode(n))
}

func tileCode(n int) int64 {
	return (tile(n)-symmetry(n))/2 + symmetry(n)
}

func tile(n int) int64 {
	if dps[n] != 0 {
		return dps[n]
	}
	if n == 1 {
		dps[1] = 1
		return dps[1]
	}
	if n == 2 {
		dps[2] = 3
		return dps[2]
	}

	dps[n] = tile(n-1) + 2*tile(n-2)

	return dps[n]
}

func symmetry(n int) int64 {
	if sym[n] != 0 {
		return sym[n]
	}
	if n == 1 {
		sym[1] = 1
		return sym[1]
	}
	if n == 2 {
		sym[2] = 3
		return sym[2]
	}
	if n == 3 {
		sym[3] = 1
		return sym[3]
	}
	if n == 4 {
		sym[4] = 5
		return sym[4]
	}

	sym[n] = symmetry(n-2) + 2*symmetry(n-4)

	return sym[n]
}
