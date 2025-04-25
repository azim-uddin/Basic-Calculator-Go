package main

import (
	"fmt"
)

func calculation(operator string, number1, number2 float64) {
	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", number1+number2)
	case "-":
		fmt.Printf("Result: %.2f\n", number1-number2)
	case "*":
		fmt.Printf("Result: %.2f\n", number1*number2)
	case "/":
		if number2 == 0 {
			fmt.Println("Math Error: Division by zero!")
		} else {
			fmt.Printf("Result: %.2f\n", number1/number2)
		}
	default:
		fmt.Println("Invalid operator")
	}
}

func main() {
	var number1, number2 float64
	var operator string

	fmt.Println("The GO Calc")
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&number1)

	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&number2)

	calculation(operator, number1, number2)
}
