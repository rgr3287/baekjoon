package main

import "fmt"

func main() {
    var N int
    fmt.Scan(&N)

    for i := 0; i < N; i++ {
        for j := 0; j < i; j++ {
            fmt.Print(" ")
        }
        for k := N - i; k > 0; k-- {
            fmt.Print("*")
        }
        fmt.Println()
    }
}
