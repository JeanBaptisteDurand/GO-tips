package main

import "fmt"

func main() {
    original := []int{1, 2, 3, 4, 5}
    slice := original[1:3] // Slice: [2, 3]

    fmt.Println("Original Slice:", original)
    fmt.Println("Sub Slice:", slice)

    slice = append(slice, 10) // Capacité pas encore dépassée, même tableau
    fmt.Println("After append 10 - Original:", original)
    fmt.Println("After append 10 - Sub Slice:", slice)

    slice = append(slice, 20, 30, 40) // Capacité dépassée, nouveau tableau
    fmt.Println("After append 20, 30, 40 - Original:", original)
    fmt.Println("After append 20, 30, 40 - Sub Slice:", slice)
}

// OUTPUT :

// Original Slice: [1 2 3 4 5]
// Sub Slice: [2 3]
// After append 10 - Original: [1 2 3 10 5]
// After append 10 - Sub Slice: [2 3 10]
// After append 20, 30, 40 - Original: [1 2 3 10 5]
// After append 20, 30, 40 - Sub Slice: [2 3 10 20 30 40]

