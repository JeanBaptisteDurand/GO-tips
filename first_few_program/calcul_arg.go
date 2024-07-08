package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args

	if len(arg) == 4 {
		switch arg[2] {
			case "+":
				fmt.Println(arg[1], "+", arg[3], "=", add(arg[1], arg[3]))
			case "-":
				fmt.Println(arg[1], "-", arg[3], "=", sub(arg[1], arg[3]))
			case "*":
				fmt.Println(arg[1], "*", arg[3], "=", mul(arg[1], arg[3]))
			case "/":
				fmt.Println(arg[1], "/", arg[3], "=", div(arg[1], arg[3]))
			case "%":
				fmt.Println(arg[1], "%", arg[3], "=", mod(arg[1], arg[3]))
			default:
				fmt.Println("Error: operator is not valid")
		}
	} else {
		fmt.Println("Not good amount of arguments. Please provide 3 arguments: number operator number")
	}
}

func add(a, b string) int {
	num1, err1 := strconv.Atoi(a)
	num2, err2 := strconv.Atoi(b)
	if err1 != nil || err2 != nil {
		fmt.Println("Error converting string to integer")
		return 0
	}
	return num1 + num2
}

func sub(a, b string) int {
	num1, err1 := strconv.Atoi(a)
	num2, err2 := strconv.Atoi(b)
	if err1 != nil || err2 != nil {
		fmt.Println("Error converting string to integer")
		return 0
	}
	return num1 - num2
}

func mul(a, b string) int {
	num1, err1 := strconv.Atoi(a)
	num2, err2 := strconv.Atoi(b)
	if err1 != nil || err2 != nil {
		fmt.Println("Error converting string to integer")
		return 0
	}
	return num1 * num2
}

func div(a, b string) float64 {
	num1, err1 := strconv.ParseFloat(a, 64)
	num2, err2 := strconv.ParseFloat(b, 64)
	if err1 != nil || err2 != nil || num2 == 0 {
		fmt.Println("Error converting string to float or division by zero")
		return 0
	}
	return num1 / num2
}

func mod(a, b string) int {
	num1, err1 := strconv.Atoi(a)
	num2, err2 := strconv.Atoi(b)
	if err1 != nil || err2 != nil {
		fmt.Println("Error converting string to integer")
		return 0
	}
	return num1 % num2
}
