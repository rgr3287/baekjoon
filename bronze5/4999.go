package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	if len(a) >= len(b) {
		fmt.Println("go")
	} else {
		fmt.Println("no")
	}
}

//package main
//
//import (
//"bufio"
//"os"
//)
//
//func main() {
//	defer wr.Flush()
//	sc.Scan()
//	p := sc.Bytes()
//	sc.Scan()
//	q := sc.Bytes()
//	if len(p) >= len(q) {
//		wr.WriteString("go")
//	} else {
//		wr.WriteString("no")
//	}
//}
//
//var (
//	sc *bufio.Scanner
//	wr *bufio.Writer
//)
//
//func init() {
//	sc = bufio.NewScanner(os.Stdin)
//	sc.Split(bufio.ScanWords)
//
//	wr = bufio.NewWriter(os.Stdout)
//}

// 시간 복잡도 줄이기
