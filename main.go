package main

import "fmt"

func number() {
	var a, b, y int
	var operation string

	fmt.Scan(&a, &operation, &b)
	if operation == "+" {
		y = a + b
		fmt.Print(y)
	} else if operation == "-" {
		y = a - b
		fmt.Print(y)
	} else if operation == "*" {
		y = a * b
		fmt.Print(y)
	} else if operation == "/" && b != 0 {
		y = a / b
		fmt.Print(y)
	} else if b == 0 {
		fmt.Print("На ноль делить нельзя")
	}

}
func main() {
	number()
}
