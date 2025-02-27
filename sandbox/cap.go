package main

import "fmt"

func main() {
    var s []int
    for i := 0; i < 20; i++ {
        s = append(s, i)
        fmt.Printf("len: %d, cap: %d, s: %v\n", len(s), cap(s), s)
    }
}
