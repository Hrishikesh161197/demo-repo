package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2, result float64
	var choice int
	fmt.Print("Enter 1st number :")
	fmt.Scan(&num1)
	fmt.Print("Enter 2nd number :")
	fmt.Scan(&num2)
	fmt.Println("Press 1 ADD / Press 2 SUB / Press 3 MUL / Press 4 DIV / Press 5 POWER")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		result = add(num1, num2)
		fmt.Println("Addition ", result)
	case 2:
		result = sub(num1, num2)
		fmt.Println("Substraction ", result)
	case 3:
		result = mul(num1, num2)
		fmt.Println("multiplication ", result)
	case 4:
		result = div(num1, num2)
		fmt.Println("Division ", result)
	case 5:
		result = math.Pow(num1, num2)
		fmt.Println("Power ", result)
	}
}
func add(num1 float64, num2 float64) float64 {
	return num1 + num2

}
func sub(num1 float64, num2 float64) float64 {
	return num1 - num2
}
func mul(num1 float64, num2 float64) float64 {
	return num1 * num2
}
func div(num1 float64, num2 float64) float64 {
	return num1 / num2
}
func Pow(num1 float64, num2 float64) float64 {
	return Pow(num1, num2)
}
