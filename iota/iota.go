package main

import "fmt"

type Color int

const (
	ColorBlue Color = iota
	ColorBlack
	ColorYellow
	ColorPink
)

func main() {
	fmt.Println("The color is", ColorBlack)
}
