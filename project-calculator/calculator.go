package main

import (
	"fmt"
)

func main() {
	var first int
	var second int
	var choice int
	fmt.Println("Enter the number for calculation")
	fmt.Print("Enter first number:")
	fmt.Scan(&first)
	fmt.Print("Enter second number:")
	fmt.Scan(&second)
	fmt.Println("Enter your choice of operation:")
	fmt.Println("1 for add\n2 for Sub\n3 for Div\n4 for Mul")
	fmt.Print("Choice:")
	fmt.Scan(&choice)
	if choice == 1 {
		add(first, second)
	}
	if choice == 2 {
		sub(first, second)
	}
	if choice == 3 {
		div(first, second)
	}
	if choice == 4 {
		mul(first, second)
	}
}

func add(a int, b int) {
	c := (a + b)
	fmt.Printf("Addidtion of %v and %v is: %v", a, b, c)
}

func sub(a int, b int) {
	c := (a - b)
	fmt.Printf("Substraction of %v and %v is: %v", a, b, c)
}

func div(a int, b int) {
	c := (a / b)
	fmt.Printf("Division of %v and %v is: %v", a, b, c)
}

func mul(a int, b int) {
	c := (a * b)
	fmt.Printf("Multiplication of %v and %v is: %v", a, b, c)
}
