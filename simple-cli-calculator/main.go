package main

import (
	"fmt"
)

func main() {
	var fNum int
	var operator string
	var sNum int

	fmt.Print("Type your first number: ")
	fmt.Scan(&fNum)

	fmt.Print("Choose thhe operation (+,-,*,/): ")
	fmt.Scan(&operator)

	fmt.Print("Type your second number: ")
	fmt.Scan(&sNum)

	switch operator {
	case "+":
		fmt.Printf("Your result is: %d\n", fNum+sNum)
	case "-":
		fmt.Printf("Your result is: %d\n", fNum-sNum)
	case "*":
		fmt.Printf("Your result is: %d\n", fNum*sNum)
	case "/":
		fmt.Printf("Your result is: %d\n", fNum/sNum)
	}
}
